package maker

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	data2 "github.com/varz1/nCovBot/data"
	"strconv"
	"strings"
	"time"
)

func Province() {
	for provinceUpdate := range channel.ProvinceUpdateChannel {
		text := strings.Builder{}
		data := data2.GetAreaData(provinceUpdate.Message.Text)
		cities := data.Cities
		tm := time.Unix(data.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		if data.IsEmpty() {
			text.WriteString("请求错误")
		} else if cities != nil {
			text.WriteString(data.CountryName + data.ProvinceName + "疫情数据:")
			text.WriteString("\n现存确诊（含港澳台以及✈️境外输入）:" + strconv.Itoa(data.CurrentConfirmedCount))
			text.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount))
			text.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount))
			text.WriteString("\n累计死亡:" + strconv.Itoa(data.DeadCount) + "\n")
			for _, v := range cities {
				text.WriteString("\n" + v.CityName + "数据:")
				text.WriteString("\n现存确诊:" + v.CurrentConfirmedCountStr)
				text.WriteString("\n累计确诊:" + strconv.Itoa(v.ConfirmedCount))
				text.WriteString("\n累计治愈:" + strconv.Itoa(v.CuredCount))
				text.WriteString("\n累计死亡:" + strconv.Itoa(v.DeadCount))
				text.WriteString("\n高风险地区数量:" + strconv.Itoa(v.HighDangerCount))
				text.WriteString("\n中风险地区数量:" + strconv.Itoa(v.MidDangerCount) + "\n")
			}
		} else {
			text.WriteString(data.CountryName + "疫情数据:")
			text.WriteString("\n现存确诊:" + strconv.Itoa(data.CurrentConfirmedCount))
			text.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount))
			text.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount))
			text.WriteString("\n累计死亡:" + strconv.Itoa(data.DeadCount))
			text.WriteString("\n暂无更多城市数据")
		}
		text.WriteString("\n数据更新时间:" + tm)
		msg := tgbotapi.NewMessage(provinceUpdate.Message.Chat.ID, text.String())
		channel.MessageChannel <- msg
		text.Reset()
	}
}
