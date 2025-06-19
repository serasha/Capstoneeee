package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
)

func DikelolaRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/dikelola", handlers.GetAllDikelola(db))
	app.Post("/api/dikelola", handlers.CreateDikelola(db))
	app.Put("/api/dikelola/:id", handlers.UpdateDikelola(db))
	app.Delete("/api/dikelola/:id", handlers.DeleteDikelola(db))
}