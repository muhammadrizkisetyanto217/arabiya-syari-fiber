package category

import (
	controllers "arabiya-syari-fiber/internals/controllers/category"
	authControllers "arabiya-syari-fiber/internals/controllers/user" // Import AuthMiddleware

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register routes
func CategoryRoutes(app *fiber.App, db *gorm.DB) {

	difficultyController := controllers.NewDifficultyController(db)
	apidifficulty := app.Group("/api/difficulties", authControllers.AuthMiddleware)
	apidifficulty.Get("/", difficultyController.GetDifficulties)
	apidifficulty.Get("/:id", difficultyController.GetDifficulty)
	apidifficulty.Post("/", difficultyController.CreateDifficulty)
	apidifficulty.Put("/:id", difficultyController.UpdateDifficulty)
	apidifficulty.Delete("/:id", difficultyController.DeleteDifficulty)

	categoryController := controllers.NewCategoryController(db)
	apicategory := app.Group("/api/categories", authControllers.AuthMiddleware)
	apicategory.Get("/", categoryController.GetCategories)
	apicategory.Get("/:id", categoryController.GetCategory)
	apicategory.Get("/difficulty/:difficulty_id", categoryController.GetCategoriesByDifficulty)
	apicategory.Post("/", categoryController.CreateCategory)
	apicategory.Put("/:id", categoryController.UpdateCategory)
	apicategory.Delete("/:id", categoryController.DeleteCategory)

	subcategoryController := controllers.NewSubcategoryController(db)
	apisubcategory := app.Group("/api/subcategories", authControllers.AuthMiddleware)
	apisubcategory.Get("/", subcategoryController.GetSubcategories)
	apisubcategory.Get("/:id", subcategoryController.GetSubcategory)
	apisubcategory.Get("/category/:category_id", subcategoryController.GetSubcategoriesByCategory)
	apisubcategory.Post("/", subcategoryController.CreateSubcategory)
	apisubcategory.Put("/:id", subcategoryController.UpdateSubcategory)
	apisubcategory.Delete("/:id", subcategoryController.DeleteSubcategory)

	themeOrLevelController := controllers.NewThemeOrLevelController(db)
	apithemesorlevels := app.Group("/api/themes-or-levels", authControllers.AuthMiddleware)
	apithemesorlevels.Get("/", themeOrLevelController.GetThemesOrLevels)
	apithemesorlevels.Get("/:id", themeOrLevelController.GetThemeOrLevel)
	apithemesorlevels.Post("/", themeOrLevelController.CreateThemeOrLevel)
	apithemesorlevels.Put("/:id", themeOrLevelController.UpdateThemeOrLevel)
	apithemesorlevels.Delete("/:id", themeOrLevelController.DeleteThemeOrLevel)

	// Unit Routes
	unitController := controllers.NewUnitController(db)
	apiunit := app.Group("/api/units", authControllers.AuthMiddleware)
	apiunit.Get("/", unitController.GetUnits)
	apiunit.Get("/:id", unitController.GetUnit)
	apiunit.Get("/themes_or_levels/:themesOrLevelId", unitController.GetUnitByThemesOrLevels)
	apiunit.Post("/", unitController.CreateUnit)
	apiunit.Put("/:id", unitController.UpdateUnit)
	apiunit.Delete("/:id", unitController.DeleteUnit)
}
