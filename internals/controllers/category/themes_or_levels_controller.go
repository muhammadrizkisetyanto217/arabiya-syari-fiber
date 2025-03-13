package category

import (
	"arabiya-syari-fiber/internals/models/category"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ThemeOrLevelController struct {
	DB *gorm.DB
}

func NewThemeOrLevelController(db *gorm.DB) *ThemeOrLevelController {
	return &ThemeOrLevelController{DB: db}
}

func (tc *ThemeOrLevelController) GetThemeOrLevels(c *fiber.Ctx) error {
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

func (tc *ThemeOrLevelController) GetThemesOrLevelsBySubcategory(c *fiber.Ctx) error {
	subcategoryID := c.Params("subcategories_id")
	log.Printf("[INFO] Fetching themes or levels for subcategory ID: %s\n", subcategoryID)

	var themesOrLevels []category.ThemesOrLevelModel
	if err := tc.DB.Where("subcategories_id = ?", subcategoryID).Find(&themesOrLevels).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch themes or levels for subcategory ID %s: %v\n", subcategoryID, err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch themes or levels"})
	}

	log.Printf("[SUCCESS] Retrieved %d themes or levels for subcategory ID %s\n", len(themesOrLevels), subcategoryID)
	return c.JSON(themesOrLevels)
}

func (tc *ThemeOrLevelController) CreateThemeOrLevel(c *fiber.Ctx) error {
	log.Println("Creating a new theme or level")

	var themeOrLevel category.ThemesOrLevelModel

	// Parsing request body
	if err := c.BodyParser(&themeOrLevel); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// 🔥 Validasi data wajib diisi
	if themeOrLevel.Name == "" || themeOrLevel.Status == "" || themeOrLevel.DescriptionShort == "" || themeOrLevel.DescriptionLong == "" || themeOrLevel.SubcategoriesID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "All fields are required, including subcategories_id"})
	}

	// 🔥 Validasi `status` harus dalam daftar yang benar
	validStatuses := map[string]bool{"active": true, "pending": true, "archived": true}
	if !validStatuses[themeOrLevel.Status] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status. Allowed values: active, pending, archived"})
	}

	// 🔥 Pastikan `subcategories_id` ada di database
	var count int64
	if err := tc.DB.Table("subcategories").Where("id = ?", themeOrLevel.SubcategoriesID).Count(&count).Error; err != nil || count == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid subcategories_id. The referenced subcategory does not exist"})
	}

	// 🔥 Insert ke database
	if err := tc.DB.Create(&themeOrLevel).Error; err != nil {
		log.Println("Error creating theme or level:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create theme or level"})
	}

	// 🔥 Kembalikan response dengan format yang lebih baik
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Theme or Level created successfully",
		"data":    themeOrLevel,
	})
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
