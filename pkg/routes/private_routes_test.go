package routes

import (
	"io"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPrivateRoutes(t *testing.T) {
	// Load .env.test file from the root folder.
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}

	// Define a structure for specifying input and output data of a single test case.
	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		tokenString   string // input token
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "delete book without JWT and body",
			route:         "/api/v1/book",
			method:        "DELETE",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "delete book without right credentials",
			route:         "/api/v1/book",
			method:        "DELETE",
			tokenString:   "Bearer " + os.Getenv("FAKE_NO_ACCESS"),
			body:          strings.NewReader(`{"id": "808b1530-89ec-4f88-a7e7-139501d129c2"}`),
			expectedError: false,
			expectedCode:  403,
		},
		{
			description:   "delete book with credentials",
			route:         "/api/v1/book",
			method:        "DELETE",
			tokenString:   "Bearer " + os.Getenv("FAKE_ONLY_DELETE_ACCESS"),
			body:          strings.NewReader(`{"id": "808b1530-89ec-4f88-a7e7-139501d129c2"}`),
			expectedError: false,
			expectedCode:  404,
		},
	}

	// Define a new Fiber app with config.
	app := fiber.New()

	// Define routes.
	PrivateRoutes(app)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case.
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Authorization", test.tokenString)
		req.Header.Set("Content-Type", "application/json")

		// Perform the request plain with the app.
		resp, err := app.Test(req, -1) // the -1 disables request latency

		// Verify, that no error occurred, that is not expected
		assert.Equalf(t, test.expectedError, err != nil, test.description)

		// As expected errors lead to broken responses,
		// the next test case needs to be processed.
		if test.expectedError {
			continue
		}

		// Verify, if the status code is as expected.
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
