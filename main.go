package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/configs"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/middleware"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/routes"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/utils"

	_ "github.com/joho/godotenv/autoload"                // load .env file automatically
	_ "github.com/koddr/tutorial-go-fiber-rest-api/docs" // load API Docs files (Swagger)
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
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
	utils.StartServer(app)
}
