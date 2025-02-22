package controllers

import (
	"arabiya-syari-fiber/internals/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DifficultyController struct {
	DB *gorm.DB
}

func NewDifficultyController(db *gorm.DB) *DifficultyController {
	return &DifficultyController{DB: db}
}

// Get all difficulties
func (dc *DifficultyController) GetDifficulties(c *fiber.Ctx) error {
	var difficulties []models.Difficulty
	if err := dc.DB.Find(&difficulties).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(difficulties)
}

// Get difficulty by ID
func (dc *DifficultyController) GetDifficulty(c *fiber.Ctx) error {
	id := c.Params("id")
	var difficulty models.Difficulty
	if err := dc.DB.First(&difficulty, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Difficulty not found"})
	}
	return c.JSON(difficulty)
}

// Create difficulty
func (dc *DifficultyController) CreateDifficulty(c *fiber.Ctx) error {
	difficulty := new(models.Difficulty)
	if err := c.BodyParser(difficulty); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := dc.DB.Create(difficulty).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(difficulty)
}

// Update difficulty
func (dc *DifficultyController) UpdateDifficulty(c *fiber.Ctx) error {
	id := c.Params("id")
	var difficulty models.Difficulty

	if err := dc.DB.First(&difficulty, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Difficulty not found"})
	}

	if err := c.BodyParser(&difficulty); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := dc.DB.Save(&difficulty).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(difficulty)
}

// Delete difficulty
func (dc *DifficultyController) DeleteDifficulty(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := dc.DB.Delete(&models.Difficulty{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Difficulty deleted"})
}
