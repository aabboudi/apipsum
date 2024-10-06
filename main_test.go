package main

import (
	"apipsum/controllers"
	"apipsum/routes"
	"bytes"
	"encoding/json"
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

// Test the /api/generate POST endpoint
func TestGenerateDataEndpoint(t *testing.T) {
	app := SetupTestApp()

	sampleSchema := map[string]interface{}{
		"name": map[string]interface{}{
			"type":       "string",
			"max_length": 20,
		},
		"age": map[string]interface{}{
			"type": "int",
			"min":  1,
			"max":  100,
		},
		"email": map[string]interface{}{
			"type": "email",
		},
	}

	body, _ := json.Marshal(sampleSchema)
	req := httptest.NewRequest("POST", "/api/generate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("count", "3")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var results []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&results)

	// Check if response count is 3
	assert.Len(t, results, 3)

	for _, result := range results {
		assert.NotEmpty(t, result["name"])
		assert.NotEmpty(t, result["age"])
		assert.NotEmpty(t, result["email"])
	}
}

// Test the /api/generate POST endpoint with invalid schema
func TestInvalidGenerateDataEndpoint(t *testing.T) {
	app := SetupTestApp()

	invalidSchema := map[string]interface{}{
		"invalid_type": map[string]interface{}{
			"type": "unknown",
		},
	}

	body, _ := json.Marshal(invalidSchema)
	req := httptest.NewRequest("POST", "/api/generate", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("count", "1")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var errorResponse map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&errorResponse)

	assert.Equal(t, float64(400), errorResponse["status"])
	assert.Contains(t, errorResponse["error"], "invalid data type")
}
