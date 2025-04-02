package user_progress

import (
	"log"

	"arabiya-syari-fiber/internals/models/progress_user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LevelPointRequirementController struct {
	DB *gorm.DB
}

func NewLevelPointRequirementController(db *gorm.DB) *LevelPointRequirementController {
	return &LevelPointRequirementController{DB: db}
}

// Get all levels
func (lc *LevelPointRequirementController) GetLevels(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all levels")
	var levels []progress_user.LevelPointRequirement
	if err := lc.DB.Find(&levels).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch levels: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch levels"})
	}
	log.Printf("[SUCCESS] Retrieved %d levels\n", len(levels))
	return c.JSON(levels)
}

// Get level by ID
func (lc *LevelPointRequirementController) GetLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching level with ID: %s\n", id)

	var level progress_user.LevelPointRequirement
	if err := lc.DB.First(&level, id).Error; err != nil {
		log.Printf("[ERROR] Level with ID %s not found\n", id)
		return c.Status(404).JSON(fiber.Map{"error": "Level not found"})
	}
	log.Printf("[SUCCESS] Retrieved level: ID=%s, Name=%s\n", id, level.NameLevel)
	return c.JSON(level)
}

// // Create level
// func (lc *LevelPointRequirementController) CreateLevel(c *fiber.Ctx) error {
//     log.Println("[INFO] Received request to create level")

//     var level progress_user.LevelPointRequirement
//     if err := c.BodyParser(&level); err != nil {
//         log.Printf("[ERROR] Invalid input: %v\n", err)
//         return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
//     }

//     if err := lc.DB.Create(&level).Error; err != nil {
//         log.Printf("[ERROR] Failed to create level: %v\n", err)
//         return c.Status(500).JSON(fiber.Map{"error": "Failed to create level"})
//     }
//     log.Printf("[SUCCESS] Level created: ID=%d, Name=%s\n", level.ID, level.NameLevel)
//     return c.Status(201).JSON(level)
// }

// Create level (support single or multiple input)
func (lc *LevelPointRequirementController) CreateLevel(c *fiber.Ctx) error {
	log.Println("[INFO] Received request to create level(s)")

	var levels []progress_user.LevelPointRequirement

	// Coba parsing langsung sebagai array
	if err := c.BodyParser(&levels); err != nil {
		log.Printf("[WARN] Input bukan array, coba parse sebagai single object\n")

		var single progress_user.LevelPointRequirement
		if err := c.BodyParser(&single); err != nil {
			log.Printf("[ERROR] Invalid input format: %v\n", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input format"})
		}

		levels = append(levels, single)
	}

	if len(levels) == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "No level data provided"})
	}

	// Simpan semua data level
	if err := lc.DB.Create(&levels).Error; err != nil {
		log.Printf("[ERROR] Failed to create levels: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create levels"})
	}

	log.Printf("[SUCCESS] Created %d level(s)\n", len(levels))
	return c.Status(201).JSON(levels)
}

// Update level
func (lc *LevelPointRequirementController) UpdateLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Updating level with ID: %s\n", id)

	var level progress_user.LevelPointRequirement
	if err := lc.DB.First(&level, id).Error; err != nil {
		log.Printf("[ERROR] Level with ID %s not found\n", id)
		return c.Status(404).JSON(fiber.Map{"error": "Level not found"})
	}

	if err := c.BodyParser(&level); err != nil {
		log.Printf("[ERROR] Invalid input: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := lc.DB.Save(&level).Error; err != nil {
		log.Printf("[ERROR] Failed to update level: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update level"})
	}

	log.Printf("[SUCCESS] Level updated: ID=%s, Name=%s\n", id, level.NameLevel)
	return c.JSON(level)
}

// Delete level
func (lc *LevelPointRequirementController) DeleteLevel(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Deleting level with ID: %s\n", id)

	if err := lc.DB.Delete(&progress_user.LevelPointRequirement{}, id).Error; err != nil {
		log.Printf("[ERROR] Failed to delete level: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete level"})
	}

	log.Printf("[SUCCESS] Level with ID %s deleted successfully\n", id)
	return c.JSON(fiber.Map{"message": "Level deleted successfully"})
}
