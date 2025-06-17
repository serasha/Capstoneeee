package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"my-app/backend/models"
	"my-app/backend/handlers"
)

type App struct {
	DB *gorm.DB
}

func setupRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/masyarakat", handlers.GetAllMasyarakat(db))
	app.Post("/api/masyarakat", handlers.CreateMasyarakat(db))

	app.Get("/api/pendaftaran", handlers.GetAllPendaftaran(db))
	app.Post("/api/pendaftaran", handlers.CreatePendaftaran(db))

	app.Get("/api/admin", handlers.GetAllAdmin(db))
	app.Post("/api/admin", handlers.CreateAdmin(db))

	app.Get("/api/superadmin", handlers.GetAllSuperAdmin(db))
	app.Post("/api/superadmin", handlers.CreateSuperAdmin(db))
}

func main() {
	// Koneksi ke PostgreSQL via GORM
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=latihan-capstone sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal membuka koneksi ke database:", err)
	}

	// Migrasi semua tabel
	if err := db.AutoMigrate(
		&models.User{},
		&models.Masyarakat{},
		&models.Pendaftaran{},
		&models.Admin{},
		&models.SuperAdmin{},
	); err != nil {
		log.Fatal("Gagal migrate semua tabel:", err)
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

	setupRoutes(app, db)

	// Serve frontend
	app.Static("/", "../frontend/dist")

	// Start server
	if err := app.Listen(":8888"); err != nil {
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


