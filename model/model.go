package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Areas struct {
	Types string
	area []string
	Query tgbotapi.Message
}