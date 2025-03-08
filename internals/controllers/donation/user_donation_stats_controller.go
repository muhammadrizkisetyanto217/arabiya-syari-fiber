package donation

import (
	"log"
	"time"

	"arabiya-syari-fiber/internals/models/donation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserDonationLogsController struct {
	DB *gorm.DB
}

func NewUserDonationLogsController(db *gorm.DB) *UserDonationLogsController {
	return &UserDonationLogsController{DB: db}
}

// Get all donation logs
func (udlc *UserDonationLogsController) GetAll(c *fiber.Ctx) error {
	log.Println("Fetching all user donation logs")
	var logs []models.UserDonationLogModel
	if err := udlc.DB.Find(&logs).Error; err != nil {
		log.Println("Error fetching donation logs:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch donation logs"})
	}
	return c.JSON(logs)
}

// Get a specific donation log by ID
func (udlc *UserDonationLogsController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Fetching donation log with ID:", id)
	var logEntry models.UserDonationLogModel
	if err := udlc.DB.First(&logEntry, id).Error; err != nil {
		log.Println("Donation log not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Donation log not found"})
	}
	return c.JSON(logEntry)
}

// Create a new donation log
func (udlc *UserDonationLogsController) Create(c *fiber.Ctx) error {
	log.Println("Creating a new user donation log")
	var logEntry models.UserDonationLogModel

	if err := c.BodyParser(&logEntry); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := udlc.DB.Create(&logEntry).Error; err != nil {
		log.Println("Error creating donation log:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create donation log"})
	}
	return c.Status(fiber.StatusCreated).JSON(logEntry)
}

// Update a donation log
func (udlc *UserDonationLogsController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Updating donation log with ID:", id)
	var logEntry models.UserDonationLogModel
	if err := udlc.DB.First(&logEntry, id).Error; err != nil {
		log.Println("Donation log not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Donation log not found"})
	}

	var input models.UserDonationLogModel
	if err := c.BodyParser(&input); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := udlc.DB.Model(&logEntry).Updates(input).Error; err != nil {
		log.Println("Error updating donation log:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update donation log"})
	}
	return c.JSON(logEntry)
}

// Soft delete a donation log
func (udlc *UserDonationLogsController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Deleting donation log with ID:", id)
	if err := udlc.DB.Model(&models.UserDonationLogModel{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		log.Println("Error deleting donation log:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete donation log"})
	}
	return c.JSON(fiber.Map{"message": "Donation log deleted successfully"})
}
