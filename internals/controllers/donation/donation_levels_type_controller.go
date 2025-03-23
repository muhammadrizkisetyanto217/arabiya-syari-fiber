package donation

import (
	"log"
	"time"

	"arabiya-syari-fiber/internals/models/donation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DonationLevelsController struct {
	DB *gorm.DB
}

func NewDonationLevelsController(db *gorm.DB) *DonationLevelsController {
	return &DonationLevelsController{DB: db}
}

// Get all donation levels
func (dlc *DonationLevelsController) GetAll(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all donation levels")

	var levels []donation.DonationLevelsTypeModel
	if err := dlc.DB.Find(&levels).Error; err != nil {
		log.Println("[ERROR] Failed to fetch donation levels:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch donation levels"})
	}

	log.Printf("[SUCCESS] Retrieved %d donation levels\n", len(levels))
	return c.JSON(fiber.Map{
		"message": "Donation levels fetched successfully",
		"total":   len(levels),
		"data":    levels,
	})
}

// Get donation level by ID
func (dlc *DonationLevelsController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("[ERROR] Invalid ID format:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Printf("[INFO] Fetching donation level with ID: %d\n", id)

	var level donation.DonationLevelsTypeModel
	if err := dlc.DB.First(&level, id).Error; err != nil {
		log.Println("[ERROR] Donation level not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Donation level not found"})
	}

	return c.JSON(fiber.Map{
		"message": "Donation level fetched successfully",
		"data":    level,
	})
}

// Create donation level
func (dlc *DonationLevelsController) Create(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new donation level")

	var level donation.DonationLevelsTypeModel
	if err := c.BodyParser(&level); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := dlc.DB.Create(&level).Error; err != nil {
		log.Println("[ERROR] Failed to create donation level:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create donation level"})
	}

	log.Printf("[SUCCESS] Donation level created: ID=%d\n", level.ID)
	return c.Status(201).JSON(fiber.Map{
		"message": "Donation level created successfully",
		"data":    level,
	})
}

// Update donation level
func (dlc *DonationLevelsController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("[ERROR] Invalid ID format:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Printf("[INFO] Updating donation level with ID: %d\n", id)

	var level donation.DonationLevelsTypeModel
	if err := dlc.DB.First(&level, id).Error; err != nil {
		log.Println("[ERROR] Donation level not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Donation level not found"})
	}

	var input donation.DonationLevelsTypeModel
	if err := c.BodyParser(&input); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := dlc.DB.Model(&level).Updates(input).Error; err != nil {
		log.Println("[ERROR] Failed to update donation level:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update donation level"})
	}

	log.Printf("[SUCCESS] Donation level updated: ID=%d\n", level.ID)
	return c.JSON(fiber.Map{
		"message": "Donation level updated successfully",
		"data":    level,
	})
}

// Soft delete donation level
func (dlc *DonationLevelsController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("[ERROR] Invalid ID format:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Printf("[INFO] Deleting donation level with ID: %d\n", id)

	if err := dlc.DB.Model(&donation.DonationLevelsTypeModel{}).
		Where("id = ?", id).
		Update("deleted_at", time.Now()).Error; err != nil {
		log.Println("[ERROR] Failed to delete donation level:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete donation level"})
	}

	log.Printf("[SUCCESS] Donation level with ID %d deleted\n", id)
	return c.JSON(fiber.Map{
		"message": "Donation level deleted successfully",
	})
}
