package data

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/json-iterator/go"
	"github.com/varz1/nCovBot/model"
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

func GetAreas() (model.Result, error) {
	resp, err := request.R().Get("https://lab.isaaclin.cn/nCoV/api/provinceName")
	if err != nil {
		return model.Result{}, err
	}
	if resp.StatusCode() != 200 {
		return model.Result{}, errors.New("request wrong")
	}
	var areas model.Result
	if resp.Body() == nil {
		return model.Result{}, errors.New("empty result")
	}
	err = json.Unmarshal(resp.Body(), &areas)
	if err != nil {
		return model.Result{}, err
	}
	return areas, nil
}
