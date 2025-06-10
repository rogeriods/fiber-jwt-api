package handlers

import (
	"rogeriods/fiber-jwt-api/configs"
	"rogeriods/fiber-jwt-api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @route	POST /register
// @desc	Create user
// @access	Public
func Register(c *fiber.Ctx) error {
	var input models.User
	// Check if JSON raw could be parsed to User model
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	input.Password = string(hashedPassword)

	// Save user to DB
	if err := configs.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create user"})
	}

	return c.JSON(fiber.Map{"message": "User created"})
}

// @route	POST /login
// @desc	Login return JWT Token
// @access	Public
func Login(c *fiber.Ctx) error {
	var input models.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var user models.User
	if err := configs.DB.Where("username = ?", input.Username, input.Password).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"user_name": user.Username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	t, err := token.SignedString(configs.JWTSecret)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

// @route	GET /api/profile
// @desc	Get profile data
// @access	Private
func Profile(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	userName := c.Locals("userName")
	return c.JSON(fiber.Map{"message": "Welcome!", "user_id": userID, "user_name": userName})
}
