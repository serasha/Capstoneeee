package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
)

func AdminRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/admin", handlers.GetAllAdmin(db))
	app.Post("/api/admin", handlers.CreateAdmin(db))
	app.Put("/api/admin/:id", handlers.UpdateAdmin(db))
	app.Delete("/api/admin/:id", handlers.DeleteAdmin(db))
	app.Get("/api/admin/dashboard-stat", handlers.StatistikDashboardAdmin(db))
}