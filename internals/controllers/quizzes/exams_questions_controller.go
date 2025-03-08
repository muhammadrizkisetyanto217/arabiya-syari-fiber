package quizzes

import (
	"log"

	models "arabiya-syari-fiber/internals/models/quizzes"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type ExamsQuestionController struct {
	DB *gorm.DB
}

// ✅ **Constructor untuk ExamsQuestionController**
func NewExamsQuestionController(db *gorm.DB) *ExamsQuestionController {
	return &ExamsQuestionController{DB: db}
}

// ✅ **Get All Questions**
func (eqc *ExamsQuestionController) GetExamsQuestions(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all exam questions")
	var questions []models.ExamsQuestionModel
	if err := eqc.DB.Find(&questions).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch questions: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch questions"})
	}
	return c.JSON(questions)
}

// ✅ **Get Question by ID**
func (eqc *ExamsQuestionController) GetExamsQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching question with ID: %s\n", id)

	var question models.ExamsQuestionModel
	if err := eqc.DB.First(&question, id).Error; err != nil {
		log.Printf("[ERROR] Question not found: %v\n", err)
		return c.Status(404).JSON(fiber.Map{"error": "Question not found"})
	}
	return c.JSON(question)
}

// ✅ **Get Questions by Exam ID**
func (eqc *ExamsQuestionController) GetQuestionsByExamID(c *fiber.Ctx) error {
	examID := c.Params("examId")
	log.Printf("[INFO] Fetching questions for exam_id: %s\n", examID)

	var questions []models.ExamsQuestionModel
	if err := eqc.DB.Where("exam_id = ?", examID).Find(&questions).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch questions for exam_id %s: %v\n", examID, err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch questions"})
	}
	log.Printf("[SUCCESS] Retrieved %d questions for exam_id %s\n", len(questions), examID)
	return c.JSON(questions)
}

// ✅ **Create a New Question**
func (eqc *ExamsQuestionController) CreateExamsQuestion(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new exam question")
	var question models.ExamsQuestionModel

	// Parsing request JSON ke struct
	if err := c.BodyParser(&question); err != nil {
		log.Printf("[ERROR] Invalid request body: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Konversi array string ke pq.StringArray
	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	// Simpan ke database
	if err := eqc.DB.Create(&question).Error; err != nil {
		log.Printf("[ERROR] Failed to create question: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create question"})
	}

	log.Printf("[SUCCESS] Created question ID=%d\n", question.ID)
	return c.Status(201).JSON(question)
}

// ✅ **Update an Existing Question**
func (eqc *ExamsQuestionController) UpdateExamsQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Updating question with ID: %s\n", id)

	var question models.ExamsQuestionModel
	if err := eqc.DB.First(&question, id).Error; err != nil {
		log.Printf("[ERROR] Question not found: %v\n", err)
		return c.Status(404).JSON(fiber.Map{"error": "Question not found"})
	}

	var requestData models.ExamsQuestionModel
	if err := c.BodyParser(&requestData); err != nil {
		log.Printf("[ERROR] Invalid request body: %v\n", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Konversi array string ke pq.StringArray
	requestData.QuestionAnswer = pq.StringArray(requestData.QuestionAnswer)

	// Update data
	if err := eqc.DB.Model(&question).Updates(requestData).Error; err != nil {
		log.Printf("[ERROR] Failed to update question: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update question"})
	}

	log.Printf("[SUCCESS] Updated question ID=%s\n", id)
	return c.JSON(question)
}

// ✅ **Delete a Question**
func (eqc *ExamsQuestionController) DeleteExamsQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Deleting question with ID: %s\n", id)

	if err := eqc.DB.Delete(&models.ExamsQuestionModel{}, id).Error; err != nil {
		log.Printf("[ERROR] Failed to delete question: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete question"})
	}

	log.Printf("[SUCCESS] Deleted question ID=%s\n", id)
	return c.JSON(fiber.Map{"message": "Question deleted successfully"})
}
