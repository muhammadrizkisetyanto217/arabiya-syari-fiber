package category

import (
	"arabiya-syari-fiber/internals/controllers/category"
	authControllers "arabiya-syari-fiber/internals/controllers/user" // Import AuthMiddleware

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register category-related routes
func CategoryRoutes(app *fiber.App, db *gorm.DB) {

	// 🔥 Proteksi seluruh kategori API dengan Middleware
	api := app.Group("/api", authControllers.AuthMiddleware)

	// 🎯 Difficulty Routes
	difficultyController := category.NewDifficultyController(db)
	difficultyRoutes := api.Group("/difficulties")
	difficultyRoutes.Get("/", difficultyController.GetDifficulties)
	difficultyRoutes.Get("/:id", difficultyController.GetDifficulty)
	difficultyRoutes.Post("/", difficultyController.CreateDifficulty)
	difficultyRoutes.Put("/:id", difficultyController.UpdateDifficulty)
	difficultyRoutes.Delete("/:id", difficultyController.DeleteDifficulty)

	// 🎯 Category Routes
	categoryController := category.NewCategoryController(db)
	categoryRoutes := api.Group("/categories")
	categoryRoutes.Get("/", categoryController.GetCategories)
	categoryRoutes.Get("/:id", categoryController.GetCategory)
	categoryRoutes.Get("/difficulty/:difficulty_id", categoryController.GetCategoriesByDifficulty)
	categoryRoutes.Post("/", categoryController.CreateCategory)
	categoryRoutes.Put("/:id", categoryController.UpdateCategory)
	categoryRoutes.Delete("/:id", categoryController.DeleteCategory)

	// 🎯 Subcategory Routes
	subcategoryController := category.NewSubcategoryController(db)
	subcategoryRoutes := api.Group("/subcategories")
	subcategoryRoutes.Get("/", subcategoryController.GetSubcategories)
	subcategoryRoutes.Get("/:id", subcategoryController.GetSubcategory)
	subcategoryRoutes.Get("/category/:category_id", subcategoryController.GetSubcategoriesByCategory)
	subcategoryRoutes.Post("/", subcategoryController.CreateSubcategory)
	subcategoryRoutes.Put("/:id", subcategoryController.UpdateSubcategory)
	subcategoryRoutes.Delete("/:id", subcategoryController.DeleteSubcategory)

	// 🎯 Themes or Levels Routes
	themeOrLevelController := category.NewThemeOrLevelController(db)
	themeOrLevelRoutes := api.Group("/themes-or-levels")
	themeOrLevelRoutes.Get("/", themeOrLevelController.GetThemesOrLevels)
	themeOrLevelRoutes.Get("/:id", themeOrLevelController.GetThemeOrLevel)
	themeOrLevelRoutes.Post("/", themeOrLevelController.CreateThemeOrLevel)
	themeOrLevelRoutes.Put("/:id", themeOrLevelController.UpdateThemeOrLevel)
	themeOrLevelRoutes.Delete("/:id", themeOrLevelController.DeleteThemeOrLevel)

	// 🎯 Unit Routes
	unitController := category.NewUnitController(db)
	unitRoutes := api.Group("/units")
	unitRoutes.Get("/", unitController.GetUnits)
	unitRoutes.Get("/:id", unitController.GetUnit)
	unitRoutes.Get("/themes-or-levels/:themesOrLevelId", unitController.GetUnitByThemesOrLevels) // ✅ Gunakan `-` bukan `_`
	unitRoutes.Post("/", unitController.CreateUnit)
	unitRoutes.Put("/:id", unitController.UpdateUnit)
	unitRoutes.Delete("/:id", unitController.DeleteUnit)
}
