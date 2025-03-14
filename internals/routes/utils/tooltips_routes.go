package utils

import (
	"arabiya-syari-fiber/internals/controllers/utils"

	"arabiya-syari-fiber/internals/controllers/user" // Middleware Auth
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// 🔥 Setup Routes untuk Tooltips
func UtilsRoutes(app *fiber.App, db *gorm.DB) {
	// 🔒 Middleware Auth diaktifkan untuk seluruh API /api/utils/*
	api := app.Group("/api", user.AuthMiddleware(db))

	// 🛠️ Grup untuk utils (tooltips)
	utilsRoutes := api.Group("/utils")

	// 🎯 Inisialisasi Controller
	tooltipsController := utils.NewTooltipsController(db)

	// 🔥 Endpoint untuk tooltips
	utilsRoutes.Post("/get-tooltips", tooltipsController.GetTooltips)
	utilsRoutes.Post("/insert-tooltip", tooltipsController.InsertTooltip)
	utilsRoutes.Get("/get-all-tooltips", tooltipsController.GetAllTooltips) // 🔥 Mendapatkan semua tooltips
}
