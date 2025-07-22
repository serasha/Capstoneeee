package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"my-app/backend/database"
	"my-app/backend/models"
	"my-app/backend/routes"
	"my-app/backend/utils"
)

var store = session.New()

type App struct {
	DB *gorm.DB
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// 🔌 Koneksi database
	database.ConnectDB()

	// 🔄 Drop table lama (opsional)
	database.DB.Migrator().DropTable(
		&models.Mengelola{},
		&models.Dikelola{},
		&models.Pendaftaran{},
		&models.Masyarakat{},
		&models.Admin{},
		&models.SuperAdmin{},
		&models.User{},
		&models.LogAktifitas{},
	)
	log.Println("Dropping old tables...")

	// ⚙️ Migrasi ulang tabel
	err := database.DB.AutoMigrate(
		&models.Admin{},
		&models.Masyarakat{},
		&models.Mendaftar{},
		&models.Pendaftaran{},
		&models.Mengelola{},
		&models.Dikelola{},
		&models.SuperAdmin{},
		&models.User{},
		&models.LogAktifitas{},
		&models.Kota{},
		&models.Dokumen{},
		&models.Timeline{},
		&models.Wawancara{},
	)
	if err != nil {
		log.Fatal("Migrasi gagal:", err)
	}
	log.Println("Migrasi berhasil!")

	// Jalankan seeder (bisa dikomentari jika tidak ingin seed setiap start)
	utils.SeedAll(database.DB)

	// 🚀 Setup Fiber
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// 🧪 Contoh API endpoint sederhana
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Halo dari Fiber!"})
	})

	// 🔀 Register route API (login & daftar)
	api := app.Group("/api") // grup semua route API
	routes.SetupLoginRoutes(api)
	routes.SetupDaftarRoutes(api)
	routes.SetupUserRoutes(api)
	routes.PendaftaranRoutes(api, database.DB)
	routes.LogAktifitasRoutes(api, database.DB)
	routes.AdminRoutes(app, database.DB)
	routes.KotaRoutes(app, database.DB)
	routes.TimelineRoutes(app, database.DB)
	routes.WawancaraRoutes(app, database.DB)
	routes.DokumenRoutes(app, database.DB)

	// 📦 Serve static files (Vue build result)
	app.Static("/", "../frontend/dist")

	// 🔁 Fallback ke index.html untuk semua non-API route (SPA mode Vue)
	app.Use(func(c *fiber.Ctx) error {
		// Skip jika request ke API
		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}
		return c.SendFile(filepath.Join("..", "frontend", "dist", "index.html"))
	})

	// 🔊 Jalankan server
	if err := app.Listen(":8070"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
