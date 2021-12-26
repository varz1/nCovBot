package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/varz1/nCovBot/bot"
	"github.com/varz1/nCovBot/maker"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	app := fiber.New()
	app.Use(logger.New())
	go bot.Run()
	app.Post("/"+os.Getenv("TOKEN"), bot.WebHookHandler)
	go maker.List()
	go maker.Overall()
	go maker.Province()
	go maker.QueryProvince()
	go maker.News()
	go maker.RiskQuery()
	err2 := app.Listen(":" + port)
	if err2 != nil {
		log.Println(err2)
	}
}
