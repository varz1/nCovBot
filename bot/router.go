package bot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gofiber/fiber/v2"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/maker"
	"log"
	"os"
	"strconv"
	"strings"
)

//func SetUpRouter(app *fiber.App) {
//	app.Post("/"+botAPI.Token, WebHookHandler)
//}

func WebHookHandler(c *fiber.Ctx) error {
	u := new(tgbotapi.Update)
	err := c.BodyParser(&u)
	if err != nil {
		log.Println("req解析失败")
		return err
	}
	log.Printf("开始处理Update\nUpdateId:%v\nCallbackQuery:%v\nInlineQuery:%v\nMessage:%v\nEditedMessage:%v\n",
		u.UpdateID, u.CallbackQuery, u.InlineQuery, u.Message, u.EditedMessage)
	channel.UpdateChannel <- *u
	return nil
}

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
	// 管理员消息控制
	if strconv.Itoa(int(update.Message.Chat.ID)) == os.Getenv("AdminId") {
		switch message {
		case "hi":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Hi👋 :) Administrator")
			channel.MessageChannel <- msg
			return
		}
	}
	// 处理菜单
	split := strings.Split(message,"")
	switch split[0] {
	case "🌏":
		var menu = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("国内各省市", "list-province"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("国内外各国家地区", "list-country-1"),
			),
		)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "请选择区域")
		msg.ReplyMarkup = menu
		channel.MessageChannel <- msg
	case "😷":
		channel.OverallUpdateChannel <- update
	case "🆕":
		channel.NewsUpdateChannel <- update
	case "⚠️":
		var menu = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("高风险地区", "risk-2-1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("中风险地区", "risk-1-1"),
			),
		)
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "请选择区域")
		msg.ReplyMarkup = menu
		channel.MessageChannel <- msg
	case "🧒🏼":
		msg = tgbotapi.NewMessage(update.Message.Chat.ID,
			"欢迎使用nCov疫情数据机器人🤖\n"+
				"功能列表:\n/start:使用提示\n/list:支持查询的地区列表\n/overall:查看疫情数据概览\n/news:查看最新新闻\n"+
				"\n使用Tip:\n发送列表中地区名可返回该地区疫情数据（注意格式）\n"+
				"示例消息:上海市\n"+
				"\n数据来自丁香园 本Bot不对数据负责")
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
				"/risk:中高风险地区列表\n"+
				"\n使用Tip:\n发送列表中地区名可返回该地区疫情数据（注意格式）\n"+
				"示例消息:上海市\n"+
				"\n数据来自丁香园 本Bot不对数据负责")
		var numericKeyboard = tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("🌏支持地区"),
				tgbotapi.NewKeyboardButton("😷疫情概览"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("🆕最新新闻"),
				tgbotapi.NewKeyboardButton("⚠️查看风险地区"),
				tgbotapi.NewKeyboardButton("🧒🏼帮助"),
			),
		)
		msg.ReplyMarkup = numericKeyboard
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
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "请选择区域")
		msg.ReplyMarkup = menu
		channel.MessageChannel <- msg
	case "/overall":
		channel.OverallUpdateChannel <- update
	case "/news":
		channel.NewsUpdateChannel <- update
	case "/risk":
		var menu = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("高风险地区", "risk-2-1"),
			),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("中风险地区", "risk-1-1"),
			),
		)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "请选择区域")
		msg.ReplyMarkup = menu
		channel.MessageChannel <- msg
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
