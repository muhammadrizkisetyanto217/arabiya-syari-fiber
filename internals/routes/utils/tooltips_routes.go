package routes

import (
	tooltips "arabiya-syari-fiber/internals/controllers/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes menghubungkan routes dengan controller
func UtilsRoutes(app *fiber.App, db *gorm.DB) {
	tooltipsController := tooltips.NewTooltipsController(db)

	app.Post("/get-tooltips", tooltipsController.GetTooltips)
	app.Post("/insert-tooltip", tooltipsController.InsertTooltip)
	app.Get("/get-all-tooltips", tooltipsController.GetAllTooltips) // 🔥 Endpoint untuk mendapatkan semua tooltips
}
 