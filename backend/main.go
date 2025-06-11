package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my-app/backend/models" // pastikan path ini sesuai module kamu
)

type App struct {
	DB *gorm.DB
}

func main() {
	// Koneksi ke PostgreSQL via GORM
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=latihan-capstone sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal membuka koneksi ke database:", err)
	}

	// Migrasi tabel jika belum ada
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Gagal migrate tabel users:", err)
	}
	log.Println("Berhasil terhubung dan migrate database!")

	// Buat instance App
	appInstance := &App{DB: db}

	// Fiber setup
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Halo dari Fiber!"})
	})
	app.Get("/api/users", appInstance.getUsers)

	// Serve frontend
	app.Static("/", "../frontend/dist")

	// Start server
	if err := app.Listen(":8070"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}

func (a *App) getUsers(c *fiber.Ctx) error {
	users, err := models.GetAllUsers(a.DB)
	if err != nil {
		log.Println("DB Error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data dari database",
		})
	}
	return c.JSON(fiber.Map{"data": users})
}
