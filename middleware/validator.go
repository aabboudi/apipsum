package middleware

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RequestLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        5,
		Expiration: 30 * time.Second,
	})
}

const maxResponseSize = 1048576 // 1MB

func ResponseLimiter(c *fiber.Ctx) error {
	countHeader := c.Get("count", "1")
	count, err := strconv.Atoi(countHeader)
	if err != nil || count <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":   fiber.StatusBadRequest,
			"response": "Invalid count",
		})
	}

	schema := c.Body()
	estimatedSize := estimateResponseSize(schema) * count
	if estimatedSize > maxResponseSize {
		return c.Status(fiber.StatusRequestEntityTooLarge).JSON(fiber.Map{
			"status":   fiber.StatusRequestEntityTooLarge,
			"response": "Estimated response too large",
		})
	}

	return c.Next()
}

func estimateResponseSize(schema []byte) int {
	// Placeholder size estimation
	sizePerItem := len(schema)
	return sizePerItem
}
