package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
)

func MengelolaRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/mengelola", handlers.GetAllMengelola(db))
	app.Post("/api/mengelola", handlers.CreateMengelola(db))
	app.Put("/api/mengelola/:id", handlers.UpdateMengelola(db))
	app.Delete("/api/mengelola/:id", handlers.DeleteMengelola(db))
}