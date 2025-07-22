package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"my-app/backend/handlers"
)

func DokumenRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/api/dokumen/:pendaftaran_id", handlers.GetDokumenByPendaftaran(db))
	app.Post("/api/dokumen/upload", handlers.UploadDokumen(db))
	app.Put("/api/dokumen/:id/verifikasi", handlers.VerifikasiDokumen(db))
}