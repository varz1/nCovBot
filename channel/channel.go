package channel

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/model"
)

var (
	MessageChannel chan tgbotapi.Chattable // 普通text消息Channel
	ListChannel    chan model.Areas        // 地区列表Channel
	OverallMsgChannel chan model.OverallMessage
)

func init() {
	MessageChannel = make(chan tgbotapi.Chattable, 100)
	ListChannel = make(chan model.Areas, 10)
	OverallMsgChannel = make(chan model.OverallMessage, 10)
}
