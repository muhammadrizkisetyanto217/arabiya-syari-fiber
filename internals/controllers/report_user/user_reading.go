package report_user

import (
	"arabiya-syari-fiber/internals/models/report_user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserReadingController struct {
	DB *gorm.DB
}

func NewUserReadingController(db *gorm.DB) *UserReadingController {
	return &UserReadingController{DB: db}
}

// POST /user-readings
func (ctrl *UserReadingController) Create(c *fiber.Ctx) error {
	var input report_user.UserReading

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	if err := ctrl.DB.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user reading",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(input)
}

// GET /user-readings
func (ctrl *UserReadingController) GetAll(c *fiber.Ctx) error {
	var readings []report_user.UserReading

	if err := ctrl.DB.Find(&readings).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch user readings",
		})
	}

	return c.JSON(readings)
}
