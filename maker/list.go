package maker

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/varz1/nCovBot/channel"
	"io/ioutil"
	"os"
	"strings"
)

type res struct {
	Results []string `json:"results"`
}

var board = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("国内各省市", "province"),
	),
)
var board1 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("国外各国家地区", "country"),
	),
)

func List() {
	// 打开json文件
	fh, err := os.Open("list.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(fh *os.File) {
		err := fh.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(fh)
	// 读取json文件，保存到jsonData中
	jsonData, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println(err)
		return
	}

	var post res
	// 解析json数据到post中
	err = json.Unmarshal(jsonData, &post)
	if err != nil {
		fmt.Println(err)
		return
	}
	for area := range channel.ListChannel {
		var c tgbotapi.Chattable
		text := "加载文件错误"
		switch area.Types {
		case "province":
			text = strings.Join(post.Results[0:34], " ")
			msg := tgbotapi.NewEditMessageText(area.Query.Chat.ID, area.Query.MessageID, text)
			msg.ReplyMarkup = &board1
			c = msg
		case "country":
			text = strings.Join(post.Results[35:], " ")
			msg := tgbotapi.NewEditMessageText(area.Query.Chat.ID, area.Query.MessageID, text)
			msg.ReplyMarkup = &board
			c = msg
		}
		channel.MessageChannel <- c
	}
}
