package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"log"
	"os"
)

var botAPI *tgbotapi.BotAPI

func initBot() *tgbotapi.BotAPI {
	api, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Println("初始化失败 检查token")
		log.Panic(err)
	}
	//初始化webHook
	url := os.Getenv("baseURL") + api.Token
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
