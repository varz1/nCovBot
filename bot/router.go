package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

func baseRouter(update tgbotapi.Update) {
	//只回应管理员消息
	if update.Message.Chat.ID != viper.GetInt64("ChatID") {
		return
	}
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	_, err := botAPI.Send(msg)
	if err != nil {
		panic("send wrong")
	}
}
