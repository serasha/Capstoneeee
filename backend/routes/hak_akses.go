package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
)

func HakAksesRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/hak-akses", handlers.GetAllHakAkses(db))
	app.Get("/api/hak-akses/:id", handlers.GetHakAksesByID(db))
	app.Post("/api/hak-akses", handlers.CreateHakAkses(db))
	app.Put("/api/hak-akses/:id", handlers.UpdateHakAkses(db))
	app.Delete("/api/hak-akses/:id", handlers.DeleteHakAkses(db))
	app.Patch("/api/hak-akses/:id/toggle-status", handlers.ToggleStatusAktif(db))
}