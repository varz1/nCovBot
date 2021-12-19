package service

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/data"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("国内各省市", "province"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("国外国家及地区", "country"),
	),
)

func ListMk(update tgbotapi.Update, pro bool) error {
	var err error
	result, err := data.GetAreas()
	if pro {
		text := data.GetList(pro, result)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyMarkup = numericKeyboard

		channel.MessageChannel <- msg
	}
	if !pro {
		text := data.GetList(pro, result)
		msg := tgbotapi.NewEditMessageText(update.Message.Chat.ID,update.Message.MessageID, text)

		channel.MessageChannel <- msg
	}

	if err != nil {
		return err
	}
	return err
}
