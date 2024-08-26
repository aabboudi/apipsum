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

func ResponseLimiter(c *fiber.Ctx) error {
	countHeader := c.Get("count", "1")
	count, err := strconv.Atoi(countHeader)
	if err != nil || count <= 0 {
		const status = fiber.StatusBadRequest
		return c.Status(status).JSON(fiber.Map{
			"status":   status,
			"response": "Invalid count",
		})
	} else if count > 1000 {
		const status = fiber.StatusRequestEntityTooLarge
		return c.Status(status).JSON(fiber.Map{
			"status":   status,
			"response": "Invalid count. Please limit your request to under 1000",
		})
	}

	schema := c.Body()
	estimatedSize := estimateResponseSize(schema) * count
	const maxResponseSize = 1048576 // 1MB
	if estimatedSize > maxResponseSize {
		const status = fiber.StatusRequestEntityTooLarge
		return c.Status(status).JSON(fiber.Map{
			"status":   status,
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
