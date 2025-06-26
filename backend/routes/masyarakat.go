package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
)

type Masyarakat struct {
	gorm.Model
	Nama     string `json:"nama"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

func MasyarakatRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/masyarakat", handlers.GetAllMasyarakat(db))
	app.Post("/api/masyarakat", handlers.CreateMasyarakat(db))
	app.Put("/api/masyarakat/:id", handlers.UpdateMasyarakat(db))
	app.Delete("/api/masyarakat/:id", handlers.DeleteMasyarakat(db))
}
