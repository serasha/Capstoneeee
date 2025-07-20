package handlers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"my-app/backend/models"
)

// RegisterRequest untuk menerima data dari frontend
type RegisterRequest struct {
	NamaLengkap string `json:"namaLengkap"`
	KataSandi   string `json:"kataSandi"`
}

// LoginRequest untuk menerima data login dari frontend
type LoginRequest struct {
	NamaLengkap string `json:"namaLengkap"`
	KataSandi   string `json:"kataSandi"`
}

// RegisterUser menyimpan user baru ke database
func RegisterUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input RegisterRequest

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		// Validasi input
		if input.NamaLengkap == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Nama lengkap wajib diisi",
			})
		}

		if input.KataSandi == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Kata sandi wajib diisi",
			})
		}

		if len(input.KataSandi) < 6 {
			return c.Status(400).JSON(fiber.Map{
				"error": "Kata sandi minimal 6 karakter",
			})
		}

		// Cek apakah username sudah digunakan
		var existing models.User
		if err := db.Where("username = ?", input.NamaLengkap).First(&existing).Error; err == nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Username sudah terdaftar",
			})
		}

		// Enkripsi password
		hashed, err := bcrypt.GenerateFromPassword([]byte(input.KataSandi), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengenkripsi password",
			})
		}

		// Buat user baru
		user := models.User{
			Username: input.NamaLengkap,
			Password: string(hashed),
		}

		if err := db.Create(&user).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal menyimpan user",
			})
		}

		// Response tanpa password
		return c.Status(201).JSON(fiber.Map{
			"message": "Registrasi berhasil",
			"user": fiber.Map{
				"id":       user.ID,
				"username": user.Username,
			},
		})
	}
}

// LoginUser memverifikasi kredensial dan mengizinkan login
func LoginUser(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input LoginRequest
		var user models.User

		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		// Validasi input
		if input.NamaLengkap == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Nama lengkap wajib diisi",
			})
		}

		if input.KataSandi == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Kata sandi wajib diisi",
			})
		}

		// Cari user berdasarkan username
		if err := db.Where("username = ?", input.NamaLengkap).First(&user).Error; err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Username tidak ditemukan",
			})
		}

		// Bandingkan password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.KataSandi)); err != nil {
			return c.Status(401).JSON(fiber.Map{
				"error": "Password salah",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Login berhasil",
			"user": fiber.Map{
				"id":       user.ID,
				"username": user.Username,
			},
		})
	}
}
