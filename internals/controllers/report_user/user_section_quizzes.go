package report_user

import (
	"arabiya-syari-fiber/internals/models/report_user"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// "your_project/models"
)

// UserSectionQuizzesController adalah controller untuk user section quizzes
type UserSectionQuizzesController struct {
	DB *gorm.DB
}

// NewUserSectionQuizzesController membuat instance controller
func NewUserSectionQuizzesController(db *gorm.DB) *UserSectionQuizzesController {
	return &UserSectionQuizzesController{DB: db}
}

// GetUserSectionQuizzes mengambil daftar kuis yang telah diselesaikan per section
func (usc *UserSectionQuizzesController) GetUserSectionQuizzes(c *fiber.Ctx) error {
	log.Println("Fetching user section quizzes")

	userID, err := c.ParamsInt("user_id")
	if err != nil {
		log.Println("Invalid user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var userSections []report_user.UserSectionQuizzesModel
	if err := usc.DB.Where("user_id = ?", userID).Find(&userSections).Error; err != nil {
		log.Println("Error fetching user section quizzes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user section quizzes"})
	}

	return c.JSON(userSections)
}
