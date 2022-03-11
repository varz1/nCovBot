package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gofiber/fiber/v2"
	"github.com/varz1/nCovBot/channel"
	"github.com/varz1/nCovBot/variables"
	"log"
)

func WebHookHandler(c *fiber.Ctx) error {
	u := new(tgbotapi.Update)
	err := c.BodyParser(&u)
	if err != nil {
		log.Println("req解析失败")
		return err
	}
	channel.UpdateChannel <- *u
	return nil
}

func NotFoundHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendString("走错路啦")
}

func HiHandler(c *fiber.Ctx) error {
	return c.SendString("hi")
}

func BlogHandler(c *fiber.Ctx) error {
	return c.Redirect(variables.Blog, fiber.StatusOK)
}
