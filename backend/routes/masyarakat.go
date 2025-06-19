package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
)

func MasyarakatRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/masyarakat", handlers.GetAllMasyarakat(db))
	app.Post("/api/masyarakat", handlers.CreateMasyarakat(db))
	app.Put("/api/masyarakat/:id", handlers.UpdateMasyarakat(db))
	app.Delete("/api/masyarakat/:id", handlers.DeleteMasyarakat(db))
}
