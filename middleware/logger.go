package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupLogger() *os.File {
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	log.SetOutput(logFile)
	return logFile
}

func Logger(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	log.Printf("%s %s - %d %s",
		c.Method(),
		c.OriginalURL(),
		c.Response().StatusCode(),
		time.Since(start))

	return err
}
