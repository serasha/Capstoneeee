package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	MasyarakatRoutes(app, db)
	PendaftaranRoutes(app, db)
	AdminRoutes(app, db)
	DikelolaRoutes(app, db)
	MengelolaRoutes(app, db)
	SuperAdminRoutes(app, db)
	AuthRoutes(app, db)
}
