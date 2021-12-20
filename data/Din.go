package data

import (
	"github.com/go-resty/resty/v2"
	"github.com/json-iterator/go"
	"github.com/varz1/nCovBot/model"
	"log"
)

var (
	request = resty.New()
	json    = jsoniter.ConfigFastest
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
}


func Overall() model.OverallData {
	resp, err := request.R().Get("https://lab.isaaclin.cn/nCoV/api/overall")
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode() != 200 {
		log.Println(err)
	}
	var overall struct{
		Results []model.OverallData `json:"results"`
	}
	err = json.Unmarshal(resp.Body(), &overall)
	return overall.Results[0]
}
