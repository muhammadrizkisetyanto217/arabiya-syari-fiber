package auth

import (
	controllers "arabiya-syari-fiber/internals/controllers/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UserRoutes to set up authentication routes
func UserRoutes(app *fiber.App, db *gorm.DB) {
	authController := controllers.NewAuthController(db)
	auth := app.Group("/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	userController := controllers.NewUserController(db)

	api := app.Group("/api/users", controllers.AuthMiddleware)
	api.Get("/", userController.GetUsers)
	api.Get("/:id", userController.GetUser)
	api.Post("/", userController.CreateUser)
	api.Put("/:id", userController.UpdateUser)
	api.Delete("/:id", userController.DeleteUser)
}
