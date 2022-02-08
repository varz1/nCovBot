package maker

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	data2 "github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/model"
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
		var msg = tgbotapi.MessageConfig{}
		//if data.IsEmpty() {
		//	text.WriteString("请求错误")
		//}
		if len(cities) == 0 {
			cities = nil
		}
		if cities == nil {
			text.WriteString(data.ProvinceName + "疫情数据:")
			text.WriteString("\n现存确诊:" + strconv.Itoa(data.CurrentConfirmedCount))
			text.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount))
			text.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount))
			text.WriteString("\n累计死亡:" + strconv.Itoa(data.DeadCount))
			text.WriteString("\n暂无更多地区数据")
		}
		if cities != nil {
			text.WriteString(data.CountryName + data.ProvinceName + "疫情数据:")
			text.WriteString("\n现存确诊（含港澳台以及✈️境外输入）:" + strconv.Itoa(data.CurrentConfirmedCount))
			text.WriteString("\n累计确诊:" + strconv.Itoa(data.ConfirmedCount))
			text.WriteString("\n累计治愈:" + strconv.Itoa(data.CuredCount))
			text.WriteString("\n累计死亡:" + strconv.Itoa(data.DeadCount))
			msg.ReplyMarkup = GetMarkup(data.ProvinceName, "false", cities)
		}
		text.WriteString("\n数据更新时间:" + tm)
		msg.Text = text.String()
		msg.ChatID = provinceUpdate.Message.Chat.ID
		channel.MessageChannel <- msg
		text.Reset()
	}
}

// GetMarkup 获取对应的keyboard
func GetMarkup(provinceName string, isCity string, cities []model.Cities) tgbotapi.InlineKeyboardMarkup {
	var markup = tgbotapi.NewInlineKeyboardMarkup()
	var row1 []tgbotapi.InlineKeyboardButton
	var row2 []tgbotapi.InlineKeyboardButton
	var row3 []tgbotapi.InlineKeyboardButton
	var row4 []tgbotapi.InlineKeyboardButton
	var row5 []tgbotapi.InlineKeyboardButton
	var row6 []tgbotapi.InlineKeyboardButton
	var row7 []tgbotapi.InlineKeyboardButton
	var row8 []tgbotapi.InlineKeyboardButton
	var row9 []tgbotapi.InlineKeyboardButton
	for k, v := range cities {
		butData := fmt.Sprintf("area-%d-%s-%s", k, provinceName, v.CityName)
		if isCity == v.CityName {
			butData = fmt.Sprintf("area-%d-%s-%s", k, provinceName, provinceName)
			v.CityName = provinceName
		}
		if k < 4 {
			row1 = append(row1, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1)
		}
		if k >= 4 && k < 9 {
			row2 = append(row2, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1, row2)
		}
		if k >= 9 && k < 14 {
			row3 = append(row3, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3)
		}
		if k >= 14 && k < 19 {
			row4 = append(row4, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4)
		}
		if k >= 19 && k < 24 {
			row5 = append(row5, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4, row5)
		}
		if k >= 24 && k < 29 {
			row6 = append(row6, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4, row5, row6)
		}
		if k >= 29 && k < 34 {
			row7 = append(row7, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4, row5, row6, row7)
		}
		if k >= 34 && k < 39 {
			row8 = append(row8, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4, row5, row6, row7, row8)
		}
		if k >= 39 && k < 44 {
			row9 = append(row9, tgbotapi.NewInlineKeyboardButtonData(v.CityName, butData))
			markup = tgbotapi.NewInlineKeyboardMarkup(row1, row2, row3, row4, row5, row6, row7, row8, row9)
		}
		//log.Println(len(row1) + len(row2) + len(row3) + len(row4) + len(row5) + len(row6) + len(row7) + len(row8) + len(row9))
	}
	return markup
}

// QueryProvince 选择具体区域的Callback处理
func QueryProvince() {
	for query := range channel.ProvinceQueryChannel {
		text := strings.Builder{}
		split := strings.Split(query.Data, "-")
		total := data2.GetAreaData(split[2])
		tm := time.Unix(total.UpdateTime/1000, 0).Format("2006-01-02 15:04")
		var markup tgbotapi.InlineKeyboardMarkup
		if split[2] == split[3] {
			text.WriteString(total.CountryName + total.ProvinceName + "疫情数据:")
			text.WriteString("\n现存确诊（含港澳台以及✈️境外输入）:" + strconv.Itoa(total.CurrentConfirmedCount))
			text.WriteString("\n累计确诊:" + strconv.Itoa(total.ConfirmedCount))
			text.WriteString("\n累计治愈:" + strconv.Itoa(total.CuredCount))
			text.WriteString("\n累计死亡:" + strconv.Itoa(total.DeadCount))
			text.WriteString("\n数据更新时间:" + tm)
			markup = GetMarkup(total.ProvinceName, "false", total.Cities)
		} else {
			k, _ := strconv.Atoi(split[1])
			city := total.Cities[k]
			text.WriteString("\n" + split[2] + city.CityName + "数据:")
			text.WriteString("\n现存确诊:" + city.CurrentConfirmedCountStr)
			text.WriteString("\n累计确诊:" + strconv.Itoa(city.ConfirmedCount))
			text.WriteString("\n累计治愈:" + strconv.Itoa(city.CuredCount))
			text.WriteString("\n累计死亡:" + strconv.Itoa(city.DeadCount))
			text.WriteString("\n高风险地区数量:" + strconv.Itoa(city.HighDangerCount))
			text.WriteString("\n中风险地区数量:" + strconv.Itoa(city.MidDangerCount))
			text.WriteString("\n数据更新时间:" + tm)
			markup = GetMarkup(split[2], split[3], total.Cities)
		}
		editedMsg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:      query.Message.Chat.ID,
				MessageID:   query.Message.MessageID,
				ReplyMarkup: &markup,
			},
			Text: text.String(),
		}
		channel.MessageChannel <- editedMsg
		text.Reset()
	}
}
