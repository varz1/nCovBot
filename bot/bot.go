package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/variables"
	"log"
)

var botAPI *tgbotapi.BotAPI

func initBot() *tgbotapi.BotAPI {
	api, err := tgbotapi.NewBotAPI(variables.EnvToken)
	if err != nil {
		log.Panic("初始化失败 检查token", err)
	}
	//初始化webHook
	url := variables.EnvBaseUrl + api.Token
	_, err = api.SetWebhook(tgbotapi.NewWebhook(url))
	return api
}

func Run() {
	botAPI = initBot()
	botAPI.Debug = false
	go sender()
	go receiver()
}

func receiver() {
	for update := range channel.UpdateChannel {
		if update.CallbackQuery != nil {
			go callBackRouter(update.CallbackQuery)
			continue
		}
		if update.Message == nil {
			continue
		}
		go baseRouter(&update)
	}
}
