package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/service"
	"log"
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
	if update.Message.IsCommand() {
		var err error
		if update.Message.Command() == "list" {
			err = service.ListMk(update, true)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func callBackRouter(update tgbotapi.Update) {
	switch update.CallbackQuery.Data {
	case "province":
		_,_ = botAPI.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID,"当前已是"))
	case "country":
		err := service.ListMk(update, false)
		if err != nil {
			log.Println(err)
		}
	}
	return
}
