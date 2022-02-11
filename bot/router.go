package bot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/maker"
	"github.com/varz1/nCovBot/model"
	"os"
	"strconv"
	"strings"
)

//func SetUpRouter(app *fiber.App) {
//	app.Post("/"+botAPI.Token, WebHookHandler)
//}

func baseRouter(update *tgbotapi.Update) {
	message := update.Message.Text
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "无该地区或格式错误")
	if update.Message.IsCommand() {
		go commandRouter(update)
		return
	}
	if maker.IsContain(message) {
		channel.ProvinceUpdateChannel <- update
		return
	}
	if strconv.Itoa(int(update.Message.Chat.ID)) != os.Getenv("AdminId") {
		id, _ := strconv.Atoi(os.Getenv("AdminId"))
		notice := tgbotapi.NewMessage(int64(id), fmt.Sprintf("User:%s\nId:%d",update.Message.Chat.UserName,update.Message.Chat.ID))
		channel.MessageChannel <- notice
	}
	// 管理员消息
	if strconv.Itoa(int(update.Message.Chat.ID)) == os.Getenv("AdminId") {
		switch message {
		case "hi":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hi👋 :) Administrator")
		case "update":
			maker.GetChMap()
			maker.GetScatter()
			maker.GetPie()
		}
	}
	channel.MessageChannel <- msg
}

func commandRouter(update *tgbotapi.Update) {
	message := update.Message.Text
	switch message {
	case "/start":
		msg := GetStartMenu(*update)
		channel.MessageChannel <- msg
	case "/list":
		msg := GetListMenu(*update)
		channel.MessageChannel <- msg
	case "/overall":
		channel.OverallUpdateChannel <- update
	case "/news":
		channel.NewsUpdateChannel <- update
	case "/risk":
		msg := GetRiskMenu(*update)
		channel.MessageChannel <- msg
	case "/trend":
		channel.TrendChannel <- update
	case "/world":
		channel.WorldUpdateChannel <- update
	}
}

func callBackRouter(query *tgbotapi.CallbackQuery) {
	commandData := strings.Split(query.Data, "-")
	switch commandData[0] {
	case "list":
		channel.ListQueryChannel <- query
	case "area":
		channel.ProvinceQueryChannel <- query
	case "risk":
		channel.RiskQueryChannel <- query
	}
}

func GetStartMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"欢迎使用nCov疫情数据机器人🤖\n"+
			"功能列表:\n/start:使用提示👋\n/list:支持查询的地区列表🌏\n/overall:查看中国疫情数据概览😷\n"+
			"/world:查看世界疫情概览🌎\n/trend:查看本土疫情趋势图📶\n"+
			"/news:查看最新新闻🆕\n"+
			"/risk:中高风险地区列表⚠️\n"+
			"\n使用Tip:\n发送列表中地区名可返回该地区疫情数据（注意格式）\n"+
			"示例消息:上海市\n"+
			"\n数据来自丁香园/腾讯/百度 本Bot不对数据负责")
	return msg
}

func GetListMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	var menu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("国内各省市", "list-province"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("国内外各国家地区", "list-country-1"),
		),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "请选择区域")
	msg.ReplyMarkup = menu
	return msg
}

func GetRiskMenu(update tgbotapi.Update) tgbotapi.MessageConfig {
	var riskdata model.Risks
	risk, _ := data.C.Get("risk")
	riskdata = risk.(model.Risks)
	var menu = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("高风险地区("+strconv.Itoa(len(riskdata.High))+"个)▶️", "risk-2-1"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("中风险地区("+strconv.Itoa(len(riskdata.Mid))+"个)▶️", "risk-1-1"),
		),
	)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "点击展开详细列表")
	msg.ReplyMarkup = menu
	return msg
}
