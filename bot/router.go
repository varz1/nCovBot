package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/model"
)

func baseRouter(update *tgbotapi.Update) {
	//只回应管理员消息
	message := update.Message.Text
	if update.Message.Chat.ID != viper.GetInt64("ChatID") {
		return
	}
	switch message {
	case "hi":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi✋ :) Administrator")
		channel.MessageChannel <- msg
	case "test":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "测试频道")
		channel.MessageChannel <- msg
	}
}

func commandRouter(update *tgbotapi.Update) {
	message := update.Message.Text
	switch message {
	case "/start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "欢迎使用")
		channel.MessageChannel <- msg
	case "/list":
		var list model.Areas
		list.Types = "menu"
		list.AreaMessage = *update.Message
		channel.ListChannel <- list
	case "/overall":
		var msg model.OverallMessage
		msg.OverallData = data.Overall()
		msg.Overall.ChatID = update.Message.Chat.ID
		channel.OverallMsgChannel <- msg
	}
}
func callBackRouter(query *tgbotapi.CallbackQuery) {
	// 查看国家列表handler
	if query.Data == "province" || query.Data == "country" {
		var list model.Areas
		list.AreaMessage = *query.Message
		list.Types = query.Data
		channel.ListChannel <- list
	}
}
