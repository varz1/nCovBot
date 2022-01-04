package maker

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	data2 "github.com/varz1/nCovBot/data"
	"os"
	"strconv"
	"strings"
	"time"
)

func Overall() {
	text := strings.Builder{}
	for overall := range channel.OverallUpdateChannel {
		data := data2.GetOverall()
		//global := data.GlobalStatistics
		tm := time.Unix(data.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		text.WriteString("🇨🇳中国疫情概况:")
		text.WriteString("\n现存确诊(含港澳台):" + strconv.Itoa(data.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(data.CurrentConfirmedIncr))
		text.WriteString("\n现存无症状:" + strconv.Itoa(data.SeriousCount) + " ⬆️" + strconv.Itoa(data.SeriousIncr))
		text.WriteString("\n境外输入:" + strconv.Itoa(data.SuspectedCount) + " ⬆️" + strconv.Itoa(data.SuspectedIncr))
		text.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount) + " ⬆️" + strconv.Itoa(data.ConfirmedIncr))
		text.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount) + " ⬆️" + strconv.Itoa(data.CuredIncr))
		text.WriteString("\n累计死亡" + strconv.Itoa(data.DeadCount) + " ⬆️" + strconv.Itoa(data.DeadIncr))
		//text.WriteString("\n🌏全球疫情概况")
		//text.WriteString("\n全球现存确诊" + strconv.Itoa(global.CurrentConfirmedCount) + " ⬆️" + strconv.Itoa(global.CurrentConfirmedIncr))
		//text.WriteString("\n全球累计确诊" + strconv.Itoa(global.ConfirmedCount) + " ⬆️" + strconv.Itoa(global.ConfirmedIncr))
		//text.WriteString("\n全球累计治愈" + strconv.Itoa(global.CuredCount) + " ⬆️" + strconv.Itoa(global.CuredIncr))
		//text.WriteString("\n全球累计死亡" + strconv.Itoa(global.DeadCount) + " ⬆️" + strconv.Itoa(global.DeadIncr))
		text.WriteString("\n数据更新时间:" + tm)
		//msg := tgbotapi.NewMessage(overall.Message.Chat.ID,text.String())
		var url = os.Getenv("baseURL") + "virus-map.png"
		//var url = "https://i.pinimg.com/736x/33/32/6d/33326dcddbf15c56d631e374b62338dc.jpg"
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
