package maker

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/data"
	"log"
	"strconv"
	"strings"
)

var size = 10

func RiskQuery() {
	for query := range channel.RiskQueryChannel {
		split := strings.Split(query.Data, "-")
		log.Println(split)
		page, _ := strconv.Atoi(split[2])
		text, markup := GetText(split[1], page)
		editedMsg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:      query.Message.Chat.ID,
				MessageID:   query.Message.MessageID,
				ReplyMarkup: &markup,
			},
			Text: text,
		}
		channel.MessageChannel <- editedMsg
	}
}

func GetText(level string, page int) (string, tgbotapi.InlineKeyboardMarkup) {
	var markup = tgbotapi.NewInlineKeyboardMarkup()
	var row []tgbotapi.InlineKeyboardButton
	text := strings.Builder{}
	switch level {
	case "return":
		var menu = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("高风险地区", "risk-2-1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("中风险地区", "risk-1-1"),
			),
		)
		markup = menu
		return "请选择区域", markup
	case "2":
		text.WriteString("高风险地区:")
	default:
		text.WriteString("中风险地区:")
	}
	areas := data.GetRiskLevel(level)
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("返回菜单", "risk-return-1"))
	maxPage := GetI(len(areas))
	if page > 1 {
		dateLeft := fmt.Sprintf("risk-%s-%d", level, page-1)
		row = append(row, tgbotapi.NewInlineKeyboardButtonData("👈上一页", dateLeft))
	}
	if page < maxPage {
		dateRight := fmt.Sprintf("risk-%s-%d", level, page+1)
		row = append(row, tgbotapi.NewInlineKeyboardButtonData("👉下一页", dateRight))
	}
	for k, v := range areas {
		// 去掉重复项
		if k < (page-1)*size || ((k+1)-(page-1)*size) > size {
			continue
		}
		text.WriteString("\n" + strconv.Itoa(k+1) + v.Area)
	}
	markup = tgbotapi.NewInlineKeyboardMarkup(row)
	return text.String(), markup
}

// GetI 获取最大页码
func GetI(len int) int {
	i := len / 10
	// 获取个位数
	for len > 9 {
		len = len % 10
	}
	if len > 0 {
		i = i + 1
	}
	return i
}
