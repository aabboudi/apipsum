package main

import (
	_ "apipsum/docs"
	"apipsum/middleware"
	"apipsum/routes"

	"github.com/gofiber/fiber/v2"
)

// @title           APIpsum
// @version         0.1.0
// @description     This is a sample server.
// @host            localhost:3000
// @BasePath        /
func main() {
	logFile := middleware.SetupLogger()
	defer logFile.Close()

	app := fiber.New()
	app.Static("/", "./static")
	app.Use(middleware.Logger)
	routes.SetupRoutes(app)

	app.Listen(":3000")
}
