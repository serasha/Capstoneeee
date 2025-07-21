package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"my-app/backend/handlers"
	"my-app/backend/middleware"
)

func PendaftaranRoutes(router fiber.Router, db *gorm.DB) {
	router.Get("/pendaftaran", handlers.GetAllPendaftaran(db))
	router.Get("/pendaftaran/user", middleware.AuthRequired, handlers.GetPendaftaranByUser(db))
	router.Get("/pendaftaran/list", handlers.GetPendaftaranList(db))
	router.Post("/pendaftaran", middleware.AuthRequired, handlers.CreatePendaftaran(db))
	router.Put("/pendaftaran/:id", handlers.UpdatePendaftaran(db))
	router.Patch("/pendaftaran/:id/verifikasi", handlers.VerifikasiPendaftaran(db))
	router.Delete("/pendaftaran/:id", handlers.DeletePendaftaran(db))
}
