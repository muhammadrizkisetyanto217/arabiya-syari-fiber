package controllers

import (
	"log"

	models "arabiya-syari-fiber/internals/models/quizzes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type QuizController struct {
	DB *gorm.DB
}

func NewQuizController(db *gorm.DB) *QuizController {
	return &QuizController{DB: db}
}

// Get all quizzes
func (qc *QuizController) GetQuizzes(c *fiber.Ctx) error {
	log.Println("Fetching all quizzes")
	var quizzes []models.Quiz
	if err := qc.DB.Find(&quizzes).Error; err != nil {
		log.Println("Error fetching quizzes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch quizzes"})
	}
	return c.JSON(quizzes)
}

// Get quiz by ID
func (qc *QuizController) GetQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Fetching quiz with ID:", id)
	var quiz models.Quiz
	if err := qc.DB.First(&quiz, id).Error; err != nil {
		log.Println("Quiz not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz not found"})
	}
	return c.JSON(quiz)
}

func (qc *QuizController) GetQuizzesBySection(c *fiber.Ctx) error {
	sectionID := c.Params("sectionId")
	log.Printf("[INFO] Fetching quizzes for section_id: %s\n", sectionID)

	var quizzes []models.Quiz
	if err := qc.DB.
		Joins("JOIN section_quizzes ON quizzes.section_quizzes_id = section_quizzes.id").
		Where("section_quizzes.id = ?", sectionID).
		Find(&quizzes).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch quizzes for section_id %s: %v\n", sectionID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch quizzes"})
	}

	log.Printf("[SUCCESS] Retrieved %d quizzes for section_id %s\n", len(quizzes), sectionID)
	return c.JSON(quizzes)
}




// Create a new quiz
func (qc *QuizController) CreateQuiz(c *fiber.Ctx) error {
	log.Println("Creating a new quiz")
	var quiz models.Quiz
	if err := c.BodyParser(&quiz); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := qc.DB.Create(&quiz).Error; err != nil {
		log.Println("Error creating quiz:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create quiz"})
	}
	return c.Status(fiber.StatusCreated).JSON(quiz)
}

// Update an existing quiz
func (qc *QuizController) UpdateQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Updating quiz with ID:", id)

	var quiz models.Quiz
	if err := qc.DB.First(&quiz, id).Error; err != nil {
		log.Println("Quiz not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz not found"})
	}

	var requestData map[string]interface{}
	if err := c.BodyParser(&requestData); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := qc.DB.Model(&quiz).Updates(requestData).Error; err != nil {
		log.Println("Error updating quiz:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update quiz"})
	}

	return c.JSON(quiz)
}

// Delete a quiz
func (qc *QuizController) DeleteQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Deleting quiz with ID:", id)
	if err := qc.DB.Delete(&models.Quiz{}, id).Error; err != nil {
		log.Println("Error deleting quiz:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete quiz"})
	}
	return c.JSON(fiber.Map{"message": "Quiz deleted successfully"})
}
