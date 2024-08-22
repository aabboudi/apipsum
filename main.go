package main

import (
	_ "apipsum/docs"
	"apipsum/router"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title           APIpsum
// @version         0.1.0
// @description     This is a sample server.
// @host            localhost:3000
// @BasePath        /
func main() {
	app := fiber.New()

	app.Static("/", "./static")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("views/index.html")
	})

	router.TestIpsum(app)
	router.SetupIpsum(app)

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/index.html")
	})

	app.Get("/docs/*", fiberSwagger.WrapHandler)

	app.Listen(":3000")
}
