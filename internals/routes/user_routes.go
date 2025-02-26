package routes

import (
	authRoutes "arabiya-syari-fiber/internals/routes/auth"
	categoryRoutes "arabiya-syari-fiber/internals/routes/category"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register routes
func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	authRoutes.UserRoutes(app, db)
	categoryRoutes.CategoryRoutes(app, db)
	categoryRoutes.QuizzesRoutes(app, db)

}
