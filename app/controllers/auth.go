package controllers

import (
	"go-fiber/databases/conn"
	"go-fiber/databases/migrations"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error)  {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func Login(ctx *fiber.Ctx) error {
	var request migrations.User

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Bad Request",
		})
	}

	// Find the user by email or username
	var user migrations.User
	result := conn.DB.Where("email = ? OR username = ?", request.Email, request.Username).First(&user)
	if result.Error != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Compare the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Passwords match, user is authenticated
	return ctx.JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
	})
}


func Register(ctx *fiber.Ctx) error {
    var user migrations.User

    if err := ctx.BodyParser(&user); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Bad Request",
        })
    }

    hashedPassword, err := Hash(user.Password)

    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Internal server error",
        })
    }

	user.Password = hashedPassword

    // Create the user in the database using RegisterRequest fields
    conn.DB.Create(&user)

    // Modify the response to include more information about the registered user
    return ctx.JSON(fiber.Map{
        "message": "Registration successful",
    })
}
