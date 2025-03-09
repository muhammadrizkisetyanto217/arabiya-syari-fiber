package category

import (
	"arabiya-syari-fiber/internals/models/category"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

type ThemeOrLevelController struct {
	DB *gorm.DB
}

func NewThemeOrLevelController(db *gorm.DB) *ThemeOrLevelController {
	return &ThemeOrLevelController{DB: db}
}

func (tc *ThemeOrLevelController) GetThemesOrLevels(c *fiber.Ctx) error {
	log.Println("Fetching all themes or levels")
	var themesOrLevels []category.ThemesOrLevelModel
	if err := tc.DB.Find(&themesOrLevels).Error; err != nil {
		log.Println("Error fetching themes or levels:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch themes or levels"})
	}
	return c.JSON(themesOrLevels)
}

func (tc *ThemeOrLevelController) GetThemeOrLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Fetching theme or level with ID:", id)
	var themeOrLevel category.ThemesOrLevelModel
	if err := tc.DB.First(&themeOrLevel, id).Error; err != nil {
		log.Println("Theme or level not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Theme or level not found"})
	}
	return c.JSON(themeOrLevel)
}

func (tc *ThemeOrLevelController) CreateThemeOrLevel(c *fiber.Ctx) error {
	log.Println("Creating a new theme or level")
	var themeOrLevel category.ThemesOrLevelModel
	if err := c.BodyParser(&themeOrLevel); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := tc.DB.Create(&themeOrLevel).Error; err != nil {
		log.Println("Error creating theme or level:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create theme or level"})
	}
	return c.Status(201).JSON(themeOrLevel)
}

func (tc *ThemeOrLevelController) UpdateThemeOrLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Updating theme or level with ID:", id)
	var themeOrLevel category.ThemesOrLevelModel
	if err := tc.DB.First(&themeOrLevel, id).Error; err != nil {
		log.Println("Theme or level not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Theme or level not found"})
	}

	if err := c.BodyParser(&themeOrLevel); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := tc.DB.Save(&themeOrLevel).Error; err != nil {
		log.Println("Error updating theme or level:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update theme or level"})
	}

	return c.JSON(themeOrLevel)
}

func (tc *ThemeOrLevelController) DeleteThemeOrLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Deleting theme or level with ID:", id)
	if err := tc.DB.Delete(&category.ThemesOrLevelModel{}, id).Error; err != nil {
		log.Println("Error deleting theme or level:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete theme or level"})
	}
	return c.JSON(fiber.Map{"message": "Theme or level deleted successfully"})
}
