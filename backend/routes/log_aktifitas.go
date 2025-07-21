package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/handlers"
)

func LogAktifitasRoutes(router fiber.Router, db *gorm.DB) {
	router.Get("/log-aktivitas", handlers.GetLogAktifitasList(db))
} 