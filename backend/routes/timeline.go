package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/handlers"
)

func TimelineRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/timeline/:pendaftaran_id", handlers.GetTimelineByPendaftaran(db))
	app.Put("/api/timeline/:id", handlers.UpdateTimeline(db))
}