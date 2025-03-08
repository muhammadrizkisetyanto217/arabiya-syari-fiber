 package quizzes

import (
	"log"
	models "arabiya-syari-fiber/internals/models/quizzes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type EvaluationController struct {
	DB *gorm.DB
}

// ✅ Inisialisasi controller
func NewEvaluationController(db *gorm.DB) *EvaluationController {
	return &EvaluationController{DB: db}
}

// ✅ Get all evaluations
func (ec *EvaluationController) GetEvaluations(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all evaluations")
	var evaluations []models.EvaluationModel
	if err := ec.DB.Find(&evaluations).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch evaluations: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch evaluations"})
	}
	log.Printf("[SUCCESS] Retrieved %d evaluations\n", len(evaluations))
	return c.JSON(evaluations)
}

// ✅ Get evaluation by ID
func (ec *EvaluationController) GetEvaluation(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Fetching evaluation with ID: %s\n", id)

	var evaluation models.EvaluationModel
	if err := ec.DB.First(&evaluation, id).Error; err != nil {
		log.Printf("[ERROR] Evaluation with ID %s not found\n", id)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Evaluation not found"})
	}
	log.Printf("[SUCCESS] Retrieved evaluation: ID=%s, Name=%s\n", id, evaluation.NameEvaluation)
	return c.JSON(evaluation)
}

// ✅ Get evaluations by Unit ID
func (ec *EvaluationController) GetEvaluationsByUnitID(c *fiber.Ctx) error {
	unitID := c.Params("unitId")
	log.Printf("[INFO] Fetching evaluations with unit ID: %s\n", unitID)

	var evaluations []models.EvaluationModel
	if err := ec.DB.Where("unit_id = ?", unitID).Find(&evaluations).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch evaluations for unit ID %s: %v\n", unitID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch evaluations"})
	}
	log.Printf("[SUCCESS] Retrieved %d evaluations for unit ID %s\n", len(evaluations), unitID)
	return c.JSON(evaluations)
}

// ✅ Create new evaluation
func (ec *EvaluationController) CreateEvaluation(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new evaluation")
	var evaluation models.EvaluationModel

	if err := c.BodyParser(&evaluation); err != nil {
		log.Printf("[ERROR] Invalid input: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := ec.DB.Create(&evaluation).Error; err != nil {
		log.Printf("[ERROR] Failed to create evaluation: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create evaluation"})
	}
	log.Printf("[SUCCESS] Evaluation created: ID=%d, Name=%s\n", evaluation.ID, evaluation.NameEvaluation)
	return c.Status(fiber.StatusCreated).JSON(evaluation)
}

// ✅ Update evaluation
func (ec *EvaluationController) UpdateEvaluation(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Updating evaluation with ID: %s\n", id)

	var evaluation models.EvaluationModel
	if err := ec.DB.First(&evaluation, id).Error; err != nil {
		log.Printf("[ERROR] Evaluation with ID %s not found\n", id)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Evaluation not found"})
	}

	if err := c.BodyParser(&evaluation); err != nil {
		log.Printf("[ERROR] Invalid input: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := ec.DB.Save(&evaluation).Error; err != nil {
		log.Printf("[ERROR] Failed to update evaluation: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update evaluation"})
	}
	log.Printf("[SUCCESS] Evaluation updated: ID=%s, Name=%s\n", id, evaluation.NameEvaluation)
	return c.JSON(evaluation)
}

// ✅ Delete evaluation
func (ec *EvaluationController) DeleteEvaluation(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Printf("[INFO] Deleting evaluation with ID: %s\n", id)

	if err := ec.DB.Delete(&models.EvaluationModel{}, id).Error; err != nil {
		log.Printf("[ERROR] Failed to delete evaluation: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete evaluation"})
	}
	log.Printf("[SUCCESS] Evaluation with ID %s deleted successfully\n", id)
	return c.JSON(fiber.Map{"message": "Evaluation deleted successfully"})
}
