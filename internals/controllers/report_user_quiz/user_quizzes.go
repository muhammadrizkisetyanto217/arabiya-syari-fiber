// package report_user_quiz

// import (
// 	"arabiya-syari-fiber/internals/models/report_user"
// 	"log"

// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/gorm"
// )

// type UserQuizzesController struct {
// 	DB *gorm.DB
// }

// // NewUserQuizzesController membuat instance controller
// func NewUserQuizzesController(db *gorm.DB) *UserQuizzesController {
// 	return &UserQuizzesController{DB: db}
// }

// // CreateOrUpdateUserQuiz menyimpan atau mengupdate hasil quiz user
// func (uqc *UserQuizzesController) CreateOrUpdateUserQuiz(c *fiber.Ctx) error {
// 	log.Println("Saving or updating user quiz result")

// 	var input report_user.UserQuizzesModel

// 	// Parse body JSON
// 	if err := c.BodyParser(&input); err != nil {
// 		log.Println("Error parsing request body:", err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	var userQuiz report_user.UserQuizzesModel
// 	result := uqc.DB.Where("user_id = ? AND quiz_id = ?", input.UserID, input.QuizID).First(&userQuiz)

// 	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
// 		// Jika data belum ada, insert baru
// 		userQuiz = report_user.UserQuizzesModel{
// 			UserID:          input.UserID,
// 			QuizID:          input.QuizID,
// 			Attempt:         1,
// 			PercentageGrade: input.PercentageGrade,
// 			TimeDuration:    input.TimeDuration,
// 			Point:           input.Point,
// 		}

// 		if err := uqc.DB.Create(&userQuiz).Error; err != nil {
// 			log.Println("Error saving user quiz result:", err)
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user quiz result"})
// 		}

// 		return c.JSON(fiber.Map{"message": "User quiz result saved successfully", "attempt": userQuiz.Attempt, "percentage_grade": userQuiz.PercentageGrade})
// 	}

// 	if result.Error != nil {
// 		log.Println("Error checking existing record:", result.Error)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process request"})
// 	}

// 	// Jika data sudah ada, lakukan update manual tanpa ON CONFLICT
// 	if input.PercentageGrade > userQuiz.PercentageGrade {
// 		userQuiz.PercentageGrade = input.PercentageGrade
// 		userQuiz.TimeDuration = input.TimeDuration
// 		userQuiz.Point = input.Point
// 	}

// 	// Selalu tingkatkan jumlah attempt
// 	userQuiz.Attempt += 1

// 	// Update data
// 	err := uqc.DB.Model(&userQuiz).Updates(userQuiz).Error
// 	if err != nil {
// 		log.Println("Error updating user quiz result:", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user quiz result"})
// 	}

// 	return c.JSON(fiber.Map{"message": "User quiz result updated successfully", "attempt": userQuiz.Attempt, "percentage_grade": userQuiz.PercentageGrade})
// }

// // GetUserQuizzes mendapatkan daftar hasil quiz user berdasarkan user_id
// func (uqc *UserQuizzesController) GetUserQuizzes(c *fiber.Ctx) error {
// 	log.Println("Fetching user quiz results")

// 	userID, err := c.ParamsInt("user_id")
// 	if err != nil {
// 		log.Println("Invalid user ID:", err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
// 	}

// 	var userQuizzes []report_user.UserQuizzesModel
// 	if err := uqc.DB.Where("user_id = ?", userID).Find(&userQuizzes).Error; err != nil {
// 		log.Println("Error fetching user quiz results:", err)
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user quiz results"})
// 	}

// 	return c.JSON(userQuizzes)
// }

package report_user_quiz

import (
	"arabiya-syari-fiber/internals/models/progress_user"
	"arabiya-syari-fiber/internals/models/report_user"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserQuizzesController struct {
	DB *gorm.DB
}

func NewUserQuizzesController(db *gorm.DB) *UserQuizzesController {
	return &UserQuizzesController{DB: db}
}

// Fungsi tambahan point berdasarkan attempt
func getAdditionalPoint(attempt int) int {
	switch attempt {
	case 1:
		return 20
	case 2:
		return 40
	default:
		return 10
	}
}

// CreateOrUpdateUserQuiz menyimpan atau mengupdate hasil quiz user
func (uqc *UserQuizzesController) CreateOrUpdateUserQuiz(c *fiber.Ctx) error {
	log.Println("Saving or updating user quiz result")

	var input report_user.UserQuizzesModel

	if err := c.BodyParser(&input); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var userQuiz report_user.UserQuizzesModel
	result := uqc.DB.Where("user_id = ? AND quiz_id = ?", input.UserID, input.QuizID).First(&userQuiz)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// Attempt pertama
		attempt := 1
		addPoint := getAdditionalPoint(attempt)

		userQuiz = report_user.UserQuizzesModel{
			UserID:          input.UserID,
			QuizID:          input.QuizID,
			Attempt:         attempt,
			PercentageGrade: input.PercentageGrade,
			TimeDuration:    input.TimeDuration,
			Point:           addPoint,
		}

		if err := uqc.DB.Create(&userQuiz).Error; err != nil {
			log.Println("Error saving user quiz result:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user quiz result"})
		}

		// Simpan log ke user_point_logs
		_ = uqc.DB.Create(&progress_user.UserPointLog{
			UserID:      userQuiz.UserID,
			Points:      addPoint,
			Description: fmt.Sprintf("Quiz ID: %d", userQuiz.QuizID),
		})

		return c.JSON(fiber.Map{
			"message":          "User quiz result saved successfully",
			"attempt":          userQuiz.Attempt,
			"percentage_grade": userQuiz.PercentageGrade,
			"point":            userQuiz.Point,
		})
	}

	if result.Error != nil {
		log.Println("Error checking existing record:", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process request"})
	}

	// Attempt ke-2, 3, dst
	userQuiz.Attempt += 1
	addPoint := getAdditionalPoint(userQuiz.Attempt)
	userQuiz.Point += addPoint

	// Update jika nilai baru lebih tinggi
	if input.PercentageGrade > userQuiz.PercentageGrade {
		userQuiz.PercentageGrade = input.PercentageGrade
		userQuiz.TimeDuration = input.TimeDuration
	}

	if err := uqc.DB.Model(&userQuiz).Updates(userQuiz).Error; err != nil {
		log.Println("Error updating user quiz result:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user quiz result"})
	}

	// Simpan log ke user_point_logs
	_ = uqc.DB.Create(&progress_user.UserPointLog{
		UserID:      userQuiz.UserID,
		Points:      addPoint,
		Description: fmt.Sprintf("Quiz ID: %d", userQuiz.QuizID),
	})

	return c.JSON(fiber.Map{
		"message":          "User quiz result updated successfully",
		"attempt":          userQuiz.Attempt,
		"percentage_grade": userQuiz.PercentageGrade,
		"point":            userQuiz.Point,
	})
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
