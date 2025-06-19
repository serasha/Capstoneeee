package routes

import (
	"my-app/backend/handlers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/api/login/admin", handlers.LoginAdmin(db))
	app.Post("/api/login/superadmin", handlers.LoginSuperAdmin(db))
	app.Post("/api/register/admin", handlers.RegisterAdmin(db))
}
