package quizzes

import (
	"log"

	"arabiya-syari-fiber/internals/models/quizzes"
	"arabiya-syari-fiber/internals/models/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type QuizzesQuestionController struct {
	DB *gorm.DB
}

func NewQuizzesQuestionController(db *gorm.DB) *QuizzesQuestionController {
	return &QuizzesQuestionController{DB: db}
}

// GET all quiz questions
func (qqc *QuizzesQuestionController) GetQuizzesQuestions(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all quiz questions")

	var questions []quizzes.QuizQuestionModel
	if err := qqc.DB.Find(&questions).Error; err != nil {
		log.Println("[ERROR] Failed to fetch quiz questions:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to fetch quiz questions",
		})
	}

	log.Printf("[SUCCESS] Retrieved %d quiz questions\n", len(questions))
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "All quiz questions fetched successfully",
		"total":   len(questions),
		"data":    questions,
	})
}

// GET single quiz question
func (qqc *QuizzesQuestionController) GetQuizzesQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching quiz question by ID: %s\n", id)

	var question quizzes.QuizQuestionModel
	if err := qqc.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Quiz question not found:", err)
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "Quiz question not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Quiz question fetched successfully by ID",
		"data":    question,
	})
}

// GET quiz questions by quiz_id
func (qqc *QuizzesQuestionController) GetQuizzesQuestionsByQuizID(c *fiber.Ctx) error {
	quizID := c.Params("quizzedId")
	log.Printf("[INFO] Fetching quiz questions for quizzes_id: %s\n", quizID)

	var questions []quizzes.QuizQuestionModel
	if err := qqc.DB.Where("quizzes_id = ?", quizID).Find(&questions).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch quiz questions for quizzes_id %s: %v\n", quizID, err)
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to fetch quiz questions by quiz ID",
		})
	}

	log.Printf("[SUCCESS] Retrieved %d quiz questions for quizzes_id %s\n", len(questions), quizID)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Quiz questions fetched successfully by quiz ID",
		"total":   len(questions),
		"data":    questions,
	})
}

// POST create quiz question
func (qqc *QuizzesQuestionController) CreateQuizzesQuestion(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new quiz question")

	var question quizzes.QuizQuestionModel
	if err := c.BodyParser(&question); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request",
		})
	}

	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := qqc.DB.Create(&question).Error; err != nil {
		log.Println("[ERROR] Failed to create quiz question:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to create quiz question",
		})
	}

	log.Printf("[SUCCESS] Quiz question created with ID: %d\n", question.ID)
	return c.Status(201).JSON(fiber.Map{
		"status":  true,
		"message": "Quiz question created successfully",
		"data":    question,
	})
}

// PUT update quiz question
func (qqc *QuizzesQuestionController) UpdateQuizzesQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Updating quiz question with ID: %s\n", id)

	var question quizzes.QuizQuestionModel
	if err := qqc.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Quiz question not found:", err)
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "Quiz question not found",
		})
	}

	if err := c.BodyParser(&question); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Invalid request",
		})
	}

	question.QuestionAnswer = pq.StringArray(question.QuestionAnswer)

	if err := qqc.DB.Save(&question).Error; err != nil {
		log.Println("[ERROR] Failed to update quiz question:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to update quiz question",
		})
	}

	log.Printf("[SUCCESS] Quiz question with ID %s updated\n", id)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Quiz question updated successfully",
		"data":    question,
	})
}

// DELETE quiz question
func (qqc *QuizzesQuestionController) DeleteQuizzesQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Deleting quiz question with ID: %s\n", id)

	if err := qqc.DB.Delete(&quizzes.QuizQuestionModel{}, id).Error; err != nil {
		log.Println("[ERROR] Failed to delete quiz question:", err)
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Failed to delete quiz question",
		})
	}

	log.Printf("[SUCCESS] Quiz question with ID %s deleted\n", id)
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Quiz question deleted successfully",
	})
}

// GET quiz question + tooltips
func (qqc *QuizzesQuestionController) GetQuizzesQuestionWithTooltips(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching quiz question with tooltips, ID: %s\n", id)

	var question quizzes.QuizQuestionModel
	if err := qqc.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Quiz question not found:", err)
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "Quiz question not found",
		})
	}

	var tooltips []utils.Tooltip
	if len(question.TooltipsID) > 0 {
		if err := qqc.DB.Where("id = ANY(?)", pq.Array(question.TooltipsID)).Find(&tooltips).Error; err != nil {
			log.Println("[ERROR] Failed to fetch tooltips:", err)
			return c.Status(500).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to fetch tooltips",
			})
		}
	}

	log.Printf("[SUCCESS] Retrieved quiz question and tooltips for ID: %s\n", id)
	return c.JSON(fiber.Map{
		"status":        true,
		"message":       "Quiz question and tooltips fetched successfully",
		"quiz_question": question,
		"tooltips":      tooltips,
	})
}

// GET only tooltips by quiz question ID
func (qqc *QuizzesQuestionController) GetOnlyQuizzesQuestionTooltips(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching tooltips only for quiz question ID: %s\n", id)

	var question quizzes.QuizQuestionModel
	if err := qqc.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Quiz question not found:", err)
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "Quiz question not found",
		})
	}

	var tooltips []utils.Tooltip
	if len(question.TooltipsID) > 0 {
		if err := qqc.DB.Where("id = ANY(?)", pq.Array(question.TooltipsID)).Find(&tooltips).Error; err != nil {
			log.Println("[ERROR] Failed to fetch tooltips:", err)
			return c.Status(500).JSON(fiber.Map{
				"status":  false,
				"message": "Failed to fetch tooltips",
			})
		}
	}

	log.Printf("[SUCCESS] Retrieved tooltips only for quiz question ID: %s\n", id)
	return c.JSON(fiber.Map{
		"status":   true,
		"message":  "Tooltips fetched successfully",
		"tooltips": tooltips,
	})
}
