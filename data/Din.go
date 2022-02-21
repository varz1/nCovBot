package data

import (
	"github.com/go-resty/resty/v2"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/varz1/nCovBot/cache"
	"github.com/varz1/nCovBot/model"
	"github.com/varz1/nCovBot/variables"
	"sort"
	"strconv"
	"time"
)

var (
	request = resty.New()
	timer   = cron.New()
	C       = cache.New()
	Debug   = false
)

func init() {
	header := map[string]string{
		"accept":                    `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`,
		"accept-language":           `zh-CN,zh-TW;q=0.9,zh;q=0.8,en-US;q=0.7,en;q=0.6`,
		"sec-ch-ua":                 ` " Not;A Brand";v="99", "Google Chrome";v="91", "Chromium";v="91"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "none",
		"sec-fetch-user":            "?1",
		"upgrade-insecure-requests": "1",
	}
	request.SetHeaders(header)
	if !Debug {
		logrus.Info("开始初始化数据")
		GetNews()
		GetRiskLevel()
		GetOverall()
		GetWorld()
		timer.AddFunc("@every 6h", func() {
			logrus.Info("开始更新数据")
			GetNews()
			GetRiskLevel()
			GetOverall()
			GetWorld()
		})
		timer.Start()
	}
}

// GetOverall 获取疫情概览
func GetOverall() {
	log1 := logrus.WithField("func", "GetOverall")
	var overall model.Overall
	resp, err := request.R().SetResult(&overall).Get(variables.OVERALL)
	if err != nil {
		log1.WithField("resp err", err).Error(err)
		return
	}
	if resp.StatusCode() != 200 {
		log1.WithField("status err", err).Error(err)
		return
	}
	log1.Info("请求本土数据概览API成功")
	C.Set("overall", overall)
}

// GetWorld 获取世界疫情概览
func GetWorld() {
	log1 := logrus.WithField("func", "GetOverall")
	var overall model.OverallWorld
	resp, err := request.R().SetResult(&overall).SetQueryString("modules=WomWorld").Get(variables.WORLD)
	if err != nil {
		log1.WithField("resp err", err).Error(err)
		return
	}
	if resp.StatusCode() != 200 {
		log1.WithField("status err", err).Error(err)
		return
	}
	log1.Info("请求世界数据概览API成功")
	C.Set("world", overall)
}

// GetAreaData 获取地区数据
func GetAreaData(area string) model.ProvinceData {
	log1 := logrus.WithField("func", "GetAreaData")
	var res struct {
		Results []model.ProvinceData `json:"results"`
	}
	log1.Println("开始请求地区数据API")
	resp, err := request.R().SetResult(&res).SetQueryString("province=" + area).Get(variables.AREA)
	if err != nil || resp.StatusCode() != 200 {
		log1.WithField("请求地区数据失败", "").Errorln(err)
		return model.ProvinceData{}
	}
	if err == nil {
		log1.Info("请求地区数据API成功")
	}
	return res.Results[0]
}

// GetNews 获取新闻
func GetNews() {
	log1 := logrus.WithField("func", "GetNews")
	var res model.News
	resp, err := request.R().SetResult(&res).Get(variables.NEWS)
	if err != nil || resp.StatusCode() != 200 {
		log1.WithField("请求失败", "新闻API").Errorln(err)
		return
	}
	log1.Info("请求新闻数据API成功")
	//NewsData = res.Results
	C.Set("news", res)
}

// GetRiskLevel 获取风险等级
func GetRiskLevel() {
	log1 := logrus.WithField("func", "GetRiskLevel")
	tm := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04")
	var res struct {
		Data []model.RiskArea `json:"data"`
	}
	count := 0
	resp, err := request.R().SetResult(&res).Get(variables.RISK)
	if err != nil || resp.StatusCode() != 200 {
		log1.WithField("请求失败", "风险地区").Error(err)
		return
	}
	log1.Info("请求风险地区API成功")
	var riskdata model.Risks
	risk := res.Data
	if len(risk) == 0 {
		riskdata.Mid = nil
		riskdata.High = nil
	}
	sort.SliceStable(risk, func(i, j int) bool {
		m, _ := strconv.Atoi(risk[i].Type)
		n, _ := strconv.Atoi(risk[j].Type)
		if m < n {
			return false
		}
		return true
	})
	for _, v := range risk {
		if v.Type == "2" {
			count = count + 1
		}
	}
	riskdata.High = risk[:count]
	riskdata.Mid = risk[count:]
	riskdata.Tm = tm
	C.Set("risk", riskdata)
}

// GetAdds 获取新增数据用于绘制图表
func GetAdds(day int) []model.Add {
	logrus.WithField("请求本土新增数据", "GetAdds")
	// day:需要的数据个数 这里只渲染最近七天数据
	var res struct {
		Data struct {
			Adds []model.Add `json:"chinaDayAddList"`
		} `json:"data"`
	}
	resp, _ := request.R().SetResult(&res).Get(variables.LOCALTimeLine)
	if !resp.IsSuccess() {
		logrus.Error("请求数据失败")
		return nil
	}
	return res.Data.Adds[len(res.Data.Adds)-day:]
}

// GetWorldData 获取大洲累积确诊返回map
func GetWorldData() (map[string]int, error) {
	logrus.WithField("请求大洲数据", "GetWorldData")
	var res struct {
		Data struct {
			WomAboard []model.World `json:"WomAboard"`
		} `json:"data"`
	}
	resp, err := request.R().SetResult(&res).SetQueryString("modules=WomAboard").Get(variables.WORLD)
	if resp.StatusCode() != 200 {
		logrus.Error("获取世界数据失败")
		return nil, err
	}
	continent := make(map[string]int)
	for _, v := range res.Data.WomAboard {
		switch v.Continent {
		case "北美洲":
			continent[variables.NA] = continent[variables.NA] + v.Confirm
		case "亚洲":
			continent[variables.AS] = continent[variables.AS] + v.Confirm
		case "南美洲":
			continent[variables.SA] = continent[variables.SA] + v.Confirm
		case "欧洲":
			continent[variables.EU] = continent[variables.EU] + v.Confirm
		case "大洋洲":
			continent[variables.OC] = continent[variables.OC] + v.Confirm
		case "非洲":
			continent[variables.AF] = continent[variables.AF] + v.Confirm
		}
	}
	return continent, nil
}
