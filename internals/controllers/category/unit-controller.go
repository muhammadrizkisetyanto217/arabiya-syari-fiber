package controllers

import (
	"arabiya-syari-fiber/internals/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

type UnitController struct {
	DB *gorm.DB
}

func NewUnitController(db *gorm.DB) *UnitController {
	return &UnitController{DB: db}
}

func (uc *UnitController) GetUnits(c *fiber.Ctx) error {
	log.Println("Fetching all units")
	var units []models.Unit
	if err := uc.DB.Find(&units).Error; err != nil {
		log.Println("Error fetching units:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch units"})
	}
	return c.JSON(units)
}

func (uc *UnitController) GetUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Fetching unit with ID:", id)
	var unit models.Unit
	if err := uc.DB.First(&unit, id).Error; err != nil {
		log.Println("Unit not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unit not found"})
	}
	return c.JSON(unit)
}

func (uc *UnitController) GetUnitByThemesOrLevels(c *fiber.Ctx) error {
	themesOrLevelID := c.Params("themesOrLevelId")
	log.Printf("[INFO] Fetching units with themes_or_level_id: %s\n", themesOrLevelID)

	var units []models.Unit
	if err := uc.DB.Where("themes_or_level_id = ?", themesOrLevelID).Find(&units).Error; err != nil {
		log.Printf("[ERROR] Failed to fetch units for themes_or_level_id %s: %v\n", themesOrLevelID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch units"})
	}
	log.Printf("[SUCCESS] Retrieved %d units for themes_or_level_id %s\n", len(units), themesOrLevelID)
	return c.JSON(units)
}

func (uc *UnitController) CreateUnit(c *fiber.Ctx) error {
	log.Println("Creating a new unit")
	var unit models.Unit
	if err := c.BodyParser(&unit); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := uc.DB.Create(&unit).Error; err != nil {
		log.Println("Error creating unit:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create unit"})
	}
	return c.Status(fiber.StatusCreated).JSON(unit)
}

func (uc *UnitController) UpdateUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Updating unit with ID:", id)
	var unit models.Unit
	if err := uc.DB.First(&unit, id).Error; err != nil {
		log.Println("Unit not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Unit not found"})
	}
	if err := c.BodyParser(&unit); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := uc.DB.Save(&unit).Error; err != nil {
		log.Println("Error updating unit:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update unit"})
	}
	return c.JSON(unit)
}

func (uc *UnitController) DeleteUnit(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Deleting unit with ID:", id)
	if err := uc.DB.Delete(&models.Unit{}, id).Error; err != nil {
		log.Println("Error deleting unit:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete unit"})
	}
	return c.JSON(fiber.Map{"message": "Unit deleted successfully"})
}
