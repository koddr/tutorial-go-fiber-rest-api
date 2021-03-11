package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koddr/tutorial-go-rest-api-fiber/app/controllers"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/book/:id", controllers.GetBook) // get one book by ID
	route.Get("/books", controllers.GetBooks)   // get list of all books
}
