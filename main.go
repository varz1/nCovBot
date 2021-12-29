package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/varz1/nCovBot/bot"
	"github.com/varz1/nCovBot/maker"
	"os"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Static("/", "./public")
	//bot.SetUpRouter(app)
	go bot.Run()
	go maker.List()
	go maker.Overall()
	go maker.Province()
	go maker.QueryProvince()
	go maker.News()
	go maker.RiskQuery()
	app.Post("/"+os.Getenv("TOKEN"), bot.WebHookHandler)
	app.Use(bot.NotFoundHandler)
	app.Listen(":" + os.Getenv("PORT"))
}
