package maker

import (
	"encoding/json"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Jsons struct {
	Results []string `json:"results"`
}

func List() {
	post, err := GetData()
	if err != nil {
		log.Println(err)
		return
	}
	var c tgbotapi.Chattable
	text := "请选择区域"
	for area := range channel.ListChannel {
		switch area.Types {
		case "menu":
			var menu = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("国内各省市", "province"),
				),
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("国内外各国家地区", "country"),
				),
			)
			msg := tgbotapi.NewMessage(area.AreaMessage.Chat.ID, text)
			msg.ReplyMarkup = menu
			c = msg
		case "province":
			var board1 = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("国内外各国家地区", "country"),
				),
			)
			text = strings.Join(post.Results[0:34], " ")
			msg := tgbotapi.NewEditMessageText(area.AreaMessage.Chat.ID, area.AreaMessage.MessageID, text)
			msg.ReplyMarkup = &board1
			c = msg
		case "country":
			var board = tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("国内各省市", "province"),
				),
			)
			text = strings.Join(post.Results[35:], " ")
			msg := tgbotapi.NewEditMessageText(area.AreaMessage.Chat.ID, area.AreaMessage.MessageID, text)
			msg.ReplyMarkup = &board
			c = msg
		}
		channel.MessageChannel <- c
	}
}

func GetData() (Jsons, error) {
	var post Jsons
	// 打开json文件
	fh, err := os.Open("list.txt")
	if err != nil {
		log.Println(err)
		return post, err
	}
	defer func(fh *os.File) {
		err := fh.Close()
		if err != nil {
			log.Println(err)
		}
	}(fh)
	// 读取json文件，保存到jsonData中
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Println(err)
		return post, err
	}

	// 解析json数据到post中
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		log.Println(err)
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
