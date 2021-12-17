package model

type Result struct {
	Areas   []string `json:"results"`
	Success bool     `json:"success"`
}
type News struct {
	Results []struct {
		CurrentConfirmedCount int `json:"currentConfirmedCount"`
		CurrentConfirmedIncr  int `json:"currentConfirmedIncr"`
		ConfirmedCount        int `json:"confirmedCount"`
		ConfirmedIncr         int `json:"confirmedIncr"`
		SuspectedCount        int `json:"suspectedCount"`
		SuspectedIncr         int `json:"suspectedIncr"`
		CuredCount            int `json:"curedCount"`
		CuredIncr             int `json:"curedIncr"`
		DeadCount             int `json:"deadCount"`
		DeadIncr              int `json:"deadIncr"`
		SeriousCount          int `json:"seriousCount"`
		SeriousIncr           int `json:"seriousIncr"`
		GlobalStatistics      struct {
			CurrentConfirmedCount       int `json:"currentConfirmedCount"`
			ConfirmedCount              int `json:"confirmedCount"`
			CuredCount                  int `json:"curedCount"`
			DeadCount                   int `json:"deadCount"`
			CurrentConfirmedIncr        int `json:"currentConfirmedIncr"`
			ConfirmedIncr               int `json:"confirmedIncr"`
			CuredIncr                   int `json:"curedIncr"`
			DeadIncr                    int `json:"deadIncr"`
			YesterdayConfirmedCountIncr int `json:"yesterdayConfirmedCountIncr"`
		} `json:"globalStatistics"`
		GeneralRemark string `json:"generalRemark"`
		Remark1       string `json:"remark1"`
		Remark2       string `json:"remark2"`
		Remark3       string `json:"remark3"`
		Remark4       string `json:"remark4"`
		Remark5       string `json:"remark5"`
		Note1         string `json:"note1"`
		Note2         string `json:"note2"`
		Note3         string `json:"note3"`
		UpdateTime    int64  `json:"updateTime"`
	} `json:"results"`
	Success bool `json:"success"`
}

type Province struct {
	Results []struct {
		LocationId            int         `json:"locationId"`
		ContinentName         string      `json:"continentName"`
		ContinentEnglishName  string      `json:"continentEnglishName"`
		CountryName           string      `json:"countryName"`
		CountryEnglishName    string      `json:"countryEnglishName"`
		CountryFullName       interface{} `json:"countryFullName"`
		ProvinceName          string      `json:"provinceName"`
		ProvinceEnglishName   string      `json:"provinceEnglishName"`
		ProvinceShortName     string      `json:"provinceShortName"`
		CurrentConfirmedCount int         `json:"currentConfirmedCount"`
		ConfirmedCount        int         `json:"confirmedCount"`
		SuspectedCount        int         `json:"suspectedCount"`
		CuredCount            int         `json:"curedCount"`
		DeadCount             int         `json:"deadCount"`
		Comment               string      `json:"comment"`
		Cities                []struct {
			CityName                 string `json:"cityName"`
			CurrentConfirmedCount    int    `json:"currentConfirmedCount"`
			ConfirmedCount           int    `json:"confirmedCount"`
			SuspectedCount           int    `json:"suspectedCount"`
			CuredCount               int    `json:"curedCount"`
			DeadCount                int    `json:"deadCount"`
			HighDangerCount          int    `json:"highDangerCount"`
			MidDangerCount           int    `json:"midDangerCount"`
			LocationId               int    `json:"locationId"`
			CurrentConfirmedCountStr string `json:"currentConfirmedCountStr"`
			CityEnglishName          string `json:"cityEnglishName,omitempty"`
		} `json:"cities"`
		UpdateTime int64 `json:"updateTime"`
	} `json:"results"`
	Success bool `json:"success"`
}