package bot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/json-iterator/go"
	"github.com/spf13/viper"
	"github.com/varz1/nCovBot/data"
	"strings"
)

func baseRouter(update tgbotapi.Update) {
	//只回应管理员消息
	if update.Message.Chat.ID != viper.GetInt64("ChatID") {
		return
	}
	switch update.Message.Text {
	case "hi":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hi administrator")
		_, err := botAPI.Send(msg)
		if err != nil {
			panic("base wrong")
		}
	}
}

func oderRouter(update tgbotapi.Update) {
	switch update.Message.Command() {
	case "list":
		result, _ := data.GetAreas()
		text, _ := jsoniter.ConfigFastest.Marshal(result.Areas)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			strings.Replace(strings.Trim(fmt.Sprint(string(text)),"[]")," ",",",-1))
		_, err := botAPI.Send(msg)
		if err != nil {
			panic("order wrong")
		}
	}
}

func boardRouter() {

}