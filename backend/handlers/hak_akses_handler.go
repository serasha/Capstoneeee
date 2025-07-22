package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/models"
)

// GetAllHakAkses mengambil semua data hak akses (gabungan Admin dan User)
func GetAllHakAkses(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var admins []models.Admin
		var users []models.User
		
		// Ambil semua admin
		if err := db.Find(&admins).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data admin",
			})
		}
		
		// Ambil semua user
		if err := db.Find(&users).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengambil data user",
			})
		}
		
		// Gabungkan data untuk response
		var hakAksesList []fiber.Map
		
		// Tambahkan admin ke list
		for _, admin := range admins {
			hakAksesList = append(hakAksesList, fiber.Map{
				"id":           admin.ID,
				"id_user":      "ADM" + strconv.Itoa(int(admin.ID)),
				"nama_user":    admin.NamaAdmin,
				"username":     admin.Username,
				"email":        admin.EmailAdmin,
				"level_akses":  "Admin",
				"status_aktif": admin.StatusAktif,
				"created_at":   admin.CreatedAt,
				"type":         "admin",
			})
		}
		
		// Tambahkan user ke list
		for _, user := range users {
			hakAksesList = append(hakAksesList, fiber.Map{
				"id":           user.ID,
				"id_user":      "USR" + strconv.Itoa(int(user.ID)),
				"nama_user":    user.Username,
				"username":     user.Username,
				"email":        "-",
				"level_akses":  user.Role,
				"status_aktif": true, // User selalu aktif
				"created_at":   time.Now(), // User model tidak punya CreatedAt
				"type":         "user",
			})
		}
		
		return c.JSON(hakAksesList)
	}
}

// CreateHakAkses membuat user/admin baru
func CreateHakAkses(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var input struct {
			NamaUser    string `json:"nama_user"`
			Username    string `json:"username"`
			Email       string `json:"email"`
			Password    string `json:"password"`
			LevelAkses  string `json:"level_akses"`
			StatusAktif bool   `json:"status_aktif"`
		}
		
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}
		
		// Validasi input
		if input.NamaUser == "" || input.Username == "" || input.Password == "" || input.LevelAkses == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Nama user, username, password, dan level akses wajib diisi",
			})
		}
		
		// Cek apakah username sudah ada
		var existingAdmin models.Admin
		var existingUser models.User
		
		if db.Where("username = ?", input.Username).First(&existingAdmin).Error == nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Username sudah digunakan oleh admin",
			})
		}
		
		if db.Where("username = ?", input.Username).First(&existingUser).Error == nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Username sudah digunakan oleh user",
			})
		}
		
		// Buat berdasarkan level akses
		if input.LevelAkses == "admin" || input.LevelAkses == "Admin" {
			admin := models.Admin{
				Username:      input.Username,
				Password:      input.Password, // Dalam production, hash password
				NamaAdmin:     input.NamaUser,
				EmailAdmin:    input.Email,
				StatusAktif:   input.StatusAktif,
				Role:          "admin",
				TanggalDibuat: time.Now(),
			}
			
			if err := db.Create(&admin).Error; err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Gagal membuat admin",
				})
			}
			
			return c.Status(201).JSON(fiber.Map{
				"message": "Admin berhasil dibuat",
				"data":    admin,
			})
		} else {
			user := models.User{
				Username: input.Username,
				Password: input.Password, // Dalam production, hash password
				Role:     input.LevelAkses,
			}
			
			if err := db.Create(&user).Error; err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Gagal membuat user",
				})
			}
			
			return c.Status(201).JSON(fiber.Map{
				"message": "User berhasil dibuat",
				"data":    user,
			})
		}
	}
}

