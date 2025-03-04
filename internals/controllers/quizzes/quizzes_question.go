package controllers

import (
	"log"

	models "arabiya-syari-fiber/internals/models/quizzes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type QuizQuestionController struct {
	DB *gorm.DB
}

func NewQuizQuestionController(db *gorm.DB) *QuizQuestionController {
	return &QuizQuestionController{DB: db}
}

// Get all quiz questions
func (qqc *QuizQuestionController) GetQuizQuestions(c *fiber.Ctx) error {
	log.Println("Fetching all quiz questions")
	var questions []models.QuizQuestion
	if err := qqc.DB.Find(&questions).Error; err != nil {
		log.Println("Error fetching quiz questions:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch quiz questions"})
	}
	return c.JSON(questions)
}

// Get quiz question by ID
func (qqc *QuizQuestionController) GetQuizQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Fetching quiz question with ID:", id)
	var question models.QuizQuestion
	if err := qqc.DB.First(&question, id).Error; err != nil {
		log.Println("Quiz question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz question not found"})
	}
	return c.JSON(question)
}

// Get quiz questions by quiz ID
func (qqc *QuizQuestionController) GetQuizQuestionsByQuizID(c *fiber.Ctx) error {
	quizID := c.Params("quizId")
	log.Printf("[INFO] Fetching questions for quiz_id: %s\n", quizID)

	var questions []models.QuizQuestion
	if err := qqc.DB.Where("quizzes_id = ?", quizID).Find(&questions).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch questions for quiz_id %s: %v\n", quizID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch questions"})
	}

	log.Printf("[SUCCESS] Retrieved %d questions for quiz_id %s\n", len(questions), quizID)
	return c.JSON(questions)
}



// Create a new quiz question
func (qqc *QuizQuestionController) CreateQuizQuestion(c *fiber.Ctx) error {
	log.Println("Creating a new quiz question")
	var question models.QuizQuestion
	if err := c.BodyParser(&question); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := qqc.DB.Create(&question).Error; err != nil {
		log.Println("Error creating quiz question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create quiz question"})
	}
	return c.Status(fiber.StatusCreated).JSON(question)
}

// Update an existing quiz question
func (qqc *QuizQuestionController) UpdateQuizQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Updating quiz question with ID:", id)

	var question models.QuizQuestion
	if err := qqc.DB.First(&question, id).Error; err != nil {
		log.Println("Quiz question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Quiz question not found"})
	}

	var requestData map[string]interface{}
	if err := c.BodyParser(&requestData); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := qqc.DB.Model(&question).Updates(requestData).Error; err != nil {
		log.Println("Error updating quiz question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update quiz question"})
	}

	return c.JSON(question)
}

// Delete a quiz question
func (qqc *QuizQuestionController) DeleteQuizQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Deleting quiz question with ID:", id)
	if err := qqc.DB.Delete(&models.QuizQuestion{}, id).Error; err != nil {
		log.Println("Error deleting quiz question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete quiz question"})
	}
	return c.JSON(fiber.Map{"message": "Quiz question deleted successfully"})
}
