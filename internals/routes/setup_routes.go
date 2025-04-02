package routes

import (
	"arabiya-syari-fiber/internals/routes/auth"
	"arabiya-syari-fiber/internals/routes/category"
	"arabiya-syari-fiber/internals/routes/donation"
	"arabiya-syari-fiber/internals/routes/user"
	"arabiya-syari-fiber/internals/routes/user_progress"
	report_user "arabiya-syari-fiber/internals/routes/user_report_quiz"
	"arabiya-syari-fiber/internals/routes/utils" // Add this line.

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register routes
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	auth.UserRoutes(app, db)
	category.CategoryRoutes(app, db)
	category.QuizzesRoutes(app, db)
	user.UsersProfileRoutes(app, db)
	donation.DonationRoutes(app, db)
	utils.UtilsRoutes(app, db) // Add this line.
	report_user.ReportUserRoutes(app, db)
	user_progress.UserPointLogRoutes(app, db)

}
