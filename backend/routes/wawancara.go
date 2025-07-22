package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/handlers"
)

func WawancaraRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/wawancara", handlers.GetAllWawancara(db))
	app.Get("/api/wawancara/:pendaftaran_id", handlers.GetWawancaraByPendaftaran(db))
	app.Post("/api/wawancara", handlers.CreateWawancara(db))
	app.Put("/api/wawancara/:id", handlers.UpdateHasilWawancara(db))
}