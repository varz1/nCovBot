package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"log"
)

var botAPI *tgbotapi.BotAPI

func initBot() *tgbotapi.BotAPI {
	api, err := tgbotapi.NewBotAPI(viper.GetString("TOKEN"))
	if err != nil {
		log.Println("初始化失败 检查token")
		log.Panic(err)
	}
	return api
}

func Run() {
	botAPI = initBot()
	botAPI.Debug = true
	go sender()
	go receiver()
}

func receiver() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := botAPI.GetUpdatesChan(u)
	if err != nil {
		panic("update wrong")
	}
	for update := range updates {
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
