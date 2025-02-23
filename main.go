package main

import (
	"github.com/etharrra/go-leetcode-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Listen(":3000")
}
