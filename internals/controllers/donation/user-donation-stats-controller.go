package controllers

import (
	"log"
	"time"

	"arabiya-syari-fiber/internals/models/donation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UsersDonationStatsController struct {
	DB *gorm.DB
}

func NewUsersDonationStatsController(db *gorm.DB) *UsersDonationStatsController {
	return &UsersDonationStatsController{DB: db}
}

// Get all donation stats
func (udsc *UsersDonationStatsController) GetAll(c *fiber.Ctx) error {
	log.Println("Fetching all user donation stats")
	var stats []models.UsersDonationStat
	if err := udsc.DB.Find(&stats).Error; err != nil {
		log.Println("Error fetching donation stats:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch donation stats"})
	}
	return c.JSON(stats)
}

// Get a specific donation stat by ID
func (udsc *UsersDonationStatsController) GetByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Fetching donation stat with ID:", id)
	var stat models.UsersDonationStat
	if err := udsc.DB.First(&stat, id).Error; err != nil {
		log.Println("Donation stat not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Donation stat not found"})
	}
	return c.JSON(stat)
}

// Create a new donation stat
func (udsc *UsersDonationStatsController) Create(c *fiber.Ctx) error {
	log.Println("Creating a new user donation stat")
	var stat models.UsersDonationStat

	if err := c.BodyParser(&stat); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := udsc.DB.Create(&stat).Error; err != nil {
		log.Println("Error creating donation stat:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create donation stat"})
	}
	return c.Status(fiber.StatusCreated).JSON(stat)
}

// Update a donation stat
func (udsc *UsersDonationStatsController) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Updating donation stat with ID:", id)
	var stat models.UsersDonationStat
	if err := udsc.DB.First(&stat, id).Error; err != nil {
		log.Println("Donation stat not found:", err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Donation stat not found"})
	}

	var input models.UsersDonationStat
	if err := c.BodyParser(&input); err != nil {
		log.Println("Invalid request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := udsc.DB.Model(&stat).Updates(input).Error; err != nil {
		log.Println("Error updating donation stat:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update donation stat"})
	}
	return c.JSON(stat)
}

// Soft delete a donation stat
func (udsc *UsersDonationStatsController) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Println("Invalid ID format:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	log.Println("Deleting donation stat with ID:", id)
	if err := udsc.DB.Model(&models.UsersDonationStat{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		log.Println("Error deleting donation stat:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete donation stat"})
	}
	return c.JSON(fiber.Map{"message": "Donation stat deleted successfully"})
}
