package routes

import (
	"github.com/gofiber/fiber/v2"
	"my-app/backend/database"
	"my-app/backend/models"
)

func SetupLoginRoutes(route fiber.Router) {
	route.Post("/login", func(c *fiber.Ctx) error {
		type LoginInput struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var input LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Data tidak valid",
			})
		}

		var user models.Masyarakat
		if err := database.DB.Where("email = ? AND password = ?", input.Email, input.Password).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Email atau password salah",
			})
		}

		// (Optional) Tambahkan token di sini jika pakai JWT
		return c.JSON(fiber.Map{
			"message": "Login berhasil",
			"user":    user,
		})
	})
}
