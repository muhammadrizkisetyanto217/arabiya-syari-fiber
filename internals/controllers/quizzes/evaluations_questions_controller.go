package quizzes

import (
	"log"

	"arabiya-syari-fiber/internals/models/quizzes"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type EvaluationsQuestionController struct {
	DB *gorm.DB
}

func NewEvaluationsQuestionController(db *gorm.DB) *EvaluationsQuestionController {
	return &EvaluationsQuestionController{DB: db}
}

// GET all evaluation questions
func (eqc *EvaluationsQuestionController) GetEvaluationsQuestions(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all evaluation questions")
	var questions []quizzes.EvaluationsQuestionModel

	if err := eqc.DB.Find(&questions).Error; err != nil {
		log.Println("[ERROR] Failed to fetch evaluation questions:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch evaluation questions"})
	}

	log.Printf("[SUCCESS] Retrieved %d evaluation questions\n", len(questions))
	return c.JSON(fiber.Map{
		"message": "Evaluation questions fetched successfully",
		"total":   len(questions),
		"data":    questions,
	})
}

// GET question by ID
func (eqc *EvaluationsQuestionController) GetEvaluationQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Fetching evaluation question with ID:", id)

	var question quizzes.EvaluationsQuestionModel
	if err := eqc.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Evaluation question not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Evaluation question not found"})
	}

	log.Printf("[SUCCESS] Evaluation question fetched: ID=%s\n", id)
	return c.JSON(fiber.Map{
		"message": "Evaluation question fetched successfully",
		"data":    question,
	})
}

// GET questions by evaluation ID
func (eqc *EvaluationsQuestionController) GetQuestionsByEvaluationID(c *fiber.Ctx) error {
	evaluationID := c.Params("evaluationId")
	log.Printf("[INFO] Fetching questions for evaluation_id: %s\n", evaluationID)

	var questions []quizzes.EvaluationsQuestionModel
	if err := eqc.DB.Where("evaluation_id = ?", evaluationID).Find(&questions).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch questions for evaluation_id %s: %v\n", evaluationID, err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch questions"})
	}

	log.Printf("[SUCCESS] Retrieved %d questions for evaluation_id %s\n", len(questions), evaluationID)
	return c.JSON(fiber.Map{
		"message": "Evaluation questions fetched successfully by evaluation ID",
		"total":   len(questions),
		"data":    questions,
	})
}

// POST create new question
func (eqc *EvaluationsQuestionController) CreateEvaluationQuestion(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new evaluation question")

	var question quizzes.EvaluationsQuestionModel
	if err := c.BodyParser(&question); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := eqc.DB.Create(&question).Error; err != nil {
		log.Println("[ERROR] Failed to create evaluation question:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create evaluation question"})
	}

	log.Printf("[SUCCESS] Evaluation question created: ID=%d\n", question.ID)
	return c.Status(201).JSON(fiber.Map{
		"message": "Evaluation question created successfully",
		"data":    question,
	})
}

// PUT update question
func (eqc *EvaluationsQuestionController) UpdateEvaluationQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Updating evaluation question with ID:", id)

	var question quizzes.EvaluationsQuestionModel
	if err := eqc.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Evaluation question not found:", err)
		return c.Status(404).JSON(fiber.Map{"error": "Evaluation question not found"})
	}

	if err := c.BodyParser(&question); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := eqc.DB.Save(&question).Error; err != nil {
		log.Println("[ERROR] Failed to update evaluation question:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update evaluation question"})
	}

	log.Printf("[SUCCESS] Evaluation question updated: ID=%s\n", id)
	return c.JSON(fiber.Map{
		"message": "Evaluation question updated successfully",
		"data":    question,
	})
}

// DELETE question
func (eqc *EvaluationsQuestionController) DeleteEvaluationQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Deleting evaluation question with ID:", id)

	if err := eqc.DB.Delete(&quizzes.EvaluationsQuestionModel{}, id).Error; err != nil {
		log.Println("[ERROR] Failed to delete evaluation question:", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete evaluation question"})
	}

	log.Printf("[SUCCESS] Evaluation question with ID %s deleted\n", id)
	return c.JSON(fiber.Map{
		"message": "Evaluation question deleted successfully",
	})
}
