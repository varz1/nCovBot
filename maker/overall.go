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

func init() {
	err := data2.GetChMap()
	if err != nil {
		return 
	}
}
func Overall() {
	text := strings.Builder{}
	for overall := range channel.OverallUpdateChannel {
		data := data2.GetOverall()
		//global := data.GlobalStatistics
		mapTime, err := data2.GetState(0)
		if err != nil {
			return
		}
		tm := time.Unix(data.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		tm1 := time.Unix(mapTime, 0).Format("2006-01-02 15:04")
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
		trendTime, err := data2.GetState(1)
		if err != nil {
			return
		}
		tm := time.Unix(trendTime, 0).Format("2006-01-02 15:04")
		text := "本土疫情趋势图" + "\n图表更新时间:" + tm
		// 时间戳更新地图
		var url = os.Getenv("baseURL") + "virusTrend.png" + "?a=" + strconv.FormatInt(time.Now().Unix(), 10)
		var p []interface{}
		pic := tgbotapi.InputMediaPhoto{
			Type:      "photo",
			Media:     url,
			Caption:   text,
			ParseMode: tgbotapi.ModeMarkdown,
		}
		p = append(p, pic)
		msg := tgbotapi.MediaGroupConfig{
			BaseChat: tgbotapi.BaseChat{
				ChatID: update.Message.Chat.ID,
			},
			InputMedia: p,
		}
		channel.MessageChannel <- msg
	}
}
