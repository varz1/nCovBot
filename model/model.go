package model

import (
	"reflect"
)

type RiskArea struct {
	Area string `json:"area"`
	Type string `json:"type"`
	Sid  int    `json:"sid"`
}

// World 世界数据
type World struct {
	Continent  string `json:"continent"`
	Name       string `json:"name"`
	PubDate    string `json:"pub_date"`
	Confirm int    `json:"confirm"`
}

// Add 新增本土case数据
type Add struct {
	Date            string `json:"date"`
	LocalConfirmAdd int    `json:"localConfirmadd"`
	Year            string `json:"y"`
}

// NewsData 请求的新闻数据
type NewsData struct {
	PubDate    string `json:"pubDate"`    //发布时间
	Title      string `json:"title"`      //标题
	Summary    string `json:"summary"`    //新闻概要
	InfoSource string `json:"infoSource"` //新闻来源
	SourceUrl  string `json:"sourceUrl"`  //新闻源地址
}

// IsEmpty ProvinceData 判空
func (msg ProvinceData) IsEmpty() bool {
	return reflect.DeepEqual(msg, ProvinceData{})
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

// OverallData 请求的概要数据
type OverallData struct {
	CurrentConfirmedCount int              `json:"currentConfirmedCount"` //现有确诊(含港澳台,境外
	CurrentConfirmedIncr  int              `json:"currentConfirmedIncr"`  //现有确诊较昨日新增
	ConfirmedCount        int              `json:"confirmedCount"`        //累计确诊
	ConfirmedIncr         int              `json:"confirmedIncr"`         //累计确诊新增
	SuspectedCount        int              `json:"suspectedCount"`        //境外输入
	SuspectedIncr         int              `json:"suspectedIncr"`         //境外输入新增
	CuredCount            int              `json:"curedCount"`            //累计治愈
	CuredIncr             int              `json:"curedIncr"`             //新增治愈
	DeadCount             int              `json:"deadCount"`             //死亡人数
	DeadIncr              int              `json:"deadIncr"`              //死亡新增
	SeriousCount          int              `json:"seriousCount"`          //现存无症状
	SeriousIncr           int              `json:"seriousIncr"`           //现存无症状新增
	GlobalStatistics      GlobalStatistics `json:"globalStatistics"`
	UpdateTime            int64            `json:"updateTime"` //更新时间戳
}
type GlobalStatistics struct {
	CurrentConfirmedCount       int `json:"currentConfirmedCount"`       //全球现存确诊
	ConfirmedCount              int `json:"confirmedCount"`              //全球累计确诊
	CuredCount                  int `json:"curedCount"`                  //全球累计治愈
	DeadCount                   int `json:"deadCount"`                   //全球累计死亡
	CurrentConfirmedIncr        int `json:"currentConfirmedIncr"`        //全球现存确诊新增
	ConfirmedIncr               int `json:"confirmedIncr"`               //全球累计确诊新增
	CuredIncr                   int `json:"curedIncr"`                   //全球累计治愈新增
	DeadIncr                    int `json:"deadIncr"`                    //全球累计确诊新增
	YesterdayConfirmedCountIncr int `json:"yesterdayConfirmedCountIncr"` //全球现存确诊新增
}
