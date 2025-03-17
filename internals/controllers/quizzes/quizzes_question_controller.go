package quizzes

import (
	"log"

	"arabiya-syari-fiber/internals/models/quizzes"
	"arabiya-syari-fiber/internals/models/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq" // Untuk TEXT[] dalam PostgreSQL
	"gorm.io/gorm"
)

type QuizzesQuestionController struct {
	DB *gorm.DB
}

func NewQuizzesQuestionController(db *gorm.DB) *QuizzesQuestionController {
	return &QuizzesQuestionController{DB: db}
}

// Get all quiz questions
func (qqc *QuizzesQuestionController) GetQuizzesQuestions(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all quiz questions")
	var questions []quizzes.QuizQuestionModel
	if err := qqc.DB.Find(&questions).Error; err != nil {
		log.Println("[ERROR] Failed to fetch quiz questions:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch quiz questions"})
	}
	return c.JSON(questions)
}

// Get quiz question by ID
func (qqc *QuizzesQuestionController) GetQuizzesQuestion(c *fiber.Ctx) error {
	id := c.Params("id") // ✅ Pastikan `id` hanya dideklarasikan sekali
	log.Println("[INFO] Fetching quiz question with ID:", id)

	var question quizzes.QuizQuestionModel
	err := qqc.DB.First(&question, id).Error
	if err != nil {
		log.Println("[ERROR] Quiz question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz question not found"})
	}

	return c.JSON(question)
}

// Get quiz questions by quiz ID
func (qqc *QuizzesQuestionController) GetQuizzesQuestionsByQuizID(c *fiber.Ctx) error {
	quizID := c.Params("quizzedId") // ✅ Pastikan `quizID` hanya dideklarasikan sekali
	log.Printf("[INFO] Fetching questions for quiz_id: %s\n", quizID)

	var questions []quizzes.QuizQuestionModel
	err := qqc.DB.Where("quizzes_id = ?", quizID).Find(&questions).Error
	if err != nil {
		log.Printf("[ERROR] Failed to fetch questions for quiz_id %s: %v\n", quizID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch questions"})
	}

	log.Printf("[SUCCESS] Retrieved %d questions for quiz_id %s\n", len(questions), quizID)
	return c.JSON(questions)
}

// Create a new quiz question
func (qqc *QuizzesQuestionController) CreateQuizzesQuestion(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new quiz question")

	var question quizzes.QuizQuestionModel
	if err := c.BodyParser(&question); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Pastikan nilai array dikonversi dengan benar
	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := qqc.DB.Create(&question).Error; err != nil {
		log.Println("[ERROR] Failed to create quiz question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create quiz question"})
	}

	return c.Status(fiber.StatusCreated).JSON(question)
}

// Update an existing quiz question
func (qqc *QuizzesQuestionController) UpdateQuizzesQuestion(c *fiber.Ctx) error {
	id := c.Params("id") // ✅ Pastikan `id` hanya dideklarasikan sekali
	log.Println("[INFO] Updating quiz question with ID:", id)

	var question quizzes.QuizQuestionModel
	err := qqc.DB.First(&question, id).Error
	if err != nil {
		log.Println("[ERROR] Quiz question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz question not found"})
	}

	if err := c.BodyParser(&question); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Pastikan array tetap dalam format `pq.StringArray`
	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := qqc.DB.Save(&question).Error; err != nil {
		log.Println("[ERROR] Failed to update quiz question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update quiz question"})
	}

	return c.JSON(question)
}

// Delete a quiz question
func (qqc *QuizzesQuestionController) DeleteQuizzesQuestion(c *fiber.Ctx) error {
	id := c.Params("id") // ✅ Pastikan `id` hanya dideklarasikan sekali
	log.Println("[INFO] Deleting quiz question with ID:", id)

	err := qqc.DB.Delete(&quizzes.QuizQuestionModel{}, id).Error
	if err != nil {
		log.Println("[ERROR] Failed to delete quiz question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete quiz question"})
	}

	return c.JSON(fiber.Map{"message": "Quiz question deleted successfully"})
}

// GetQuizQuestionWithTooltips mengambil quiz question beserta tooltips berdasarkan ID
func (qc *QuizzesQuestionController) GetQuizzesQuestionWithTooltips(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching quiz question with ID: %s\n", id)

	var quizQuestion quizzes.QuizQuestionModel
	if err := qc.DB.First(&quizQuestion, id).Error; err != nil {
		log.Println("[ERROR] Quiz question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz question not found"})
	}

	var tooltips []utils.Tooltip
	if len(quizQuestion.TooltipsID) > 0 {
		if err := qc.DB.Where("id = ANY(?)", pq.Array(quizQuestion.TooltipsID)).Find(&tooltips).Error; err != nil {
			log.Println("[ERROR] Failed to fetch tooltips:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tooltips"})
		}
	}

	log.Printf("[SUCCESS] Retrieved quiz question with ID: %s\n", id)
	return c.JSON(fiber.Map{
		"quiz_question": quizQuestion,
		"tooltips":      tooltips,
	})
}

// GetOnlyQuizQuestionTooltips hanya mengambil tooltips berdasarkan ID quiz question
func (qc *QuizzesQuestionController) GetOnlyQuizzesQuestionTooltips(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching only tooltips for quiz question ID: %s\n", id)

	var quizQuestion quizzes.QuizQuestionModel
	if err := qc.DB.First(&quizQuestion, id).Error; err != nil {
		log.Println("[ERROR] Quiz question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz question not found"})
	}

	var tooltips []utils.Tooltip
	if len(quizQuestion.TooltipsID) > 0 {
		if err := qc.DB.Where("id = ANY(?)", pq.Array(quizQuestion.TooltipsID)).Find(&tooltips).Error; err != nil {
			log.Println("[ERROR] Failed to fetch tooltips:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch tooltips"})
		}
	}

	log.Printf("[SUCCESS] Retrieved only tooltips for quiz question ID: %s\n", id)
	return c.JSON(fiber.Map{
		"tooltips": tooltips,
	})
}
