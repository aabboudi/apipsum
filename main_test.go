package main

import (
	"apipsum/controllers"
	"apipsum/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Init Fiber app for testing
func SetupTestApp() *fiber.App {
	app := fiber.New()
	routes.SetupRoutes(app)
	return app
}

// Test GenerateData() controller
func TestGenerateData(t *testing.T) {
	tests := []struct {
		name      string
		schema    map[string]interface{}
		expectErr bool
	}{
		{
			name: "Generate boolean data",
			schema: map[string]interface{}{
				"test_bool": map[string]interface{}{
					"type": "bool",
				},
			},
			expectErr: false,
		},
		{
			name: "Generate string data with max_length",
			schema: map[string]interface{}{
				"test_string": map[string]interface{}{
					"type":       "string",
					"max_length": 10,
				},
			},
			expectErr: false,
		},
		{
			name: "Generate email data",
			schema: map[string]interface{}{
				"test_email": map[string]interface{}{
					"type": "email",
				},
			},
			expectErr: false,
		},
		{
			name: "Invalid data type",
			schema: map[string]interface{}{
				"invalid_type": map[string]interface{}{
					"type": "unknown",
				},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := controllers.GenerateData(tt.schema)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, result)
			}
		})
	}
}

// Test availability of home page
func TestMainEndpoint(t *testing.T) {
	app := SetupTestApp()

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// Test availability of docs page
func TestDocsEndpoint(t *testing.T) {
	app := SetupTestApp()

	req := httptest.NewRequest("GET", "/docs/index.html", nil)
	resp, err := app.Test(req, -1)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
