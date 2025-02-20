package routes

import (
	"arabiya-syari-fiber/models"
	"arabiya-syari-fiber/database"

	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"
)

// GET: Ambil semua user
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

// GET: Ambil user berdasarkan ID
func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// POST: Buat user baru
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Create(user)
	return c.Status(201).JSON(user)
}

// PUT: Update user berdasarkan ID
func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	database.DB.Save(&user)
	return c.JSON(user)
}

// DELETE: Hapus user berdasarkan ID
func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := database.DB.Delete(&models.User{}, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(fiber.Map{"message": "User deleted"})
}

// Register routes
func SetupUserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")
	userGroup.Get("/", GetUsers)
	userGroup.Get("/:id", GetUser)
	userGroup.Post("/", CreateUser)
	userGroup.Put("/:id", UpdateUser)
	userGroup.Delete("/:id", DeleteUser)
}
