package category

import (
	// "arabiya-syari-fiber/internals/models"
	// "arabiya-syari-fiber/internals/database"
	"arabiya-syari-fiber/internals/controllers"

	// db "arabiya-syari-fiber/internals/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register routes
func CategoryRoutes(app *fiber.App, db *gorm.DB) {

	// userController := controllers.NewUserController(db)

	// api := app.Group("/api/users", controllers.AuthMiddleware)
	// api.Get("/", userController.GetUsers)
	// api.Get("/:id", userController.GetUser)
	// api.Post("/", userController.CreateUser)
	// api.Put("/:id", userController.UpdateUser)
	// api.Delete("/:id", userController.DeleteUser)


	difficultyController := controllers.NewDifficultyController(db)

	apidifficulty := app.Group("/api/difficulties")
	apidifficulty.Get("/", difficultyController.GetDifficulties)
	apidifficulty.Get("/:id", difficultyController.GetDifficulty)
	apidifficulty.Post("/", difficultyController.CreateDifficulty)
	apidifficulty.Put("/:id", difficultyController.UpdateDifficulty)
	apidifficulty.Delete("/:id", difficultyController.DeleteDifficulty)


	categoryController := controllers.NewCategoryController(db)
	apicategory := app.Group("/api/categories")
	apicategory.Get("/", categoryController.GetCategories)
	apicategory.Get("/:id", categoryController.GetCategory)
	apicategory.Get("/difficulty/:difficulty_id", categoryController.GetCategoriesByDifficulty)
	apicategory.Post("/", categoryController.CreateCategory)
	apicategory.Put("/:id", categoryController.UpdateCategory)
	apicategory.Delete("/:id", categoryController.DeleteCategory)
}
