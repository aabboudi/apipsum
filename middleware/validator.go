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
			"response": "Invalid count. Please limit the count in your request to under 1000",
		})
	}

	schema := make(map[string]interface{})
	if err := c.BodyParser(&schema); err != nil {
		const status = fiber.StatusBadRequest
		return c.Status(status).JSON(fiber.Map{
			"status": status,
			"error":  "Invalid JSON schema",
		})
	}

	estimatedSize := estimateResponseSize(schema) * count
	const maxResponseSize = 1048576 // 1MB
	if estimatedSize > maxResponseSize {
		const status = fiber.StatusRequestEntityTooLarge
		return c.Status(status).JSON(fiber.Map{
			"status":         status,
			"response":       "Estimated response too large",
			"estimated_size": "Estimated a response of " + strconv.Itoa(estimatedSize) + " bytes but the limit is " + strconv.Itoa(maxResponseSize),
		})
	}

	return c.Next()
}

func estimateResponseSize(schema map[string]interface{}) int {
	objectSize := 0

	for _, fieldInterface := range schema {
		field := fieldInterface.(map[string]interface{})
		fieldType := field["type"].(string)

		switch fieldType {
		case "bool":
			objectSize++
		case "string", "email", "url", "slug":
			maxLength, exists := field["max_length"].(int)
			if !exists {
				maxLength = 100
			}
			objectSize += maxLength
		case "int":
			objectSize += 4
		case "float":
			objectSize += 8
		case "datetime":
			objectSize += 7
		}
	}

	return objectSize
}
