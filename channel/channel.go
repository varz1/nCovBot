package channel

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/model"
)

var (
	MessageChannel     chan tgbotapi.Chattable      // 普通text消息Channel
	ListQueryChannel   chan *tgbotapi.CallbackQuery // 地区列表Channel
	OverallUpdateChannel  chan *tgbotapi.Update    // 概览Channel With MsgConfig
	ProvinceUpdateChannel chan *tgbotapi.Update       //地区数据Channel With MsgConfig
	NewsMsgChannel     chan model.NewsMsg
)

func init() {
	MessageChannel = make(chan tgbotapi.Chattable, 100)
	ListQueryChannel = make(chan *tgbotapi.CallbackQuery, 10)
	OverallUpdateChannel = make(chan *tgbotapi.Update, 10)
	ProvinceUpdateChannel = make(chan *tgbotapi.Update)
	NewsMsgChannel = make(chan model.NewsMsg)
}
