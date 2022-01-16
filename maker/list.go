package maker

import (
	"encoding/json"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/varz1/nCovBot/channel"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var listJson = Jsons{}

func init() {
	log1 := logrus.WithField("打开List文件", "init listJson")
	//var post Jsons
	p,_ := os.Getwd()
	jsonData, err := ioutil.ReadFile(p+"/list.json")
	if err != nil {
		log1.Errorln(err)
	}
	// 解析json数据到post中
	err = json.Unmarshal(jsonData, &listJson)
	if err != nil {
		log1.Errorln(err)
	}
}

type Jsons struct {
	Results []string `json:"results"`
}

func List() {
	//post, err := GetData()
	//if err != nil {
	//	log1.Errorln("打开文件错误", err)
	//	return
	//}
	text := ""
	var row []tgbotapi.InlineKeyboardButton
	var board = tgbotapi.NewInlineKeyboardMarkup()
	for query := range channel.ListQueryChannel {
		split := strings.Split(query.Data, "-")
		switch split[1] {
		case "province":
			row = append(row[0:0], tgbotapi.NewInlineKeyboardButtonData("各国家地区", "list-country-1"))
			text = strings.Join(listJson.Results[0:34], " ")
			board = tgbotapi.NewInlineKeyboardMarkup(row)
		case "country":
			i, _ := strconv.Atoi(split[2])
			text = strings.Join(listJson.Results[35+(i-1)*50:35+i*50], " ")
			board = GetPage(i)
		}
		editedMsg := tgbotapi.EditMessageTextConfig{
			BaseEdit: tgbotapi.BaseEdit{
				ChatID:      query.Message.Chat.ID,
				MessageID:   query.Message.MessageID,
				ReplyMarkup: &board,
			},
			Text: text,
		}
		channel.MessageChannel <- editedMsg
	}
}

// GetPage 分页
func GetPage(currentPage int) (markup tgbotapi.InlineKeyboardMarkup) {
	var row []tgbotapi.InlineKeyboardButton
	pageUp := ""
	pageDown := ""
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("国内各省市", "list-province"))
	if currentPage > 1 {
		pageDown = fmt.Sprintf("list-country-%d", currentPage-1)
		row = append(row, tgbotapi.NewInlineKeyboardButtonData("👈上一页", pageDown))
	}
	if currentPage < 5 {
		pageUp = fmt.Sprintf("list-country-%d", currentPage+1)
		row = append(row, tgbotapi.NewInlineKeyboardButtonData("下一页👉", pageUp))
	}
	markup = tgbotapi.NewInlineKeyboardMarkup(row)
	return
}

//func GetData() (Jsons, error) {
//	log1 := logrus.WithField("打开List文件", "GetData")
//	var post Jsons
//	// 打开json文件
//	fh, err := os.Open("list.json")
//	if err != nil {
//		log1.Errorln(err)
//		return post, err
//	}
//	defer func(fh *os.File) {
//		err := fh.Close()
//		if err != nil {
//			log1.Errorln(err)
//		}
//		log.Println("关闭文件")
//	}(fh)
//	// 读取json文件，保存到jsonData中
//	jsonData, err := ioutil.ReadAll(fh)
//	if err != nil {
//		log1.Errorln(err)
//		return post, err
//	}
//	// 解析json数据到post中
//	err = json.Unmarshal(jsonData, &post)
//	if err != nil {
//		log1.Errorln(err)
//		return post, err
//	}
//	return post, nil
//}

func IsContain(data string) bool {
	//var post, _ = GetData()
	for _, v := range listJson.Results {
		if v == data {
			return true
		}
	}
	return false
}
