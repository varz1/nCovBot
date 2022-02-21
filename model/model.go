package model

import (
	"bytes"
)

type Risks struct {
	High []RiskArea
	Mid []RiskArea
	Tm string
}

// Chartt 图表结构体
type Chartt struct {
	Pie bytes.Buffer
	Date string
}

// RiskArea 风险地区数据
type RiskArea struct {
	Area string `json:"area"`
	Type string `json:"type"`
	Sid  int    `json:"sid"`
}

// World 世界数据
type World struct {
	Continent string `json:"continent"`
	Name      string `json:"name"`
	PubDate   string `json:"pub_date"`
	Confirm   int    `json:"confirm"`
}

// Add 新增本土case数据时间线
type Add struct {
	Date            string `json:"date"`
	LocalConfirmAdd int    `json:"localConfirmadd"`
	Year            string `json:"y"`
}

// News 请求的新闻数据
type News struct {
	Results []struct {
		PubDate    string      `json:"pubDate"`
		Title      string      `json:"title"`
		Summary    string      `json:"summary"`
		InfoSource string      `json:"infoSource"`
		SourceUrl  string      `json:"sourceUrl"`
	} `json:"results"`
}
// ProvinceData 请求的省份详细数据
type ProvinceData struct {
	LocationId            int         `json:"locationId"`
	ContinentName         string      `json:"continentName"`
	CountryName           string      `json:"countryName"`
	CountryFullName       interface{} `json:"countryFullName"`
	ProvinceName          string      `json:"provinceName"`
	CurrentConfirmedCount int         `json:"currentConfirmedCount"` //现存确诊(含境外
	ConfirmedCount        int         `json:"confirmedCount"`        //累计确诊
	SuspectedCount        int         `json:"suspectedCount"`
	CuredCount            int         `json:"curedCount"` //累计治愈
	DeadCount             int         `json:"deadCount"`  //死亡
	Cities                []Cities    `json:"cities"`
	UpdateTime            int64       `json:"updateTime"`
}

type Cities struct {
	CityName                 string `json:"cityName"`
	ConfirmedCount           int    `json:"confirmedCount"` //累计确诊
	SuspectedCount           int    `json:"suspectedCount"`
	CuredCount               int    `json:"curedCount"`      //累计治愈
	DeadCount                int    `json:"deadCount"`       //死亡
	HighDangerCount          int    `json:"highDangerCount"` //高风险地区数量
	MidDangerCount           int    `json:"midDangerCount"`  //中风险地区数量
	LocationId               int    `json:"locationId"`
	CurrentConfirmedCountStr string `json:"currentConfirmedCountStr"` //现存本土确诊
	CityEnglishName          string `json:"cityEnglishName,omitempty"`
}

type Overall struct {
	Data struct {
		Diseaseh5Shelf struct {
			LastUpdateTime string `json:"lastUpdateTime"`
			ChinaTotal     struct {
				LocalConfirm       int `json:"localConfirm"`
				LocalConfirmH5     int `json:"localConfirmH5"`
				Confirm            int `json:"confirm"`
				NoInfect           int `json:"noInfect"`
				Dead               int `json:"dead"`
				NowConfirm         int `json:"nowConfirm"`
				LocalAccConfirm    int `json:"local_acc_confirm"`
				NoInfectH5         int `json:"noInfectH5"`
				Heal               int `json:"heal"`
				Suspect            int `json:"suspect"`
				NowSevere          int `json:"nowSevere"`
				ImportedCase       int `json:"importedCase"`
				ShowLocalConfirm   int `json:"showLocalConfirm"`
				Showlocalinfeciton int `json:"showlocalinfeciton"`
			} `json:"chinaTotal"`
			ChinaAdd struct {
				NoInfect       int `json:"noInfect"`
				LocalConfirm   int `json:"localConfirm"`
				NoInfectH5     int `json:"noInfectH5"`
				Confirm        int `json:"confirm"`
				Heal           int `json:"heal"`
				NowSevere      int `json:"nowSevere"`
				ImportedCase   int `json:"importedCase"`
				LocalConfirmH5 int `json:"localConfirmH5"`
				Dead           int `json:"dead"`
				NowConfirm     int `json:"nowConfirm"`
				Suspect        int `json:"suspect"`
			} `json:"chinaAdd"`
		} `json:"diseaseh5Shelf"`
	} `json:"data"`
}

type OverallWorld struct {
	Data struct {
		WomWorld struct {
			PubDate        string `json:"PubDate"`
			Y              string `json:"y"`
			Date           string `json:"date"`
			NowConfirm     int    `json:"nowConfirm"`
			NowConfirmAdd  int    `json:"nowConfirmAdd"`
			Confirm        int    `json:"confirm"`
			ConfirmAdd     int    `json:"confirmAdd"`
			Heal           int    `json:"heal"`
			HealAdd        int    `json:"healAdd"`
			Dead           int    `json:"dead"`
			DeadAdd        int    `json:"deadAdd"`
			Deathrate      int    `json:"deathrate"`
			Curerate       int    `json:"curerate"`
			LastUpdateTime string `json:"lastUpdateTime"`
		} `json:"WomWorld"`
	} `json:"data"`
}