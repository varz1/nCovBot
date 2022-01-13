package channel

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	UpdateChannel         chan tgbotapi.Update
	MessageChannel        chan tgbotapi.Chattable      // 普通消息Channel
	ListQueryChannel      chan *tgbotapi.CallbackQuery // 地区列表CallbackChannel
	ProvinceQueryChannel  chan *tgbotapi.CallbackQuery // 省份数据CallbackChannel
	OverallUpdateChannel  chan *tgbotapi.Update        // 概览UpdateChannel
	WorldUpdateChannel    chan *tgbotapi.Update        // 世界概览UpdateChannel
	ProvinceUpdateChannel chan *tgbotapi.Update        // 地区数据UpdateChannel
	NewsUpdateChannel     chan *tgbotapi.Update        // 新闻数据UpdateChannel
	RiskQueryChannel      chan *tgbotapi.CallbackQuery // 风险地区Callback
	TrendChannel          chan *tgbotapi.Update        // 风险地区Callback
	//NumChannel chan string
)

func init() {
	UpdateChannel = make(chan tgbotapi.Update)
	MessageChannel = make(chan tgbotapi.Chattable, 100)
	ListQueryChannel = make(chan *tgbotapi.CallbackQuery)
	ProvinceQueryChannel = make(chan *tgbotapi.CallbackQuery)
	OverallUpdateChannel = make(chan *tgbotapi.Update)
	WorldUpdateChannel = make(chan *tgbotapi.Update)
	ProvinceUpdateChannel = make(chan *tgbotapi.Update)
	NewsUpdateChannel = make(chan *tgbotapi.Update)
	RiskQueryChannel = make(chan *tgbotapi.CallbackQuery)
	TrendChannel = make(chan *tgbotapi.Update)
	//NumChannel = make(chan string)
}
