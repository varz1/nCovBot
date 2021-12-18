package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/data"
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
	}
}

func oderRouter(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "list":
		var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("1", "1"),
				tgbotapi.NewInlineKeyboardButtonData("2", "2"),
				tgbotapi.NewInlineKeyboardButtonData("3", "3"),
				tgbotapi.NewInlineKeyboardButtonData("4", "4"),
				tgbotapi.NewInlineKeyboardButtonData("5", "5"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("6", "6"),
				tgbotapi.NewInlineKeyboardButtonData("7", "7"),
				tgbotapi.NewInlineKeyboardButtonData("8", "8"),
				tgbotapi.NewInlineKeyboardButtonData("9", "9"),
				tgbotapi.NewInlineKeyboardButtonData("10", "10"),
			),
		)
		result, _ := data.GetAreas()
		text := data.GetList(1, result)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyMarkup = numericKeyboard
		_, err := botAPI.Send(msg)
		if err != nil {
			panic("order wrong")
		}
	}
}

//func queryRouter() (update tgbotapi.Update) {
//	switch update.CallbackQuery.Data {
//	case "1":
//
//	}
//}
