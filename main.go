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
	bot.SetUpRouter(app)
	go bot.Run()
	go maker.List()
	go maker.Overall()
	go maker.WorldOverall()
	go maker.Trend()
	go maker.Province()
	go maker.QueryProvince()
	go maker.News()
	go maker.RiskQuery()
	app.Listen(":" + os.Getenv("PORT"))
}
