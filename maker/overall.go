package maker

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	data2 "github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/model"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	SCATTER = model.Chartt{}
	Pie     = model.Chartt{}
)

func init() {
	GetScatter()
	GetPie()
}

func Overall() {
	text := strings.Builder{}
	for overall := range channel.OverallUpdateChannel {
		data := data2.GetOverall() //
		tm := time.Unix(data.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		text.WriteString("🇨🇳中国疫情概况:")
		text.WriteString("\n现存确诊(含港澳台):" + strconv.Itoa(data.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(data.CurrentConfirmedIncr))
		text.WriteString("\n现存无症状:" + strconv.Itoa(data.SeriousCount) + " ⬆️" + strconv.Itoa(data.SeriousIncr))
		text.WriteString("\n境外输入:" + strconv.Itoa(data.SuspectedCount) + " ⬆️" + strconv.Itoa(data.SuspectedIncr))
		text.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount) + " ⬆️" + strconv.Itoa(data.ConfirmedIncr))
		text.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount) + " ⬆️" + strconv.Itoa(data.CuredIncr))
		text.WriteString("\n累计死亡" + strconv.Itoa(data.DeadCount) + " ⬆️" + strconv.Itoa(data.DeadIncr))
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
		if SCATTER.Pie.Bytes() == nil {
			errMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "渲染错误")
			channel.MessageChannel <- errMsg
			return
		}
		msg := tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat: tgbotapi.BaseChat{ChatID: update.Message.Chat.ID},
				File: tgbotapi.FileBytes{
					Name:  "trend.jpg",
					Bytes: SCATTER.Pie.Bytes(),
				},
			},
			Caption: "七天内本土新增病例\n横轴代表日期 纵轴代表病例数\n图表更新时间:"+SCATTER.Date,
		}
		channel.MessageChannel <- msg
	}
}

func WorldOverall() {
	for update := range channel.WorldUpdateChannel {
		data := data2.GetOverall()
		global := data.GlobalStatistics
		if Pie.Pie.Bytes() == nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "图表为空")
			channel.MessageChannel <- msg
			return
		}
		caption := strings.Builder{}
		tm := time.Unix(data.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		caption.WriteString("\n🌏全球疫情概况")
		caption.WriteString("\n全球现存确诊" + strconv.Itoa(global.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(global.CurrentConfirmedIncr))
		caption.WriteString("\n全球累计确诊" + strconv.Itoa(global.ConfirmedCount) + " ⬆️" + strconv.Itoa(global.ConfirmedIncr))
		caption.WriteString("\n全球累计治愈" + strconv.Itoa(global.CuredCount) + " ⬆️" + strconv.Itoa(global.CuredIncr))
		caption.WriteString("\n全球累计死亡" + strconv.Itoa(global.DeadCount) + " ⬆️" + strconv.Itoa(global.DeadIncr))
		caption.WriteString("\n数据更新时间:" + tm)
		caption.WriteString("\n图表为各大洲累计病例数占比 更新时间:"+Pie.Date)
		p := tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat: tgbotapi.BaseChat{
					ChatID: update.Message.Chat.ID,
				},
				File: tgbotapi.FileBytes{
					Name:  "world.jpg",
					Bytes: Pie.Pie.Bytes(),
				},
			},
			Caption: caption.String(),
		}
		channel.MessageChannel <- p
		caption.Reset()
	}
}

func Time2TimeStamp(t string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("20060102", "2022"+t, loc)
	return tt.Unix()
}
