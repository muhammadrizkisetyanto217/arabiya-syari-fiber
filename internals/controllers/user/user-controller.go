package controllers

import (
	"github.com/gofiber/fiber/v2"

	"arabiya-syari-fiber/internals/models/user"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

// Get all users
func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	var users []models.User
	if err := uc.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}
	return c.JSON(users)
}

// Get user by ID
func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := uc.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// Create new user
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := uc.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(201).JSON(user)
}

// Update user by ID
func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := uc.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	uc.DB.Save(&user)
	return c.JSON(user)
}

// Delete user by ID
func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := uc.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}