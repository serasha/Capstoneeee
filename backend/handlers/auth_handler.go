// handlers/auth_handler.go
package handlers

import (
	"my-app/backend/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Format login tidak valid",
			})
		}

		var admin models.Admin
		// Hanya bisa login jika role admin
		if err := db.Where("username = ? AND password = ? AND role = ?", input.Username, input.Password, "admin").First(&admin).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Username atau password salah atau bukan akun admin",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Login admin berhasil",
			"data":    admin,
		})
	}
}

func LoginSuperAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Format login tidak valid",
			})
		}

		var superadmin models.SuperAdmin
		if err := db.Where("username = ? AND password = ?", input.Username, input.Password).First(&superadmin).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Username atau password salah",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Login superadmin berhasil",
			"data":    superadmin,
		})
	}
}

func RegisterAdmin(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Admin
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Input tidak valid",
			})
		}

		if input.Username == "" || input.Password == "" || input.EmailAdmin == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username, password, dan email wajib diisi",
			})
		}

		if err := db.Create(&input).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menyimpan admin",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Registrasi admin berhasil",
			"data":    input,
		})
	}
}

func LoginMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input LoginInput
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Format login tidak valid",
			})
		}

		var masyarakat models.Masyarakat
		if err := db.Where("email = ? AND password = ?", input.Username, input.Password).First(&masyarakat).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Email atau password salah",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Login masyarakat berhasil",
			"data":    masyarakat,
		})
	}
}

func RegisterMasyarakat(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input models.Masyarakat
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Input tidak valid",
			})
		}

		// Validasi minimal
		if input.NamaLengkap == "" || input.NIK == "" || input.Email == "" || input.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Nama lengkap, NIK, Email, dan Password wajib diisi",
			})
		}

		input.TglPendaftaran = time.Now()

		if err := db.Create(&input).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menyimpan data masyarakat",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Registrasi masyarakat berhasil",
			"data":    input,
		})
	}
}
