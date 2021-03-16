package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/utils"
)

// JWTRoutes func for describe group of public routes.
func JWTRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Post("/auth", func(c *fiber.Ctx) error {
		token, err := utils.GenerateNewJWTAccessToken(uuid.NewString(), []string{})
		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"error": false,
			"msg":   token,
		})
	})
}
