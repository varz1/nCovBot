package maker

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/data"
	"github.com/varz1/nCovBot/model"
	"strconv"
	"strings"
	"time"
)


func News() {
	for newsUpdate := range channel.NewsUpdateChannel {
		text := strings.Builder{}
		data2,exist := data.C.Get("news") // 请求API
		if !exist {
			text.WriteString("获取新闻数据失败")
			msg := tgbotapi.NewMessage(newsUpdate.Message.Chat.ID, text.String())
			channel.MessageChannel <- msg
			text.Reset()
			return
		}
		data3 := data2.(model.News)
		data1 := data3.Results
		//data1 := data.NewsData // 请求API
		var row1 []tgbotapi.InlineKeyboardButton
		var row2 []tgbotapi.InlineKeyboardButton
		for k, newsData := range data1 {
			date, _ := strconv.ParseInt(newsData.PubDate, 10, 64)
			tm := time.Unix(date/1000, 0).Format("2006-01-02 15:04")
			text.WriteString("\n" + strconv.Itoa(k+1) + " " + newsData.Title)
			text.WriteString("\n发布时间:" + tm)
			if k <= 4 {
				row1 = append(row1, tgbotapi.NewInlineKeyboardButtonURL(strconv.Itoa(k+1), newsData.SourceUrl))
			}
			if k > 4 {
				row2 = append(row2, tgbotapi.NewInlineKeyboardButtonURL(strconv.Itoa(k+1), newsData.SourceUrl))
			}
		}
		var de = tgbotapi.NewInlineKeyboardMarkup(row1, row2)
		msg := tgbotapi.NewMessage(newsUpdate.Message.Chat.ID, text.String())
		msg.ReplyMarkup = de
		channel.MessageChannel <- msg
		text.Reset()
	}
}
