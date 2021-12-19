package channel

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/model"
)

var (
	MessageChannel chan tgbotapi.Chattable
	ListChannel chan model.Areas
)

func init() {
	MessageChannel = make(chan tgbotapi.Chattable, 100)
	ListChannel = make(chan model.Areas)
}
