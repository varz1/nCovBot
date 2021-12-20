package model

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Areas struct {
	Types       string
	area        []string
	AreaMessage tgbotapi.Message
}

type OverallMessage struct {
	Overall     tgbotapi.MessageConfig
	OverallData `json:"results"`
}

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
