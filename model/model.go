package model

import "github.com/go-telegram-bot-api/telegram-bot-api"

type Areas struct {
	Types string
	area []string
	AreaMessage tgbotapi.Message
}