package controllers

import (
	"arabiya-syari-fiber/internals/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController(db *gorm.DB) *CategoryController {
	return &CategoryController{DB: db}
}

// Get all categories
func (cc *CategoryController) GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	if err := cc.DB.Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch categories"})
	}
	return c.JSON(categories)
}

// Get single category
func (cc *CategoryController) GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := cc.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}
	return c.JSON(category)
}

// Get categories by difficulty_id
func (cc *CategoryController) GetCategoriesByDifficulty(c *fiber.Ctx) error {
	difficultyID := c.Params("difficulty_id")
	var categories []models.Category

	if err := cc.DB.Where("difficulty_id = ?", difficultyID).Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch categories"})
	}

	return c.JSON(categories)
}

// Create category
func (cc *CategoryController) CreateCategory(c *fiber.Ctx) error {
	var category models.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := cc.DB.Create(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create category"})
	}
	return c.Status(201).JSON(category)
}

// Update category
func (cc *CategoryController) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	var category models.Category
	if err := cc.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
	}

	if err := c.BodyParser(&category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	cc.DB.Save(&category)
	return c.JSON(category)
}

// Delete category
func (cc *CategoryController) DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := cc.DB.Delete(&models.Category{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete category"})
	}
	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}
