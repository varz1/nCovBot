package bot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gofiber/fiber/v2"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/maker"
	"github.com/varz1/nCovBot/model"
	"github.com/varz1/nCovBot/variables"
	"strconv"
	"strings"
)

func SetUpRouter(app *fiber.App) {
	app.Post("/"+variables.EnvToken, WebHookHandler)
	app.Get("/", BlogHandler)
	app.Get("/hi", HiHandler)
	app.Use(NotFoundHandler)
}

func baseRouter(update *tgbotapi.Update) {
	message := update.Message.Text
	admin, _ := strconv.Atoi(variables.EnvAdminId)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "无该地区或格式错误")
	if update.Message.IsCommand() {
		go commandRouter(update)
		return
	}
	if maker.IsContain(message) {
		msg = tgbotapi.NewMessage(int64(admin), "正在查找...")
		channel.MessageChannel <- msg
		channel.ProvinceUpdateChannel <- update
		return
	}
	// 管理员消息
	if update.Message.Chat.ID == int64(admin) {
		switch message {
		case "hi":
			msg = tgbotapi.NewMessage(int64(admin), "Hi👋 :) Administrator")
		case "/map":
			msg = tgbotapi.NewMessage(int64(admin), "开始更新图表数据...")
			channel.MessageChannel <- msg
			maker.GetChMap()
			maker.GetScatter()
			maker.GetPie()
			msg1 := tgbotapi.NewMessage(int64(admin), "图表数据更新完毕")
			channel.MessageChannel <- msg1
			return
		case "/data":
			msg = tgbotapi.NewMessage(int64(admin), "开始更新数据...")
			channel.MessageChannel <- msg
			data.GetNews()
			data.GetRiskLevel()
			data.GetOverall()
			data.GetWorld()
			msg1 := tgbotapi.NewMessage(int64(admin), "数据更新完毕")
			channel.MessageChannel <- msg1
			return
		}
	} else {
		notice := tgbotapi.NewMessage(int64(admin), fmt.Sprintf("User:%v\nId:%d", update.Message.Chat.UserName, update.Message.Chat.ID))
		channel.MessageChannel <- notice
	}
	channel.MessageChannel <- msg
}

func commandRouter(update *tgbotapi.Update) {
	message := update.Message.Text
	switch message {
	case "/start":
		var msg tgbotapi.MessageConfig
		if strconv.Itoa(int(update.Message.Chat.ID)) == variables.EnvAdminId {
			msg = GetStartMenu(*update, true)
		} else {
			msg = GetStartMenu(*update, false)
		}
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

func GetStartMenu(update tgbotapi.Update, isAdmin bool) tgbotapi.MessageConfig {
	var msg tgbotapi.MessageConfig
	if isAdmin {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID,
			"你好管理员！欢迎使用nCov疫情数据机器人🤖\n"+
				"功能列表:\n/start:使用提示👋\n/list:支持查询的地区列表🌏\n/overall:查看中国疫情数据概览😷\n"+
				"/world:查看世界疫情概览🌎\n/trend:查看本土疫情趋势图📶\n"+
				"/news:查看最新新闻🆕\n"+
				"/risk:中高风险地区列表⚠️\n"+
				"/map:更新图表\n"+
				"/data:更新数据\n"+
				"\n使用Tip:\n发送列表中地区名可返回该地区疫情数据（注意格式）\n"+
				"示例消息:上海市\n"+
				"\n数据来自丁香园/腾讯/百度 本Bot不对数据负责")
	} else {
		msg = tgbotapi.NewMessage(update.Message.Chat.ID,
			"欢迎使用nCov疫情数据机器人🤖\n"+
				"功能列表:\n/start:使用提示👋\n/list:支持查询的地区列表🌏\n/overall:查看中国疫情数据概览😷\n"+
				"/world:查看世界疫情概览🌎\n/trend:查看本土疫情趋势图📶\n"+
				"/news:查看最新新闻🆕\n"+
				"/risk:中高风险地区列表⚠️\n"+
				"\n使用Tip:\n发送列表中地区名可返回该地区疫情数据（注意格式）\n"+
				"示例消息:上海市\n"+
				"\n数据来自丁香园/腾讯/百度 本Bot不对数据负责")
	}
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
