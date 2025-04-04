package quizzes

import (
	"log"

	"arabiya-syari-fiber/internals/models/quizzes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type QuizController struct {
	DB *gorm.DB
}

func NewQuizController(db *gorm.DB) *QuizController {
	return &QuizController{DB: db}
}

// GET all quizzes
func (qc *QuizController) GetQuizzes(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all quizzes")
	var quizList []quizzes.QuizModel

	if err := qc.DB.Find(&quizList).Error; err != nil {
		log.Println("[ERROR] Failed to fetch quizzes:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch quizzes"})
	}

	log.Printf("[SUCCESS] Retrieved %d quizzes\n", len(quizList))
	return c.JSON(fiber.Map{
		"message": "Quizzes fetched successfully",
		"total":   len(quizList),
		"data":    quizList,
	})
}

// GET quiz by ID
func (qc *QuizController) GetQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching quiz with ID: %s\n", id)

	var quiz quizzes.QuizModel
	if err := qc.DB.First(&quiz, id).Error; err != nil {
		log.Println("[ERROR] Quiz not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Quiz not found"})
	}

	return c.JSON(fiber.Map{
		"message": "Quiz fetched successfully",
		"data":    quiz,
	})
}

// GET quizzes by section ID
func (qc *QuizController) GetQuizzesBySection(c *fiber.Ctx) error {
	sectionID := c.Params("sectionId")
	log.Printf("[INFO] Fetching quizzes for section_id: %s\n", sectionID)

	var quizzesBySection []quizzes.QuizModel
	if err := qc.DB.
		Joins("JOIN section_quizzes ON quizzes.section_quizzes_id = section_quizzes.id").
		Where("section_quizzes.id = ?", sectionID).
		Find(&quizzesBySection).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch quizzes for section_id %s: %v\n", sectionID, err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch quizzes"})
	}

	log.Printf("[SUCCESS] Retrieved %d quizzes for section_id %s\n", len(quizzesBySection), sectionID)
	return c.JSON(fiber.Map{
		"message": "Quizzes fetched successfully by section",
		"total":   len(quizzesBySection),
		"data":    quizzesBySection,
	})
}

// POST create quiz
func (qc *QuizController) CreateQuiz(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new quiz")

	var quiz quizzes.QuizModel
	if err := c.BodyParser(&quiz); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := qc.DB.Create(&quiz).Error; err != nil {
		log.Println("[ERROR] Failed to create quiz:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create quiz"})
	}

	log.Printf("[SUCCESS] Quiz created: ID=%d\n", quiz.ID)
	return c.Status(201).JSON(fiber.Map{
		"message": "Quiz created successfully",
		"data":    quiz,
	})
}

// PUT update quiz
func (qc *QuizController) UpdateQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Updating quiz with ID: %s\n", id)

	var quiz quizzes.QuizModel
	if err := qc.DB.First(&quiz, id).Error; err != nil {
		log.Println("[ERROR] Quiz not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Quiz not found"})
	}

	var requestData map[string]interface{}
	if err := c.BodyParser(&requestData); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := qc.DB.Model(&quiz).Updates(requestData).Error; err != nil {
		log.Println("[ERROR] Failed to update quiz:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update quiz"})
	}

	log.Printf("[SUCCESS] Quiz updated: ID=%s\n", id)
	return c.JSON(fiber.Map{
		"message": "Quiz updated successfully",
		"data":    quiz,
	})
}

// DELETE quiz
func (qc *QuizController) DeleteQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Deleting quiz with ID: %s\n", id)

	if err := qc.DB.Delete(&quizzes.QuizModel{}, id).Error; err != nil {
		log.Println("[ERROR] Failed to delete quiz:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete quiz"})
	}

	log.Printf("[SUCCESS] Quiz with ID %s deleted\n", id)
	return c.JSON(fiber.Map{
		"message": "Quiz deleted successfully",
	})
}
