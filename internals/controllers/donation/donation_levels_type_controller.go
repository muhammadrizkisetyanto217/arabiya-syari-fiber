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
	log.Println("Fetching all donation levels")
	var levels []models.DonationLevelsTypeModel
	if err := dlc.DB.Find(&levels).Error; err != nil {
		log.Println("Error fetching donation levels:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch donation levels"})
	}
	return c.JSON(levels)
}

// Get a specific donation level by ID
func (dlc *DonationLevelsController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Fetching donation level with ID:", id)
	var level models.DonationLevelsTypeModel
	if err := dlc.DB.First(&level, id).Error; err != nil {
		log.Println("Donation level not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Donation level not found"})
	}
	return c.JSON(level)
}

// Create a new donation level
func (dlc *DonationLevelsController) Create(c *fiber.Ctx) error {
	log.Println("Creating a new donation level")
	var level models.DonationLevelsTypeModel

	if err := c.BodyParser(&level); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := dlc.DB.Create(&level).Error; err != nil {
		log.Println("Error creating donation level:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create donation level"})
	}
	return c.Status(fiber.StatusCreated).JSON(level)
}

// Update a donation level
func (dlc *DonationLevelsController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Updating donation level with ID:", id)
	var level models.DonationLevelsTypeModel
	if err := dlc.DB.First(&level, id).Error; err != nil {
		log.Println("Donation level not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Donation level not found"})
	}

	var input models.DonationLevelsTypeModel
	if err := c.BodyParser(&input); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Update hanya field yang dikirim
	if err := dlc.DB.Model(&level).Updates(input).Error; err != nil {
		log.Println("Error updating donation level:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update donation level"})
	}
	return c.JSON(level)
}

// Soft delete a donation level
func (dlc *DonationLevelsController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Deleting donation level with ID:", id)
	if err := dlc.DB.Model(&models.DonationLevelsTypeModel{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		log.Println("Error deleting donation level:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete donation level"})
	}
	return c.JSON(fiber.Map{"message": "Donation level deleted successfully"})
}
