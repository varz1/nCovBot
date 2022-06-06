package variables

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	EnvToken     = os.Getenv("TOKEN")
	EnvTestToken = os.Getenv("TestToken")
	EnvBaseUrl   = os.Getenv("baseURL")
	EnvAdminId   = os.Getenv("AdminId")
)

func init() {
	if EnvAdminId == "" || EnvBaseUrl == "" || EnvToken == "" || EnvTestToken == "" {
		logrus.Warn("缺少环境变量！")
	}
}

const Blog = "https://varz1.github.io/"

// apis
const (
	OVERALL1      = "https://lab.isaaclin.cn/nCoV/api/overall"                                                    //新闻概览API
	AREA          = "https://lab.isaaclin.cn/nCoV/api/area?"                                                      //地区数据API
	NEWS          = "https://lab.isaaclin.cn/nCoV/api/news"                                                       //新闻API
	RISK          = "https://eyesight.news.qq.com/sars/riskarea"                                                  //风险地区API
	LOCALTimeLine = "https://api.inews.qq.com/newsqa/v1/query/inner/publish/modules/list?modules=chinaDayAddList" //本土新增时间线
	WORLD         = "https://api.inews.qq.com/newsqa/v1/automation/modules/list?"                                 //世界数据API
	OVERALL       = "https://api.inews.qq.com/newsqa/v1/query/inner/publish/modules/list?modules=diseaseh5Shelf"  //疫情概览
)

//words
const (
	NA = "北美洲"
	AS = "亚洲"
	SA = "南美洲"
	EU = "欧洲"
	OC = "大洋州"
	AF = "非洲"
)

//title
const (
	PIE   = "各大洲累计确诊比"
	TREND = "七天内新增本土病例"
)

// DAY data
const (
	DAY float64 = 86400 // 一天的时间戳
)
