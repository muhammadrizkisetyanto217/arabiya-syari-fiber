package routes

import (
	authRoutes "arabiya-syari-fiber/internals/routes/auth"
	categoryRoutes "arabiya-syari-fiber/internals/routes/category"
	DonationRoutes "arabiya-syari-fiber/internals/routes/donation"
	UserRoutes "arabiya-syari-fiber/internals/routes/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	authRoutes.UserRoutes(app, db)
	categoryRoutes.CategoryRoutes(app, db)
	categoryRoutes.QuizzesRoutes(app, db)
	UserRoutes.SetupRoutes(app, db)
	DonationRoutes.DonationRoutes(app, db)


}
