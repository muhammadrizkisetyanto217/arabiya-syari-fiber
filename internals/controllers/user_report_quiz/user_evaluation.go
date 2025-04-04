// package report_user_quiz

// import (
// 	"github.com/gofiber/fiber/v2"
// 	"gorm.io/gorm"

// 	"arabiya-syari-fiber/internals/models/report_user"
// )

// type UserEvaluationController struct {
// 	DB *gorm.DB
// }

// func NewUserEvaluationController(db *gorm.DB) *UserEvaluationController {
// 	return &UserEvaluationController{DB: db}
// }

// // POST /api/user_evaluations
// func (ctrl *UserEvaluationController) Create(c *fiber.Ctx) error {
// 	var input report_user.UserEvaluationModel
// 	if err := c.BodyParser(&input); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	if err := ctrl.DB.Create(&input).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user evaluation"})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(input)
// }

// // GET /api/user_evaluations/:user_id
// func (ctrl *UserEvaluationController) GetByUserID(c *fiber.Ctx) error {
// 	userID := c.Params("user_id")
// 	var evaluations []report_user.UserEvaluationModel

// 	if err := ctrl.DB.Where("user_id = ?", userID).Find(&evaluations).Error; err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get evaluations"})
// 	}

// 	return c.JSON(evaluations)
// }



package report_user_quiz

import (
	// "arabiya-syari-fiber/internals/models/progress_user"
	"arabiya-syari-fiber/internals/models/report_user"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserEvaluationController struct {
	DB *gorm.DB
}

func NewUserEvaluationController(db *gorm.DB) *UserEvaluationController {
	return &UserEvaluationController{DB: db}
}

// Fungsi untuk menentukan poin berdasarkan attempt
func getAdditionalPointEvaluation(attempt int) int {
	switch attempt {
	case 1:
		return 20
	case 2:
		return 40
	default:
		return 10
	}
}

// POST /api/user_evaluations
func (ctrl *UserEvaluationController) CreateOrUpdateUserEvaluation(c *fiber.Ctx) error {
	globalStart := time.Now()
	log.Println("[START] CreateOrUpdateUserEvaluation")

	var input report_user.UserEvaluationModel
	if err := c.BodyParser(&input); err != nil {
		log.Println("[ERROR] Body parsing failed:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	tx := ctrl.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var evaluation report_user.UserEvaluationModel
	err := tx.Where("user_id = ? AND evaluation_id = ?", input.UserID, input.EvaluationID).First(&evaluation).Error

	if err == gorm.ErrRecordNotFound {
		// First time evaluation
		attempt := 1
		addPoint := getAdditionalPointEvaluation(attempt)

		evaluation = report_user.UserEvaluationModel{
			UserID:       input.UserID,
			EvaluationID: input.EvaluationID,
			Attempt:      attempt,
			Point:        addPoint,
		}

		if err := tx.Create(&evaluation).Error; err != nil {
			log.Println("[ERROR] Gagal insert user_evaluation:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user evaluation"})
		}

		logUserPoint(tx, input.UserID, addPoint, input.EvaluationID, "evaluation")
		HandleUserPointProgress(tx, input.UserID, addPoint)

		log.Printf("[DONE] Evaluation created dalam %.2fs", time.Since(globalStart).Seconds())
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message":        "User evaluation created successfully",
			"attempt":        attempt,
			"point":          addPoint,
			"user_id":        input.UserID,
			"evaluation_id":  input.EvaluationID,
		})
	}

	if err != nil {
		log.Println("[ERROR] Query record gagal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process request"})
	}

	// Sudah pernah → Update data
	evaluation.Attempt += 1
	addPoint := getAdditionalPointEvaluation(evaluation.Attempt)
	evaluation.Point += addPoint

	if err := tx.Model(&evaluation).Updates(evaluation).Error; err != nil {
		log.Println("[ERROR] Update user_evaluation gagal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user evaluation"})
	}

	logUserPoint(tx, evaluation.UserID, addPoint, evaluation.EvaluationID, "evaluation")
	HandleUserPointProgress(tx, evaluation.UserID, addPoint)

	log.Printf("[DONE] Evaluation updated dalam %.2fs", time.Since(globalStart).Seconds())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "User evaluation updated successfully",
		"attempt":        evaluation.Attempt,
		"point":          evaluation.Point,
		"user_id":        evaluation.UserID,
		"evaluation_id":  evaluation.EvaluationID,
	})
}

// GET /api/user_evaluations/:user_id
func (ctrl *UserEvaluationController) GetByUserID(c *fiber.Ctx) error {
	start := time.Now()
	log.Println("[START] GetUserEvaluations")

	userID := c.Params("user_id")
	var evaluations []report_user.UserEvaluationModel

	if err := ctrl.DB.Where("user_id = ?", userID).Find(&evaluations).Error; err != nil {
		log.Println("[ERROR] Gagal ambil user_evaluation:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get evaluations"})
	}

	log.Printf("[DONE] GetUserEvaluations dalam %.2fms", time.Since(start).Seconds()*1000)
	return c.JSON(evaluations)
}
