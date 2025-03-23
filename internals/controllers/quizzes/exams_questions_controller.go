package quizzes

import (
	"log"

	"arabiya-syari-fiber/internals/models/quizzes"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ExamsQuestionController struct {
	DB *gorm.DB
}

// Constructor
func NewExamsQuestionController(db *gorm.DB) *ExamsQuestionController {
	return &ExamsQuestionController{DB: db}
}

// GET all exam questions
func (eqc *ExamsQuestionController) GetExamsQuestions(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all exam questions")

	var questions []quizzes.ExamsQuestionModel
	if err := eqc.DB.Find(&questions).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch questions: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch questions"})
	}

	log.Printf("[SUCCESS] Retrieved %d exam questions\n", len(questions))
	return c.JSON(fiber.Map{
		"message": "Exam questions fetched successfully",
		"total":   len(questions),
		"data":    questions,
	})
}

// GET single question by ID
func (eqc *ExamsQuestionController) GetExamsQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching question with ID: %s\n", id)

	var question quizzes.ExamsQuestionModel
	if err := eqc.DB.First(&question, id).Error; err != nil {
		log.Printf("[ERROR] Question not found: %v\n", err)
		return c.Status(404).JSON(fiber.Map{"error": "Question not found"})
	}

	return c.JSON(fiber.Map{
		"message": "Exam question fetched successfully",
		"data":    question,
	})
}

// GET questions by exam_id
func (eqc *ExamsQuestionController) GetQuestionsByExamID(c *fiber.Ctx) error {
	examID := c.Params("examId")
	log.Printf("[INFO] Fetching questions for exam_id: %s\n", examID)

	var questions []quizzes.ExamsQuestionModel
	if err := eqc.DB.Where("exam_id = ?", examID).Find(&questions).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch questions for exam_id %s: %v\n", examID, err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch questions"})
	}

	log.Printf("[SUCCESS] Retrieved %d questions for exam_id %s\n", len(questions), examID)
	return c.JSON(fiber.Map{
		"message": "Exam questions fetched by exam ID successfully",
		"total":   len(questions),
		"data":    questions,
	})
}

// POST create question
func (eqc *ExamsQuestionController) CreateExamsQuestion(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new exam question")

	var question quizzes.ExamsQuestionModel
	if err := c.BodyParser(&question); err != nil {
		log.Printf("[ERROR] Invalid request body: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := eqc.DB.Create(&question).Error; err != nil {
		log.Printf("[ERROR] Failed to create question: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create question"})
	}

	log.Printf("[SUCCESS] Created question ID=%d\n", question.ID)
	return c.Status(201).JSON(fiber.Map{
		"message": "Exam question created successfully",
		"data":    question,
	})
}

// PUT update question
func (eqc *ExamsQuestionController) UpdateExamsQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Updating question with ID: %s\n", id)

	var question quizzes.ExamsQuestionModel
	if err := eqc.DB.First(&question, id).Error; err != nil {
		log.Printf("[ERROR] Question not found: %v\n", err)
		return c.Status(404).JSON(fiber.Map{"error": "Question not found"})
	}

	var requestData quizzes.ExamsQuestionModel
	if err := c.BodyParser(&requestData); err != nil {
		log.Printf("[ERROR] Invalid request body: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	requestData.QuestionAnswer = pq.StringArray(requestData.QuestionAnswer)

	if err := eqc.DB.Model(&question).Updates(requestData).Error; err != nil {
		log.Printf("[ERROR] Failed to update question: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update question"})
	}

	log.Printf("[SUCCESS] Updated question ID=%s\n", id)
	return c.JSON(fiber.Map{
		"message": "Exam question updated successfully",
		"data":    question,
	})
}

// DELETE question
func (eqc *ExamsQuestionController) DeleteExamsQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Deleting question with ID: %s\n", id)

	if err := eqc.DB.Delete(&quizzes.ExamsQuestionModel{}, id).Error; err != nil {
		log.Printf("[ERROR] Failed to delete question: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete question"})
	}

	log.Printf("[SUCCESS] Deleted question ID=%s\n", id)
	return c.JSON(fiber.Map{
		"message": "Exam question deleted successfully",
	})
}
