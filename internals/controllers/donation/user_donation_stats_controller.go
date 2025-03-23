package donation

import (
	"log"

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
	log.Println("[INFO] Fetching all user donation logs")

	var logs []donation.UserDonationLogModel
	if err := udlc.DB.Find(&logs).Error; err != nil {
		log.Println("[ERROR] Failed to fetch donation logs:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch donation logs"})
	}

	log.Printf("[SUCCESS] Retrieved %d donation logs\n", len(logs))
	return c.JSON(fiber.Map{
		"message": "Donation logs fetched successfully",
		"total":   len(logs),
		"data":    logs,
	})
}

// Get donation log by ID
func (udlc *UserDonationLogsController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("[ERROR] Invalid ID format:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Printf("[INFO] Fetching donation log with ID: %d\n", id)

	var logEntry donation.UserDonationLogModel
	if err := udlc.DB.First(&logEntry, id).Error; err != nil {
		log.Println("[ERROR] Donation log not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Donation log not found"})
	}

	return c.JSON(fiber.Map{
		"message": "Donation log fetched successfully",
		"data":    logEntry,
	})
}

// Create new donation log
func (udlc *UserDonationLogsController) Create(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new user donation log")

	var logEntry donation.UserDonationLogModel
	if err := c.BodyParser(&logEntry); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := udlc.DB.Create(&logEntry).Error; err != nil {
		log.Println("[ERROR] Failed to create donation log:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create donation log"})
	}

	log.Printf("[SUCCESS] Donation log created: ID=%d\n", logEntry.ID)
	return c.Status(201).JSON(fiber.Map{
		"message": "Donation log created successfully",
		"data":    logEntry,
	})
}

// Update donation log
func (udlc *UserDonationLogsController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("[ERROR] Invalid ID format:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Printf("[INFO] Updating donation log with ID: %d\n", id)

	var logEntry donation.UserDonationLogModel
	if err := udlc.DB.First(&logEntry, id).Error; err != nil {
		log.Println("[ERROR] Donation log not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Donation log not found"})
	}

	var input donation.UserDonationLogModel
	if err := c.BodyParser(&input); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := udlc.DB.Model(&logEntry).Updates(input).Error; err != nil {
		log.Println("[ERROR] Failed to update donation log:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update donation log"})
	}

	log.Printf("[SUCCESS] Donation log updated: ID=%d\n", logEntry.ID)
	return c.JSON(fiber.Map{
		"message": "Donation log updated successfully",
		"data":    logEntry,
	})
}

// Permanently delete donation log
func (udlc *UserDonationLogsController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("[ERROR] Invalid ID format:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Printf("[INFO] Permanently deleting donation log with ID: %d\n", id)

	if err := udlc.DB.Unscoped().Where("id = ?", id).Delete(&donation.UserDonationLogModel{}).Error; err != nil {
		log.Println("[ERROR] Failed to delete donation log:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete donation log"})
	}

	log.Printf("[SUCCESS] Donation log with ID %d deleted\n", id)
	return c.JSON(fiber.Map{
		"message": "Donation log deleted successfully",
	})
}
