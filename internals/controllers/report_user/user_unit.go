package report_user

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"arabiya-syari-fiber/internals/models/report_user"
)

type UserUnitController struct {
	DB *gorm.DB
}

func NewUserUnitController(db *gorm.DB) *UserUnitController {
	return &UserUnitController{DB: db}
}

// GET /api/user_units/:user_id
func (ctrl *UserUnitController) GetUserUnitsByUserID(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}

	var units []report_user.UserUnitModel
	if err := ctrl.DB.
		Where("user_id = ?", userID).
		Find(&units).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch user units"})
	}

	return c.JSON(units)
}
