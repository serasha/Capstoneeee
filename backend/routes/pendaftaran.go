package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
)

func PendaftaranRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/pendaftaran", handlers.GetAllPendaftaran(db))
	app.Post("/api/pendaftaran", handlers.CreatePendaftaran(db))
	app.Put("/api/pendaftaran/:id", handlers.UpdatePendaftaran(db))
	app.Delete("/api/pendaftaran/:id", handlers.DeletePendaftaran(db))
}
