package quizzes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"arabiya-syari-fiber/internals/models/quizzes"
)

type SectionQuizController struct {
	DB *gorm.DB
}

func NewSectionQuizController(db *gorm.DB) *SectionQuizController {
	return &SectionQuizController{DB: db}
}

func (sqc *SectionQuizController) GetSectionQuizzes(c *fiber.Ctx) error {
	log.Println("Fetching all section quizzes")
	var quizzes []models.SectionQuizModel
	if err := sqc.DB.Find(&quizzes).Error; err != nil {
		log.Println("Error fetching section quizzes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch section quizzes"})
	}
	return c.JSON(quizzes)
}

func (sqc *SectionQuizController) GetSectionQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Fetching section quiz with ID:", id)
	var quiz models.SectionQuizModel
	if err := sqc.DB.First(&quiz, id).Error; err != nil {
		log.Println("Section quiz not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Section quiz not found"})
	}
	return c.JSON(quiz)
}

func (sqc *SectionQuizController) GetSectionQuizzesByUnit(c *fiber.Ctx) error {
	unitID := c.Params("unitId")
	log.Printf("[INFO] Fetching section_quizzes for unit_id: %s\n", unitID)

	var sectionQuizzes []models.SectionQuizModel
	if err := sqc.DB.Where("unit_id = ?", unitID).Find(&sectionQuizzes).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch section_quizzes for unit_id %s: %v\n", unitID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch section_quizzes"})
	}

	log.Printf("[SUCCESS] Retrieved %d section_quizzes for unit_id %s\n", len(sectionQuizzes), unitID)
	return c.JSON(sectionQuizzes)
}


func (sqc *SectionQuizController) CreateSectionQuiz(c *fiber.Ctx) error {
	log.Println("Creating a new section quiz")
	var quiz models.SectionQuizModel
	if err := c.BodyParser(&quiz); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := sqc.DB.Create(&quiz).Error; err != nil {
		log.Println("Error creating section quiz:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create section quiz"})
	}
	return c.Status(fiber.StatusCreated).JSON(quiz)
}

func (sqc *SectionQuizController) UpdateSectionQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Updating section quiz with ID:", id)

	var quiz models.SectionQuizModel
	if err := sqc.DB.First(&quiz, id).Error; err != nil {
		log.Println("Section quiz not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Section quiz not found"})
	}

	var requestData map[string]interface{}
	if err := c.BodyParser(&requestData); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := sqc.DB.Model(&quiz).Updates(requestData).Error; err != nil {
		log.Println("Error updating section quiz:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update section quiz"})
	}

	return c.JSON(quiz)
}

func (sqc *SectionQuizController) DeleteSectionQuiz(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Deleting section quiz with ID:", id)
	if err := sqc.DB.Delete(&models.SectionQuizModel{}, id).Error; err != nil {
		log.Println("Error deleting section quiz:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete section quiz"})
	}
	return c.JSON(fiber.Map{"message": "Section quiz deleted successfully"})
}
