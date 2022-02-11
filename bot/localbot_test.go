package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/varz1/nCovBot/maker"
	"log"
	"testing"
)

func TestBot(t *testing.T) {
	go BotTestRun()
	go maker.List()
	go maker.Overall()
	go maker.WorldOverall()
	go maker.Trend()
	go maker.Province()
	go maker.QueryProvince()
	go maker.News()
	go maker.RiskQuery()
}

func BotTestRun() {
	botAPI = initTestBot()
	botAPI.Debug = false
	go sender()
	go rec()
}

func initTestBot() *tgbotapi.BotAPI {
	api, err := tgbotapi.NewBotAPI("5053020953:AAFV81i2e8lkjFdDGM7g2XskxiZRDKA3jVQ")
	if err != nil {
		log.Println("初始化失败 检查token")
		log.Panic(err)
	}
	return api
}

func rec() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := botAPI.GetUpdatesChan(u)

	if err != nil {
		logrus.WithField("func", "receiver").WithField("err in", "GetUpdatesChan").Panicln(err)
		return
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