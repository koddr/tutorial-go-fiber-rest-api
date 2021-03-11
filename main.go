package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koddr/tutorial-go-rest-api-fiber/pkg/configs"
	"github.com/koddr/tutorial-go-rest-api-fiber/pkg/middleware"
	"github.com/koddr/tutorial-go-rest-api-fiber/pkg/routes"
	"github.com/koddr/tutorial-go-rest-api-fiber/pkg/utils"

	_ "github.com/joho/godotenv/autoload"                // load .env file automatically
	_ "github.com/koddr/tutorial-go-rest-api-fiber/docs" // load API Docs files (Swagger)
)

// @title Fiber Template API
// @version 1.0
// @description This is an auto-generated API Docs for Fiber Template.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)
}
