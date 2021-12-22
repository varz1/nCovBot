package maker

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"strconv"
	"strings"
	"time"
)

func News() {
	for news := range channel.NewsMsgChannel {
		text := strings.Builder{}
		config := news.Config
		data := news.Data
		var row1 []tgbotapi.InlineKeyboardButton
		var row2 []tgbotapi.InlineKeyboardButton
		for k, newsData := range data {
			timeTem := "2006-01-02 15:04:05"
			date, _ := strconv.ParseInt(newsData.PubDate, 10, 64)
			tm := time.Unix(date/1000, 0).Format(timeTem)
			text.WriteString("\n" + strconv.Itoa(k+1) + " " + newsData.Title)
			text.WriteString("\n发布时间:" + tm)
			if k <= 4 {
				row1 = append(row1, tgbotapi.NewInlineKeyboardButtonURL(strconv.Itoa(k+1), newsData.SourceUrl))
			}
			if k > 4 {
				row2 = append(row2, tgbotapi.NewInlineKeyboardButtonURL(strconv.Itoa(k+1), newsData.SourceUrl))
			}
			config.Text = text.String()
		}
		var de = tgbotapi.NewInlineKeyboardMarkup(row1, row2)
		config.ReplyMarkup = de
		channel.MessageChannel <- config
		text.Reset()
	}
}
