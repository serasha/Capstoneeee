package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
)

func SuperAdminRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/superadmin", handlers.GetAllSuperAdmin(db))
	app.Post("/api/superadmin", handlers.CreateSuperAdmin(db))
	app.Put("/api/superadmin/:id", handlers.UpdateSuperAdmin(db))
	app.Delete("/api/superadmin/:id", handlers.DeleteSuperAdmin(db))
}
