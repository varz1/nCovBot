package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/maker"
	"github.com/varz1/nCovBot/model"
)

func baseRouter(update *tgbotapi.Update) {
	if update.Message.IsCommand() {
		go commandRouter(update)
		return
	}
	//只回应管理员消息
	message := update.Message.Text
	if update.Message.Chat.ID != viper.GetInt64("ChatID") {
		return
	}
	switch message {
	case "hi":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi✋ :) Administrator")
		channel.MessageChannel <- msg
	case "test":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "测试频道")
		channel.MessageChannel <- msg
	}
	if !maker.IsContain(message) {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "无该地区数据或者输入错误")
		channel.MessageChannel <- msg
	}
	if maker.IsContain(message) {
		var area model.ProvinceMsg
		area.Data = data.AreaData(message)
		area.Config.ChatID = update.Message.Chat.ID
		channel.ProvinceMsgChannel <- area
	}
}

func commandRouter(update *tgbotapi.Update) {
	if update.Message.Chat.ID != viper.GetInt64("ChatID") {
		return
	}
	message := update.Message.Text
	switch message {
	case "/start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"欢迎使用nCov疫情数据机器人🤖\n" +
			"功能列表:\n/start:使用提示\n/list:支持查询的地区列表\n/overall:查看疫情数据概览\n" +
			"\n发送列表中地区名可返回该地区疫情数据（注意格式）\n" +
			"示例消息:上海市")
		channel.MessageChannel <- msg
	case "/list":
		var list model.Areas
		list.Types = "menu"
		list.AreaMessage = *update.Message
		channel.ListChannel <- list
	case "/overall":
		var msg model.OverallMessage
		msg.OverallData = data.Overall()
		msg.Overall.ChatID = update.Message.Chat.ID
		channel.OverallMsgChannel <- msg
	}
}
func callBackRouter(query *tgbotapi.CallbackQuery) {
	// 查看国家列表handler
	if query.Data == "province" || query.Data == "country" {
		var list model.Areas
		list.AreaMessage = *query.Message
		list.Types = query.Data
		channel.ListChannel <- list
	}
}
