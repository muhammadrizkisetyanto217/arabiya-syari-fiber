package report_user_quiz

import (
	"log"
	"time"

	"arabiya-syari-fiber/internals/models/report_user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type QuestionSavedController struct {
	DB *gorm.DB
}

func NewQuestionSavedController(db *gorm.DB) *QuestionSavedController {
	return &QuestionSavedController{DB: db}
}

// CREATE /api/question_saved
func (ctrl *QuestionSavedController) Create(c *fiber.Ctx) error {
	start := time.Now()
	log.Println("[START] CreateQuestionSaved")

	var single report_user.QuestionSavedModel
	var multiple []report_user.QuestionSavedModel

	// Coba bind ke slice dulu
	if err := c.BodyParser(&multiple); err == nil && len(multiple) > 0 {
		for _, item := range multiple {
			if item.SourceTypeID < 1 || item.SourceTypeID > 3 {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid source_type_id (must be 1–3)"})
			}
		}
		if err := ctrl.DB.Create(&multiple).Error; err != nil {
			log.Println("[ERROR] Failed to save questions (bulk):", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save questions"})
		}
		log.Printf("[DONE] Created multiple questions in %.2fms", time.Since(start).Seconds()*1000)
		return c.Status(fiber.StatusCreated).JSON(multiple)
	}

	// Kalau bukan slice, coba bind sebagai single
	if err := c.BodyParser(&single); err != nil {
		log.Println("[ERROR] Invalid input format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if single.SourceTypeID < 1 || single.SourceTypeID > 3 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid source_type_id (must be 1–3)"})
	}

	if err := ctrl.DB.Create(&single).Error; err != nil {
		log.Println("[ERROR] Failed to save question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save question"})
	}

	log.Printf("[DONE] Created single question in %.2fms", time.Since(start).Seconds()*1000)
	return c.Status(fiber.StatusCreated).JSON(single)
}

// READ (GET by user_id) /api/question_saved/:user_id
func (ctrl *QuestionSavedController) GetByUserID(c *fiber.Ctx) error {
	start := time.Now()
	log.Println("[START] GetQuestionSavedByUser")

	userID := c.Params("user_id")
	var results []report_user.QuestionSavedModel

	if err := ctrl.DB.Where("user_id = ?", userID).Find(&results).Error; err != nil {
		log.Println("[ERROR] Failed to get questions:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get questions"})
	}

	log.Printf("[DONE] Retrieved %d records in %.2fms", len(results), time.Since(start).Seconds()*1000)
	return c.JSON(results)
}

// READ (GET by ID) /api/question_saved/detail/:id
func (ctrl *QuestionSavedController) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var question report_user.QuestionSavedModel

	if err := ctrl.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Question not found"})
	}

	return c.JSON(question)
}

// UPDATE /api/question_saved/:id
func (ctrl *QuestionSavedController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var question report_user.QuestionSavedModel

	if err := ctrl.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Question not found"})
	}

	var input report_user.QuestionSavedModel
	if err := c.BodyParser(&input); err != nil {
		log.Println("[ERROR] Invalid input:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	question.UserID = input.UserID
	question.SourceTypeID = input.SourceTypeID
	question.QuestionID = input.QuestionID

	if err := ctrl.DB.Save(&question).Error; err != nil {
		log.Println("[ERROR] Failed to update question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update question"})
	}

	return c.JSON(question)
}

// DELETE /api/question_saved/:id
func (ctrl *QuestionSavedController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var question report_user.QuestionSavedModel

	if err := ctrl.DB.First(&question, id).Error; err != nil {
		log.Println("[ERROR] Question not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Question not found"})
	}

	if err := ctrl.DB.Delete(&question).Error; err != nil {
		log.Println("[ERROR] Failed to delete question:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete question"})
	}

	return c.JSON(fiber.Map{"message": "Question deleted successfully"})
}
