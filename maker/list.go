package maker

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/varz1/nCovBot/channel"
	"io/ioutil"
	"os"
	"strings"
)
type Jsons struct {
	Results []string `json:"results"`
}

func List() {
	log := logrus.WithField("func","ListQueryChannel")
	log.Info("打开文件")
	post, err := GetData()
	if err != nil {
		log.Errorln("请求API错误",err)
		return
	}
	text := ""
	var board = tgbotapi.NewInlineKeyboardMarkup()
	for query := range channel.ListQueryChannel {
		switch query.Data {
		case "list-province":
			board = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("国内外各国家地区", "list-country"),
				),
			)
			text = strings.Join(post.Results[0:34], " ")
		case "list-country":
			board = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("国内各省市", "list-province"),
				),
			)
			text = strings.Join(post.Results[35:], " ")
		}
		editedMsg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:      query.Message.Chat.ID,
				MessageID:   query.Message.MessageID,
				ReplyMarkup: &board,
			},
			Text:      text,
		}
		channel.MessageChannel <- editedMsg
	}
}

func GetData() (Jsons, error) {
	log := logrus.WithField("打开List文件","GetData")
	var post Jsons
	// 打开json文件
	fh, err := os.Open("list.txt")
	if err != nil {
		log.Errorln(err)
		return post, err
	}
	defer func(fh *os.File) {
		err := fh.Close()
		if err != nil {
			log.Errorln(err)
		}
		log.Println("关闭文件")
	}(fh)
	// 读取json文件，保存到jsonData中
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Errorln(err)
		return post, err
	}

	// 解析json数据到post中
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		log.Errorln(err)
		return post, err
	}
	return post, nil
}

func IsContain(data string) bool {
	var post, _ = GetData()
	for _, v := range post.Results {
		if v == data {
			return true
		}
	}
	return false
}
