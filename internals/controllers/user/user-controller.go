package user

import (
	"github.com/gofiber/fiber/v2"

	"arabiya-syari-fiber/internals/models/user"
	"gorm.io/gorm"
)

//* Kita membuat sebuah struct bernama UserController, yang memiliki satu property bernama DB. (Property adalah variabel yang terdapat dalam sebuah struct).
//& Property DB ini adalah pointer ke objek database (gorm.DB), yang akan digunakan untuk mengakses database.
type UserController struct {
	DB *gorm.DB
}
//^ Bayangkan UserController ini seperti seorang kasir toko.
// 1. Agar bisa bekerja, kasir butuh akses ke database toko (misalnya, daftar barang dan harga).
// 2. Dalam kode ini, DB adalah akses ke database yang diberikan ke kasir (UserController).
// 3. Tanpa DB, kasir tidak bisa mencari barang, menambahkan transaksi, dll.


//*  Fungsi NewUserController adalah "constructor"
// Constructor ini digunakan untuk membuat objek UserController dengan database yang bisa disesuaikan.
func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}
// 1. Saat Anda mempekerjakan kasir baru (UserController), Anda harus memberi mereka akses ke database toko (DB).
// 2. NewUserController(db) adalah cara memberi kasir akses ke database saat mereka mulai bekerja.



// Get all users
func (uc *UserController) GetUsers(c *fiber.Ctx) error {
	var users []user.UserModel
	if err := uc.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to retrieve users"})
	}
	return c.JSON(users)
}

// Get user by ID
func (uc *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user user.UserModel
	if err := uc.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

// Create new user
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var user user.UserModel
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
	var user user.UserModel
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
	if err := uc.DB.Delete(&user.UserModel{}, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}
	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}