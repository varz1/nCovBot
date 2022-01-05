package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"testing"
)

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("国内各省市", "province"),
	),
)

func Test1bot(t *testing.T) {
	bot, err := tgbotapi.NewBotAPI("5053020953:AAHGrarAJJo0Xt9kn4-AhmDKs6dWFgKZZF8")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			switch update.Message.Text {
			case "open":
				msg.ReplyMarkup = numericKeyboard
			case "close":
				msg.ReplyMarkup = nil
			case "photo":

			}
			if _, err = bot.Send(msg); err != nil {
				panic(err)
			}
		} else if update.CallbackQuery != nil {
			query := *update.CallbackQuery
			//callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			//if _, err := bot.Request(callback); err != nil {
			//	panic(err)
			//}
			var c tgbotapi.Chattable
			msg := tgbotapi.NewEditMessageText(query.Message.Chat.ID,
				query.Message.MessageID, "query.Data")
			msg.ReplyMarkup = &numericKeyboard
			c = msg
			_, _ = botAPI.Send(c)
			//msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			//if _, err := bot.Send(msg); err != nil {
			//	panic(err)
			//}
		}
	}
}
