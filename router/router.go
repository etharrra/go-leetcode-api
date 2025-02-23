package router

import (
	"github.com/etharrra/go-leetcode-api/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/:username", handler.FetchUserProfile)
}
