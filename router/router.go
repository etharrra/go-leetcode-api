package router

import (
	"fmt"

	"github.com/etharrra/go-leetcode-api/handler"
	logger "github.com/etharrra/go-leetcode-api/loggger"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		err := logger.WriteLog("Logged")
		if err != nil {
			fmt.Println("Error writing log:", err)
		} else {
			fmt.Println("Log written successfully!")
		}
		return c.SendString("pong")
	})

	app.Get("/:username", handler.FetchUserProfile)
}
