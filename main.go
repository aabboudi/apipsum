package main

import (
	"apipsum/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.SetupIpsum(app)

	app.Listen(":3000")
}
