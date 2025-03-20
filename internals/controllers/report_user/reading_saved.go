package report_user

import (
	"arabiya-syari-fiber/internals/models/report_user"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ReadingSavedController struct {
	DB *gorm.DB
}

// NewReadingSavedController membuat instance controller
func NewReadingSavedController(db *gorm.DB) *ReadingSavedController {
	return &ReadingSavedController{DB: db}
}

// SaveReading menyimpan reading ke daftar saved user
func (rc *ReadingSavedController) SaveReading(c *fiber.Ctx) error {
	log.Println("Saving a reading for user")

	var readingSaved report_user.ReadingSavedModel

	if err := c.BodyParser(&readingSaved); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := rc.DB.Create(&readingSaved).Error; err != nil {
		log.Println("Error saving reading:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save reading"})
	}

	return c.JSON(fiber.Map{"message": "Reading saved successfully"})
}

// UnsaveReading menghapus reading dari daftar saved user
func (rc *ReadingSavedController) UnsaveReading(c *fiber.Ctx) error {
	log.Println("Removing a saved reading")

	var readingSaved report_user.ReadingSavedModel

	if err := c.BodyParser(&readingSaved); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := rc.DB.Where("user_id = ? AND reading_id = ?", readingSaved.UserID, readingSaved.ReadingID).
		Delete(&report_user.ReadingSavedModel{}).Error; err != nil {
		log.Println("Error removing saved reading:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to unsave reading"})
	}

	return c.JSON(fiber.Map{"message": "Reading unsaved successfully"})
}

// GetSavedReadings mendapatkan daftar reading yang disimpan user
func (rc *ReadingSavedController) GetSavedReadings(c *fiber.Ctx) error {
	log.Println("Fetching saved readings for user")

	userID, err := c.ParamsInt("user_id")
	if err != nil {
		log.Println("Invalid user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var savedReadings []report_user.ReadingSavedModel
	if err := rc.DB.Where("user_id = ?", userID).Find(&savedReadings).Error; err != nil {
		log.Println("Error fetching saved readings:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve saved readings"})
	}

	return c.JSON(savedReadings)
}
