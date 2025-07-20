package routes

import (
	"my-app/backend/database"
	"my-app/backend/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"strconv"
)

var jwtSecret = []byte("supersecretkey")

func generateJWT(userID uint, username string, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"username": username,
		"role": role,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	if cookie == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}
	c.Locals("user_id", claims["user_id"])
	c.Locals("username", claims["username"])
	return c.Next()
}

func SetupUserRoutes(route fiber.Router) {
	userGroup := route.Group("/user")

	// Register endpoint yang sesuai dengan frontend
	userGroup.Post("/register", func(c *fiber.Ctx) error {
		type RegisterRequest struct {
			NamaLengkap string `json:"namaLengkap"`
			KataSandi   string `json:"kataSandi"`
			Role        string `json:"role"`
		}

		var req RegisterRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Data tidak valid",
			})
		}

		// Validasi input
		if req.NamaLengkap == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Nama lengkap wajib diisi",
			})
		}

		if req.KataSandi == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Kata sandi wajib diisi",
			})
		}

		if len(req.KataSandi) < 6 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Kata sandi minimal 6 karakter",
			})
		}

		// Cek apakah username sudah digunakan
		var existingUser models.User
		if err := database.DB.Where("username = ?", req.NamaLengkap).First(&existingUser).Error; err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username sudah terdaftar",
			})
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.KataSandi), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal mengenkripsi password",
			})
		}

		// Buat user baru
		role := req.Role
		if role != "admin" { role = "user" }
		user := models.User{
			Username: req.NamaLengkap,
			Password: string(hashedPassword),
			Role: role,
		}

		// Simpan ke database
		if err := database.DB.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal menyimpan user",
			})
		}

		// Buat data masyarakat otomatis dengan NIK dan Email unik
		nikAuto := "AUTO-" + strconv.Itoa(int(user.ID))
		emailAuto := user.Username + "+auto@yourapp.local"
		masyarakat := models.Masyarakat{
			NamaLengkap: req.NamaLengkap,
			NIK: nikAuto,
			Email: emailAuto,
			Password: string(hashedPassword),
			TglPendaftaran: time.Now(),
		}
		if err := database.DB.Create(&masyarakat).Error; err != nil {
			// Jika gagal, hapus user yang baru saja dibuat agar tidak orphan
			database.DB.Delete(&user)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Gagal membuat data masyarakat, pastikan data unik (NIK/Email) tidak bentrok",
			})
		}

		// Return success response
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Registrasi berhasil",
			"user": fiber.Map{
				"id":       user.ID,
				"username": user.Username,
				"role": user.Role,
			},
		})
	})

	// Login endpoint - JWT
	userGroup.Post("/login", func(c *fiber.Ctx) error {
		type LoginRequest struct {
			NamaLengkap string `json:"namaLengkap"`
			KataSandi   string `json:"kataSandi"`
		}

		var req LoginRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}

		// Validasi input
		if req.NamaLengkap == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Nama lengkap wajib diisi",
			})
		}

		if req.KataSandi == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Kata sandi wajib diisi",
			})
		}

		// Cari user di database
		var user models.User
		if err := database.DB.Where("username = ?", req.NamaLengkap).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Username tidak ditemukan",
			})
		}

		// Verifikasi password
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.KataSandi)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Password salah",
			})
		}

		// Generate JWT
		tokenString, err := generateJWT(user.ID, user.Username, user.Role)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal membuat token"})
		}
		// Set JWT di cookie
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    tokenString,
			Expires:  time.Now().Add(24 * time.Hour),
			HTTPOnly: true,
			Secure:   false, // set true jika pakai https
			SameSite: fiber.CookieSameSiteLaxMode,
		})

		return c.JSON(fiber.Map{
			"message": "Login berhasil",
			"user": fiber.Map{
				"id":       user.ID,
				"username": user.Username,
				"role": user.Role,
			},
		})
	})

	// Logout endpoint (hapus cookie JWT)
	userGroup.Post("/logout", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-1 * time.Hour),
			HTTPOnly: true,
			Secure:   false, // set true jika pakai https
			SameSite: fiber.CookieSameSiteLaxMode,
		})
		return c.JSON(fiber.Map{"message": "Logout berhasil"})
	})

	// Endpoint cek user login (me)
	userGroup.Get("/me", AuthRequired, func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		username := c.Locals("username")
		role := "user"
		cookie := c.Cookies("jwt")
		if cookie != "" {
			token, _ := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) { return jwtSecret, nil })
			if token != nil && token.Valid {
				claims, ok := token.Claims.(jwt.MapClaims)
				if ok && claims["role"] != nil {
					role = claims["role"].(string)
				}
			}
		}
		return c.JSON(fiber.Map{
			"id": userID,
			"username": username,
			"role": role,
		})
	})
}
