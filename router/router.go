package router

import (
	"math/rand"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func generateData(schema map[string]string) map[string]interface{} {
	data := make(map[string]interface{})

	for key, valueType := range schema {
		switch valueType {
		case "string":
			data[key] = "random_string_" + strconv.Itoa(rand.Intn(1000))
		case "int":
			data[key] = rand.Intn(100)
		case "float":
			data[key] = rand.Float64() * 100
		case "bool":
			data[key] = rand.Intn(2) == 1
		default:
			data[key] = nil
		}
	}

	return data
}

// @Summary Generate JSON data
// @Description Generate JSON objects based on the schema provided in the request body
// @Tags Generate
// @Accept json
// @Produce json
// @Param count header int true "Number of objects to generate"
// @Param schema body map[string]string true "Schema of the JSON object"
// @Success 200 {array} map[string]interface{}
// @Failure 400 {string} string "Invalid request"
// @Router /api/generate [post]
func SetupIpsum(app *fiber.App) {
	app.Post("/api/generate", func(c *fiber.Ctx) error {
		countHeader := c.Get("count", "1")
		count, err := strconv.Atoi(countHeader)
		if err != nil || count <= 0 {
			return c.Status(400).SendString("Invalid count")
		}

		schema := make(map[string]string)
		if err := c.BodyParser(&schema); err != nil {
			return c.Status(400).SendString("Invalid JSON schema")
		}

		var results []map[string]interface{}
		for i := 0; i < count; i++ {
			results = append(results, generateData(schema))
		}

		return c.JSON(results)
	})
}

// @Summary Test API endpoint
// @Description Respond with status 200 if a GET request is sent to this endpoint. Used to verify the availability and responsiveness of the /api/generate endpoint.
// @Tags Generate
// @Produce json
// @Success 200 {object} map[string]interface{} "API is working"
// @Failure 400 {string} string "Invalid request"
// @Router /api/generate [get]
func TestIpsum(app *fiber.App) {
	app.Get("/api/generate", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":   200,
			"response": "Yes. The endpoint is working.",
		})
	})
}
