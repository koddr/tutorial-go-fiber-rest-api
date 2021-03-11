package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koddr/tutorial-go-rest-api-fiber/app/controllers"
	"github.com/koddr/tutorial-go-rest-api-fiber/pkg/middleware"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1", middleware.JWTProtected())

	// Routes for POST method:
	route.Post("/book", controllers.CreateBook) // create a new book

	// Routes for PATCH method:
	route.Patch("/book", controllers.UpdateBook) // update one book by ID

	// Routes for DELETE method:
	route.Delete("/book", controllers.DeleteBook) // delete one book by ID
}
