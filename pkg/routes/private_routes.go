package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/controllers"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for POST method:
	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook) // create a new book

	// Routes for PUT method:
	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook) // update one book by ID

	// Routes for DELETE method:
	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook) // delete one book by ID
}
