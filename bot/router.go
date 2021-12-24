package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/spf13/viper"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/maker"
	"log"
	"strings"
)

func baseRouter(update *tgbotapi.Update) {
	message := update.Message.Text
	if update.Message.IsCommand() {
		go commandRouter(update)
		return
	}
	//只回应管理员消息
	if update.Message.Chat.ID != viper.GetInt64("ChatID") {
		return
	}
	switch message {
	case "hi":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi✋ :) Administrator")
		channel.MessageChannel <- msg
		return
	}
	if maker.IsContain(message) {
		channel.ProvinceUpdateChannel <- update
	}else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "无该地区数据或者输入错误")
		channel.MessageChannel <- msg
	}
}

func commandRouter(update *tgbotapi.Update) {
	message := update.Message.Text
	switch message {
	case "/start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"欢迎使用nCov疫情数据机器人🤖\n"+
				"功能列表:\n/start:使用提示\n/list:支持查询的地区列表\n/overall:查看疫情数据概览\n/news:查看最新新闻\n"+
				"\n使用Tip:\n发送列表中地区名可返回该地区疫情数据（注意格式）\n"+
				"示例消息:上海市\n"+
				"\n数据来自丁香园 本Bot不对数据负责")
		channel.MessageChannel <- msg
	case "/list":
		var menu = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("国内各省市", "list-province"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("国内外各国家地区", "list-country-1"),
			),
		)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,"请选择区域")
		msg.ReplyMarkup = menu
		channel.MessageChannel <- msg
	case "/overall":
		channel.OverallUpdateChannel <- update
	case "/news":
		channel.NewsUpdateChannel <- update
	}
}

func callBackRouter(query *tgbotapi.CallbackQuery) {
	commandData := strings.Fields(query.Data)
	log.Println(commandData[0])
	// 查看国家列表handler
	if strings.Contains(commandData[0], "list") {
		channel.ListQueryChannel <- query
	}
	if strings.Contains(commandData[0],"area"){
		channel.ProvinceQueryChannel <- query
	}
}
