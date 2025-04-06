package report_user_quiz

import (
	"arabiya-syari-fiber/internals/models/report_user"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserExamController struct {
	DB *gorm.DB
}

func NewUserExamController(db *gorm.DB) *UserExamController {
	return &UserExamController{DB: db}
}

// Fungsi untuk menentukan poin berdasarkan attempt
func getAdditionalPointExam(attempt int) int {
	switch attempt {
	case 1:
		return 20
	case 2:
		return 40
	default:
		return 10
	}
}

// POST /api/user_exams
func (ctrl *UserExamController) CreateOrUpdateUserExam(c *fiber.Ctx) error {
	globalStart := time.Now()
	log.Println("[START] CreateOrUpdateUserExam")

	var input report_user.UserExamModel
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

	var exam report_user.UserExamModel
	err := tx.Where("user_id = ? AND exam_id = ?", input.UserID, input.ExamID).First(&exam).Error

	if err == gorm.ErrRecordNotFound {
		attempt := 1
		addPoint := getAdditionalPointExam(attempt)

		exam = report_user.UserExamModel{
			UserID:          input.UserID,
			ExamID:          input.ExamID,
			Attempt:         attempt,
			Point:           addPoint,
			PercentageGrade: input.PercentageGrade,
			TimeDuration:    input.TimeDuration,
		}

		if err := tx.Create(&exam).Error; err != nil {
			log.Println("[ERROR] Gagal insert user_exam:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user exam"})
		}

		logUserPoint(tx, input.UserID, addPoint, input.ExamID, "exam")
		HandleUserPointProgress(tx, input.UserID, addPoint)

		log.Printf("[DONE] Exam created dalam %.2fs", time.Since(globalStart).Seconds())
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User exam created successfully",
			"attempt": attempt,
			"point":   addPoint,
			"user_id": input.UserID,
			"exam_id": input.ExamID,
		})
	}

	if err != nil {
		log.Println("[ERROR] Query record gagal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process request"})
	}

	exam.Attempt += 1
	addPoint := getAdditionalPointExam(exam.Attempt)
	exam.Point += addPoint

	// Simpan nilai terbesar dari percentage_grade dan time_duration
	if input.PercentageGrade > exam.PercentageGrade {
		exam.PercentageGrade = input.PercentageGrade
		exam.TimeDuration = input.TimeDuration
	}

	if err := tx.Model(&exam).Updates(exam).Error; err != nil {
		log.Println("[ERROR] Update user_exam gagal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user exam"})
	}

	logUserPoint(tx, exam.UserID, addPoint, exam.ExamID, "exam")
	HandleUserPointProgress(tx, exam.UserID, addPoint)

	log.Printf("[DONE] Exam updated dalam %.2fs", time.Since(globalStart).Seconds())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User exam updated successfully",
		"attempt": exam.Attempt,
		"point":   exam.Point,
		"user_id": exam.UserID,
		"exam_id": exam.ExamID,
	})
}

// GET /api/user_exams/:user_id
func (ctrl *UserExamController) GetByUserID(c *fiber.Ctx) error {
	start := time.Now()
	log.Println("[START] GetUserExams")

	userID := c.Params("user_id")
	var exams []report_user.UserExamModel

	if err := ctrl.DB.Where("user_id = ?", userID).Find(&exams).Error; err != nil {
		log.Println("[ERROR] Gagal ambil user_exam:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get exams"})
	}

	log.Printf("[DONE] GetUserExams dalam %.2fms", time.Since(start).Seconds()*1000)
	return c.JSON(exams)
}
