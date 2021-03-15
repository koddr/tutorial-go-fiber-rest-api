package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/configs"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/routes"
	"github.com/stretchr/testify/assert"
)

func TestPublicRoutes(t *testing.T) {
	// Define a structure for specifying input and output
	// data of a single test case. This structure is then used
	// to create a so called test map, which contains all test
	// cases, that should be run for testing this function
	tests := []struct {
		description   string
		route         string // input route
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "get all books",
			route:         "/api/v1/books",
			expectedError: false,
			expectedCode:  200,
		},
		{
			description:   "get book by ID",
			route:         "/api/v1/book/1",
			expectedError: false,
			expectedCode:  200,
		},
	}

	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Routes.
	routes.PublicRoutes(app) // Register a public routes for app.

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app.
		// The -1 disables request latency.
		resp, err := app.Test(req, -1)

		// Verify that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses, the next
		// test case needs to be processed
		if test.expectedError {
			continue
		}

		// Verify if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
