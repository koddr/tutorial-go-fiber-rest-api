package controllers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/koddr/tutorial-go-fiber-rest-api/app/models"
	"github.com/koddr/tutorial-go-fiber-rest-api/pkg/utils"
	"github.com/koddr/tutorial-go-fiber-rest-api/platform/database"
)

// UserLogin method auth user and return JWT (Access & Refresh).
// @Description Auth user and return JWT.
// @Summary auth user and return JWT
// @Tags Public
// @Accept json
// @Produce json
// @Param email body string true "User Email"
// @Param password body string true "User Password"
// @Success 200 {object} models.User
// @Router /api/v1/user/login [post]
func UserLogin(c *fiber.Ctx) error {
	// Create a new user auth struct.
	auth := &models.Auth{}

	// Checking received data from JSON body.
	if err := c.BodyParser(auth); err != nil {
		// Return, if JSON data is not correct.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get user by email.
	foundedUser, err := db.GetUserByEmail(auth.Email)
	if err != nil {
		// Return, if user not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with the given login is not found",
		})
	}

	// Generate JWT Access token.
	accessToken, err := utils.GenerateNewJWTAccessToken(
		foundedUser.ID.String(),
		[]string{"book:create", "book:update", "book:delete"},
	)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Generate JWT Refresh token.
	refreshToken, err := utils.GenerateNewJWTRefreshToken()
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"tokens": fiber.Map{
			"access":  accessToken,
			"refresh": refreshToken,
		},
	})
}

// CreateUser func for create a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags Public
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Success 201 {object} models.User
// @Router /api/v1/user/register [post]
func CreateUser(c *fiber.Ctx) error {
	// Create a new user auth struct.
	auth := &models.Auth{}

	// Checking received data from JSON body.
	if err := c.BodyParser(auth); err != nil {
		// Return, if JSON data is not correct.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a User model.
	validate := validator.New()

	// Validate auth fields.
	if err := validate.Struct(auth); err != nil {
		// Return, if some fields are not valid.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   utils.ValidatorErrors(err),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new user struct.
	user := &models.User{}

	// Set initialized default data for user:
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.Email = auth.Email
	user.PasswordHash = utils.GeneratePassword(auth.Password)
	user.UserStatus = 1 // 0 == draft, 1 == active

	// Create a new user with validated data.
	if err := db.CreateUser(user); err != nil {
		// Return status 500 and create user process error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Clear password hash field for JSON view.
	user.PasswordHash = ""

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}
