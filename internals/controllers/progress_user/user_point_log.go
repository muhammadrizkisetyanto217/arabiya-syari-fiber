package progress_user

import (
	"arabiya-syari-fiber/internals/models/progress_user"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserPointLogController struct {
	DB *gorm.DB
}

func NewUserPointLogController(db *gorm.DB) *UserPointLogController {
	return &UserPointLogController{DB: db}
}

// GetUserPointLogs mengambil semua log poin berdasarkan user_id
func (ctrl *UserPointLogController) GetUserPointLogs(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		log.Println("Invalid user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var logs []progress_user.UserPointLog
	if err := ctrl.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&logs).Error; err != nil {
		log.Println("Gagal mengambil user_point_logs:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Gagal mengambil data log poin",
		})
	}

	return c.JSON(logs)
}
