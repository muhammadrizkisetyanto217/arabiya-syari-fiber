package report_user

import (
	"arabiya-syari-fiber/internals/models/report_user"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserQuizzesController struct {
	DB *gorm.DB
}

// NewUserQuizzesController membuat instance controller
func NewUserQuizzesController(db *gorm.DB) *UserQuizzesController {
	return &UserQuizzesController{DB: db}
}

// CreateOrUpdateUserQuiz menyimpan atau mengupdate hasil quiz user
func (uqc *UserQuizzesController) CreateOrUpdateUserQuiz(c *fiber.Ctx) error {
	log.Println("Saving or updating user quiz result")

	var input report_user.UserQuizzesModel

	// Parse body JSON
	if err := c.BodyParser(&input); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var userQuiz report_user.UserQuizzesModel
	result := uqc.DB.Where("user_id = ? AND quiz_id = ?", input.UserID, input.QuizID).First(&userQuiz)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// Jika data belum ada, insert baru
		userQuiz = report_user.UserQuizzesModel{
			UserID:          input.UserID,
			QuizID:          input.QuizID,
			Attempt:         1,
			PercentageGrade: input.PercentageGrade,
			TimeDuration:    input.TimeDuration,
			Point:           input.Point,
		}

		if err := uqc.DB.Create(&userQuiz).Error; err != nil {
			log.Println("Error saving user quiz result:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user quiz result"})
		}

		return c.JSON(fiber.Map{"message": "User quiz result saved successfully", "attempt": userQuiz.Attempt, "percentage_grade": userQuiz.PercentageGrade})
	}

	if result.Error != nil {
		log.Println("Error checking existing record:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process request"})
	}

	// Jika data sudah ada, lakukan update manual tanpa ON CONFLICT
	if input.PercentageGrade > userQuiz.PercentageGrade {
		userQuiz.PercentageGrade = input.PercentageGrade
		userQuiz.TimeDuration = input.TimeDuration
		userQuiz.Point = input.Point
	}

	// Selalu tingkatkan jumlah attempt
	userQuiz.Attempt += 1

	// Update data
	err := uqc.DB.Model(&userQuiz).Updates(userQuiz).Error
	if err != nil {
		log.Println("Error updating user quiz result:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user quiz result"})
	}

	return c.JSON(fiber.Map{"message": "User quiz result updated successfully", "attempt": userQuiz.Attempt, "percentage_grade": userQuiz.PercentageGrade})
}

// GetUserQuizzes mendapatkan daftar hasil quiz user berdasarkan user_id
func (uqc *UserQuizzesController) GetUserQuizzes(c *fiber.Ctx) error {
	log.Println("Fetching user quiz results")

	userID, err := c.ParamsInt("user_id")
	if err != nil {
		log.Println("Invalid user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var userQuizzes []report_user.UserQuizzesModel
	if err := uqc.DB.Where("user_id = ?", userID).Find(&userQuizzes).Error; err != nil {
		log.Println("Error fetching user quiz results:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user quiz results"})
	}

	return c.JSON(userQuizzes)
}
