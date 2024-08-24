package routes

import (
	"apipsum/controllers"
	"strconv"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SetupRoutes(app *fiber.App) {
	testGenerateRoute(app)
	generateRoute(app)
	swaggerRoute(app)
	homePage(app)
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
func generateRoute(app *fiber.App) {
	app.Post("/api/generate", func(c *fiber.Ctx) error {
		countHeader := c.Get("count", "1")
		count, err := strconv.Atoi(countHeader)
		if err != nil || count <= 0 {
			return c.Status(400).SendString("Invalid count")
		}

		schema := make(map[string]interface{})
		if err := c.BodyParser(&schema); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"status": 400,
				"error":  "Invalid JSON schema",
			})
		}

		var results []map[string]interface{}
		for i := 0; i < count; i++ {
			data, err := controllers.GenerateData(schema)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"status": 400,
					"error":  err.Error(),
				})
			}
			results = append(results, data)
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
func testGenerateRoute(app *fiber.App) {
	app.Get("/api/generate", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":   200,
			"response": "Yes. The endpoint is working.",
		})
	})
}

func swaggerRoute(app *fiber.App) {
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("/docs/index.html")
	})
	app.Get("/docs/*", fiberSwagger.WrapHandler)
}

func homePage(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("views/index.html")
	})
}
