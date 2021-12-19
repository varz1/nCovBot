package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
)

var botAPI *tgbotapi.BotAPI

func initBot() *tgbotapi.BotAPI {
	if viper.GetString("TOKEN") != "" && viper.GetInt64("ChatID") != 0 {
		api, _ := tgbotapi.NewBotAPI(viper.GetString("TOKEN"))
		return api
	} else {
		panic("init wrong!")
	}
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
			go callBackRouter(update)
		}
		if update.Message == nil {
			continue
		}
		go baseRouter(update)
	}
}
