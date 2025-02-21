package routes

import (
	// "arabiya-syari-fiber/internals/models"
	// "arabiya-syari-fiber/internals/database"
	"arabiya-syari-fiber/internals/controllers"
	// db "arabiya-syari-fiber/internals/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register routes
func SetupUserRoutes(app *fiber.App, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	api := app.Group("/api/users")
	api.Get("/", userController.GetUsers)
	api.Get("/:id", userController.GetUser)
	api.Post("/", userController.CreateUser)
	api.Put("/:id", userController.UpdateUser)
	api.Delete("/:id", userController.DeleteUser)
}
