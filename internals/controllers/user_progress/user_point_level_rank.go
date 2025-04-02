package user_progress

import (
	"arabiya-syari-fiber/internals/models/progress_user"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserPointLevelRankController struct {
	DB *gorm.DB
}

func NewUserPointLevelRankController(db *gorm.DB) *UserPointLevelRankController {
	return &UserPointLevelRankController{DB: db}
}

func (ctrl *UserPointLevelRankController) GetUserPointLevelRankByUserID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		log.Println("[ERROR] Invalid user_id:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	var rank progress_user.UserPointLevelRank
	if err := ctrl.DB.Where("user_id = ?", userID).First(&rank).Error; err != nil {
		log.Println("[ERROR] Data not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User point level rank not found",
		})
	}

	return c.JSON(rank)
}
