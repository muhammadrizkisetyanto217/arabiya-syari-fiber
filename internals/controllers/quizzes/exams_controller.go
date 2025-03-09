package quizzes

import (
	"log"

	"arabiya-syari-fiber/internals/models/quizzes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ExamController struct {
	DB *gorm.DB
}

func NewExamController(db *gorm.DB) *ExamController {
	return &ExamController{DB: db}
}

// ✅ Get all exams
func (ec *ExamController) GetExams(c *fiber.Ctx) error {
	log.Println("[INFO] Fetching all exams")
	var exams []quizzes.ExamModel
	if err := ec.DB.Find(&exams).Error; err != nil {
		log.Println("[ERROR] Failed to fetch exams:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch exams"})
	}
	return c.JSON(exams)
}

// ✅ Get exam by ID
func (ec *ExamController) GetExam(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Fetching exam with ID:", id)

	var exam quizzes.ExamModel
	if err := ec.DB.First(&exam, id).Error; err != nil {
		log.Println("[ERROR] Exam not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Exam not found"})
	}
	return c.JSON(exam)
}

// ✅ Get exams by unit_id
func (ec *ExamController) GetExamsByUnitID(c *fiber.Ctx) error {
	unitID := c.Params("unitId")
	log.Printf("[INFO] Fetching exams for unit_id: %s\n", unitID)

	var exams []quizzes.ExamModel
	if err := ec.DB.Where("unit_id = ?", unitID).Find(&exams).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch exams for unit_id %s: %v\n", unitID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch exams"})
	}

	log.Printf("[SUCCESS] Retrieved %d exams for unit_id %s\n", len(exams), unitID)
	return c.JSON(exams)
}


// ✅ Create a new exam
func (ec *ExamController) CreateExam(c *fiber.Ctx) error {
	log.Println("[INFO] Creating a new exam")
	var exam quizzes.ExamModel

	// Parsing JSON ke model
	if err := c.BodyParser(&exam); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validasi manual sebelum menyimpan
	if err := exam.Validate(); err != nil {
		log.Println("[ERROR] Validation failed:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed", "details": err.Error()})
	}

	if err := ec.DB.Create(&exam).Error; err != nil {
		log.Println("[ERROR] Failed to create exam:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create exam"})
	}

	return c.Status(fiber.StatusCreated).JSON(exam)
}

// ✅ Update an existing exam
func (ec *ExamController) UpdateExam(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Updating exam with ID:", id)

	var exam quizzes.ExamModel
	if err := ec.DB.First(&exam, id).Error; err != nil {
		log.Println("[ERROR] Exam not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Exam not found"})
	}

	var requestData quizzes.ExamModel
	if err := c.BodyParser(&requestData); err != nil {
		log.Println("[ERROR] Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Validasi sebelum update
	if err := requestData.Validate(); err != nil {
		log.Println("[ERROR] Validation failed:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed", "details": err.Error()})
	}

	if err := ec.DB.Model(&exam).Updates(requestData).Error; err != nil {
		log.Println("[ERROR] Failed to update exam:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update exam"})
	}

	return c.JSON(exam)
}

// ✅ Delete an exam
func (ec *ExamController) DeleteExam(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("[INFO] Deleting exam with ID:", id)

	if err := ec.DB.Delete(&quizzes.ExamModel{}, id).Error; err != nil {
		log.Println("[ERROR] Failed to delete exam:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete exam"})
	}

	return c.JSON(fiber.Map{"message": "Exam deleted successfully"})
}
