package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/handlers"
)

func KotaRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/kota", handlers.GetAllKota(db))
	app.Post("/api/kota", handlers.CreateKota(db))
	app.Put("/api/kota/:id", handlers.UpdateKota(db))
	app.Delete("/api/kota/:id", handlers.DeleteKota(db))
}