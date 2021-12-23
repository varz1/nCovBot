package channel

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	MessageChannel        chan tgbotapi.Chattable      // 普通消息Channel
	ListQueryChannel      chan *tgbotapi.CallbackQuery // 地区列表CallbackChannel
	OverallUpdateChannel  chan *tgbotapi.Update        // 概览UpdateChannel
	ProvinceUpdateChannel chan *tgbotapi.Update        // 地区数据UpdateChannel
	NewsUpdateChannel     chan *tgbotapi.Update        // 新闻数据UpdateChannel
)

func init() {
	MessageChannel = make(chan tgbotapi.Chattable, 100)
	ListQueryChannel = make(chan *tgbotapi.CallbackQuery)
	OverallUpdateChannel = make(chan *tgbotapi.Update)
	ProvinceUpdateChannel = make(chan *tgbotapi.Update)
	NewsUpdateChannel = make(chan *tgbotapi.Update)
}
