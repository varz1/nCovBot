package channel

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var (
	MessageChannel chan tgbotapi.Chattable
)

func init()  {
	MessageChannel = make(chan tgbotapi.Chattable,100)
}
