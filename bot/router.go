package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/model"
)

var board = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("国内各省市", "province"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("国外各国家地区", "country"),
	),
)
func baseRouter(update tgbotapi.Update) {
	//只回应管理员消息
	if update.Message.Chat.ID != viper.GetInt64("ChatID") {
		return
	}
	switch update.Message.Text {
	case "hi":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hi✋ administrator")
		channel.MessageChannel <- msg
	case "open":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "请选择区域")
		msg.ReplyMarkup = board
		channel.MessageChannel <- msg
	}
}

func callBackRouter(query *tgbotapi.CallbackQuery) {
	if query.Data == "province" || query.Data == "country" {
		var list model.Areas
		list.Query = *query.Message
		list.Types = query.Data
		channel.ListChannel <- list
	}
}
