package auth

import (
	"arabiya-syari-fiber/internals/controllers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserRoutes to set up authentication routes
func UserRoutes(app *fiber.App, db *gorm.DB) {
	authController := controllers.NewAuthController(db)
	auth := app.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
}
