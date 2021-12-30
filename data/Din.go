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

// Cro19map 定时更新地图
func Cro19map() {
	c := cron.New()
	c.AddFunc("@every 12h", func() {
		GetChMap()
	})
	c.Start()
}

// GetChMap 获取中国疫情地图
func GetChMap() {
	log1 := logrus.WithField("func GetMap", "chromeDp爬取地图")
	var url = "https://voice.baidu.com/act/newpneumonia/newpneumonia"
	var sel = "#virus-map"
	pwd, _ := os.Getwd()
	file := "/public/map.png"
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent(`Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36`),
	}
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	chromedp.ExecPath(os.Getenv("GOOGLE_CHROME_SHIM"))
	// 超时设置
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	var buf []byte
	info, err := os.Stat(pwd + file)
	if err != nil {
		log1.WithError(err).Logln(2, "获取文件更新时间失败")
	}
	log1.Info("上次更新时间为" + info.ModTime().String())
	if err := chromedp.Run(ctx,
		Screenshot(url, sel, &buf)); err != nil {
		log1.Error(err)
	}
	if err := ioutil.WriteFile(pwd+file, buf, 0o644); err != nil {
		log1.Error(err)
	}
	log1.Info("地图已更新")
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
	log1.Info("开始请求新闻概览API")
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
		log1.Info("请求新闻概览API成功")
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