// UpdateHakAkses memperbarui data hak akses
func UpdateHakAkses(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		userType := c.Query("type") // "admin" atau "user"
		
		var input struct {
			NamaUser    string `json:"nama_user"`
			Username    string `json:"username"`
			Email       string `json:"email"`
			LevelAkses  string `json:"level_akses"`
			StatusAktif bool   `json:"status_aktif"`
		}
		
		if err := c.BodyParser(&input); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Permintaan tidak valid",
			})
		}
		
		if userType == "admin" {
			var admin models.Admin
			if err := db.First(&admin, id).Error; err != nil {
				return c.Status(404).JSON(fiber.Map{
					"error": "Admin tidak ditemukan",
				})
			}
			
			admin.NamaAdmin = input.NamaUser
			admin.Username = input.Username
			admin.EmailAdmin = input.Email
			admin.StatusAktif = input.StatusAktif
			
			if err := db.Save(&admin).Error; err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Gagal memperbarui admin",
				})
			}
			
			return c.JSON(fiber.Map{
				"message": "Admin berhasil diperbarui",
				"data":    admin,
			})
		} else {
			var user models.User
			if err := db.First(&user, id).Error; err != nil {
				return c.Status(404).JSON(fiber.Map{
					"error": "User tidak ditemukan",
				})
			}
			
			user.Username = input.Username
			user.Role = input.LevelAkses
			
			if err := db.Save(&user).Error; err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Gagal memperbarui user",
				})
			}
			
			return c.JSON(fiber.Map{
				"message": "User berhasil diperbarui",
				"data":    user,
			})
		}
	}
}

// DeleteHakAkses menghapus data hak akses
func DeleteHakAkses(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		userType := c.Query("type") // "admin" atau "user"
		
		if userType == "admin" {
			var admin models.Admin
			if err := db.First(&admin, id).Error; err != nil {
				return c.Status(404).JSON(fiber.Map{
					"error": "Admin tidak ditemukan",
				})
			}
			
			if err := db.Delete(&admin).Error; err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Gagal menghapus admin",
				})
			}
			
			return c.JSON(fiber.Map{
				"message": "Admin berhasil dihapus",
			})
		} else {
			var user models.User
			if err := db.First(&user, id).Error; err != nil {
				return c.Status(404).JSON(fiber.Map{
					"error": "User tidak ditemukan",
				})
			}
			
			if err := db.Delete(&user).Error; err != nil {
				return c.Status(500).JSON(fiber.Map{
					"error": "Gagal menghapus user",
				})
			}
			
			return c.JSON(fiber.Map{
				"message": "User berhasil dihapus",
			})
		}
	}
}

// GetHakAksesByID mengambil data hak akses berdasarkan ID
func GetHakAksesByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		userType := c.Query("type")
		
		if userType == "admin" {
			var admin models.Admin
			if err := db.First(&admin, id).Error; err != nil {
				return c.Status(404).JSON(fiber.Map{
					"error": "Admin tidak ditemukan",
				})
			}
			
			return c.JSON(fiber.Map{
				"id":           admin.ID,
				"id_user":      "ADM" + strconv.Itoa(int(admin.ID)),
				"nama_user":    admin.NamaAdmin,
				"username":     admin.Username,
				"email":        admin.EmailAdmin,
				"level_akses":  "Admin",
				"status_aktif": admin.StatusAktif,
				"created_at":   admin.CreatedAt,
				"type":         "admin",
			})
		} else {
			var user models.User
			if err := db.First(&user, id).Error; err != nil {
				return c.Status(404).JSON(fiber.Map{
					"error": "User tidak ditemukan",
				})
			}
			
			return c.JSON(fiber.Map{
				"id":           user.ID,
				"id_user":      "USR" + strconv.Itoa(int(user.ID)),
				"nama_user":    user.Username,
				"username":     user.Username,
				"email":        "-",
				"level_akses":  user.Role,
				"status_aktif": true,
				"created_at":   time.Now(),
				"type":         "user",
			})
		}
	}
}

// ToggleStatusAktif mengubah status aktif admin
func ToggleStatusAktif(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		userType := c.Query("type")
		
		if userType != "admin" {
			return c.Status(400).JSON(fiber.Map{
				"error": "Hanya admin yang memiliki status aktif",
			})
		}
		
		var admin models.Admin
		if err := db.First(&admin, id).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{
				"error": "Admin tidak ditemukan",
			})
		}
		
		admin.StatusAktif = !admin.StatusAktif
		
		if err := db.Save(&admin).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": "Gagal mengubah status",
			})
		}
		
		return c.JSON(fiber.Map{
			"message": "Status berhasil diubah",
			"status_aktif": admin.StatusAktif,
		})
	}
}