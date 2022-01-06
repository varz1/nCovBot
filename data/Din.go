package data

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/go-resty/resty/v2"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"github.com/varz1/nCovBot/model"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"time"
)

var request = resty.New()

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
}

// Cro19map 定时更新数据图表
func Cro19map() {
	c := cron.New()
	c.AddFunc("@every 6h", func() {
		if err := GetChMap(); err != nil {
			logrus.Info("更新失败请重试")
		} else {
			logrus.Info("更新成功")
		}
	})
	c.AddFunc("@every 60s", func() {
		Ping()
	})
	c.Start()
}

// GetChMap 无头浏览器爬取数据图表
func GetChMap() error {
	logrus.WithField("GetChMap", "开始爬取图表")
	var url = "https://voice.baidu.com/act/newpneumonia/newpneumonia"
	var selMap = "#virus-map"
	var selTrend = "div.VirusTrend_1-1-321_1U4OF0"
	pwd, _ := os.Getwd()
	fileMap := "/public/virusMap.png"
	fileTrend := "/public/virusTrend.png"
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	chromedp.ExecPath(os.Getenv("GOOGLE_CHROME_SHIM"))
	// 超时设置
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	var buf []byte
	if err := chromedp.Run(ctx,
		Screenshot(url, selMap, &buf)); err != nil {
		return err
	}
	if err := ioutil.WriteFile(pwd+fileMap, buf, 0o644); err != nil {
		return err
	}
	if err := chromedp.Run(ctx,
		Screenshot(url, selTrend, &buf)); err != nil {
		return err
	}
	if err := ioutil.WriteFile(pwd+fileTrend, buf, 0o644); err != nil {
		return err
	}
	return nil
}

func GetState(num int) int64 {
	log1 := logrus.WithField("GetState", "查看状态")
	pwd, _ := os.Getwd()
	var updateTime int64 = 0
	var files []string
	files = append(files, "/public/virusMap.png")
	files = append(files, "/public/virusTrend.png")
	for k, f := range files {
		info, err := os.Stat(pwd + f)
		if k == num {
			updateTime = info.ModTime().Unix()
		}
		if err != nil {
			log1.Info("尚未更新map")
		} else {
			log1.Info(f + "已更新" + "上次更新时间为" + info.ModTime().String())
		}
	}
	return updateTime
}

// Screenshot 截图
func Screenshot(url, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// GetOverall 获取疫情概览
func GetOverall() model.OverallData {
	log1 := logrus.WithField("func", "GetOverall")
	log1.Info("开始请求数据概览API")
	var overall struct {
		Results []model.OverallData `json:"results"`
	}
	resp, err := request.R().SetResult(&overall).Get("https://lab.isaaclin.cn/nCoV/api/overall")
	if err != nil {
		log1.WithField("resp err", err).Error(err)
	}
	if resp.StatusCode() != 200 {
		log1.WithField("status err", err).Error(err)
	}
	if err == nil {
		log1.Info("请求数据概览API成功")
	}
	return overall.Results[0]
}

// GetAreaData 获取地区数据
func GetAreaData(area string) model.ProvinceData {
	log1 := logrus.WithField("func", "GetAreaData")
	var res struct {
		Results []model.ProvinceData `json:"results"`
	}
	log1.Println("开始请求地区数据API")
	resp, err := request.R().SetResult(&res).SetQueryString("province=" + area).
		Get("https://lab.isaaclin.cn/nCoV/api/area?")
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
func GetNews() []model.NewsData {
	log1 := logrus.WithField("func", "GetNews")
	var res struct {
		Results []model.NewsData `json:"results"`
	}
	log1.Println("开始请求新闻数据API")
	resp, err := request.R().SetResult(&res).Get("https://lab.isaaclin.cn/nCoV/api/news")
	if err != nil || resp.StatusCode() != 200 {
		log1.WithField("请求失败", "新闻API").Errorln(err)
		return []model.NewsData{}
	}
	if err == nil {
		log1.Info("请求新闻数据API成功")
	}
	return res.Results
}

// GetRiskLevel 获取风险等级
func GetRiskLevel(level string) []model.RiskArea {
	log1 := logrus.WithField("func", "GetRiskLevel")
	var res struct {
		Data []model.RiskArea `json:"data"`
	}
	count := 0
	log1.Println("开始请求风险地区API")
	resp, err := request.R().SetResult(&res).Get("https://eyesight.news.qq.com/sars/riskarea")
	if err != nil || resp.StatusCode() != 200 {
		log1.WithField("请求失败", "风险地区").Error(err)
	}
	if err == nil {
		log1.Info("请求风险地区API成功")
	}
	risk := res.Data
	if len(risk) == 0 {
		return nil
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
	switch level {
	case "2":
		return risk[:count]
	default:
		return risk[count:]
	}
}

func Ping() {
	resp, err := request.R().Get("https://ncovbott.herokuapp.com/hi")
	if err != nil {
		logrus.Info("ping err")
	} else {
		logrus.Printf("Ping 成功 状态码: %v", resp.StatusCode())
	}
}
