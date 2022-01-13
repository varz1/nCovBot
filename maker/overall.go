package maker

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	data2 "github.com/varz1/nCovBot/data"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Overall() {
	text := strings.Builder{}
	for overall := range channel.OverallUpdateChannel {
		mapTime, err := data2.GetState()
		if err != nil {
			log.Println(err)
			msg := tgbotapi.NewMessage(overall.Message.Chat.ID, "获取图表失败")
			channel.MessageChannel <- msg
			return
		}
		data := data2.GetOverall()//
		tm := time.Unix(data.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		tm1 := time.Unix(mapTime, 0).Format("2006-01-02 15:04")
		text.WriteString("🇨🇳中国疫情概况:")
		text.WriteString("\n现存确诊(含港澳台):" + strconv.Itoa(data.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(data.CurrentConfirmedIncr))
		text.WriteString("\n现存无症状:" + strconv.Itoa(data.SeriousCount) + " ⬆️" + strconv.Itoa(data.SeriousIncr))
		text.WriteString("\n境外输入:" + strconv.Itoa(data.SuspectedCount) + " ⬆️" + strconv.Itoa(data.SuspectedIncr))
		text.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount) + " ⬆️" + strconv.Itoa(data.ConfirmedIncr))
		text.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount) + " ⬆️" + strconv.Itoa(data.CuredIncr))
		text.WriteString("\n累计死亡" + strconv.Itoa(data.DeadCount) + " ⬆️" + strconv.Itoa(data.DeadIncr))
		text.WriteString("\n地图更新时间:" + tm1)
		text.WriteString("\n数据更新时间:" + tm)
		var url = os.Getenv("baseURL") + "virusMap.png" + "?a=" + strconv.FormatInt(time.Now().Unix(), 10)
		var p []interface{}
		pic := tgbotapi.InputMediaPhoto{
			Type:      "photo",
			Media:     url,
			Caption:   text.String(),
			ParseMode: tgbotapi.ModeMarkdown,
		}
		p = append(p, pic)
		msg := tgbotapi.MediaGroupConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: overall.Message.Chat.ID,
			},
			InputMedia: p,
		}
		channel.MessageChannel <- msg
		text.Reset()
	}
}

func Trend() {
	for update := range channel.TrendChannel {
		log.Println("开始绘图Trend")
		const Day = 86400
		adds := data2.GetAdds(7) //获取七天本地新增
		if adds == nil {
			errMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "请求数据错误")
			channel.MessageChannel <- errMsg
			return
		}
		var xRange, yRange []float64
		for _, v := range adds {
			s := strings.ReplaceAll(v.Date, ".", "")
			res := Time2TimeStamp(s)
			xRange = append(xRange, float64(res+Day))
			yRange = append(yRange, float64(v.LocalConfirmAdd))
		}
		uT := "2022." + adds[len(adds)-1].Date
		buf := Scatter(xRange, yRange, "7 Days Local Case Increment")
		if buf == nil {
			errMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "渲染错误")
			channel.MessageChannel <- errMsg
			return
		}
		fi := tgbotapi.FileBytes{
			Name:  "trend.jpg",
			Bytes: buf.Bytes(),
		}
		msg := tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat: tgbotapi.BaseChat{ChatID: update.Message.Chat.ID},
				File:     fi,
			},
			Caption: "七天内本土新增病例\n横轴代表日期 纵轴代表病例数\n数据更新时间" + uT,
		}
		channel.MessageChannel <- msg
	}
}

func Time2TimeStamp(t string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("20060102", "2022"+t, loc)
	return tt.Unix()
}
