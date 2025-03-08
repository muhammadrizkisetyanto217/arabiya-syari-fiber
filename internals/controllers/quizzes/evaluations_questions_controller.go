package quizzes

import (
	"log"

	models "arabiya-syari-fiber/internals/models/quizzes"
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

// ✅ Get all evaluations questions
func (eqc *EvaluationsQuestionController) GetEvaluationsQuestions(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all evaluations questions")
	var questions []models.EvaluationsQuestionModel
	if err := eqc.DB.Find(&questions).Error; err != nil {
		log.Println("[ERROR] Failed to fetch evaluations questions:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch evaluations questions"})
	}
	return c.JSON(questions)
}

// ✅ Get a single evaluation question by ID
func (eqc *EvaluationsQuestionController) GetEvaluationQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Fetching evaluation question with ID:", id)

	var question models.EvaluationsQuestionModel
	if err := eqc.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Evaluation question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Evaluation question not found"})
	}
	return c.JSON(question)
}

// ✅ Get evaluations questions by evaluation ID
func (eqc *EvaluationsQuestionController) GetQuestionsByEvaluationID(c *fiber.Ctx) error {
	evaluationID := c.Params("evaluationId")
	log.Printf("[INFO] Fetching questions for evaluation_id: %s\n", evaluationID)

	var questions []models.EvaluationsQuestionModel
	if err := eqc.DB.Where("evaluation_id = ?", evaluationID).Find(&questions).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch questions for evaluation_id %s: %v\n", evaluationID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch questions"})
	}

	log.Printf("[SUCCESS] Retrieved %d questions for evaluation_id %s\n", len(questions), evaluationID)
	return c.JSON(questions)
}

// ✅ Create a new evaluation question
func (eqc *EvaluationsQuestionController) CreateEvaluationQuestion(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new evaluation question")
	var question models.EvaluationsQuestionModel

	// Parsing JSON ke struct
	if err := c.BodyParser(&question); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Konversi array string ke pq.StringArray
	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := eqc.DB.Create(&question).Error; err != nil {
		log.Println("[ERROR] Failed to create evaluation question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create evaluation question"})
	}

	return c.Status(fiber.StatusCreated).JSON(question)
}

// ✅ Update an existing evaluation question
func (eqc *EvaluationsQuestionController) UpdateEvaluationQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Updating evaluation question with ID:", id)

	var question models.EvaluationsQuestionModel
	if err := eqc.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Evaluation question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Evaluation question not found"})
	}

	// Parsing JSON
	if err := c.BodyParser(&question); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Konversi array string ke pq.StringArray
	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := eqc.DB.Save(&question).Error; err != nil {
		log.Println("[ERROR] Failed to update evaluation question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update evaluation question"})
	}

	return c.JSON(question)
}

// ✅ Delete an evaluation question
func (eqc *EvaluationsQuestionController) DeleteEvaluationQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Deleting evaluation question with ID:", id)

	if err := eqc.DB.Delete(&models.EvaluationsQuestionModel{}, id).Error; err != nil {
		log.Println("[ERROR] Failed to delete evaluation question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete evaluation question"})
	}

	return c.JSON(fiber.Map{"message": "Evaluation question deleted successfully"})
}
