package user

import (
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	config "arabiya-syari-fiber/internals/configs"
	models "arabiya-syari-fiber/internals/models/user"
)

// 🔑 Load Secret Key dari Environment
var SecretKey = config.GetEnv("JWT_SECRET")

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{DB: db}
}

// 🔥 REGISTER USER
func (ac *AuthController) Register(c *fiber.Ctx) error {
	var input models.UserModel

	// 📌 Parsing body JSON
	if err := c.BodyParser(&input); err != nil {
		log.Printf("[ERROR] Failed to parse request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}


	// 📌 Validasi input (Cek di Model)
	if err := input.Validate(); err != nil {
		log.Printf("[ERROR] Validation failed: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// 📌 Hash password sebelum menyimpan ke database
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[ERROR] Failed to hash password: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to secure password"})
	}
	input.Password = string(passwordHash)

	// 📌 Simpan user ke database
	if err := ac.DB.Create(&input).Error; err != nil {
		log.Printf("[ERROR] Failed to save user to database: %v", err)

		// 🔥 Deteksi error unik (email sudah ada)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return c.Status(400).JSON(fiber.Map{"error": "Email already registered"})
		}

		return c.Status(500).JSON(fiber.Map{"error": "Failed to register user"})
	}

	log.Printf("[SUCCESS] User registered: ID=%d, Email=%s", input.ID, input.Email)
	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

// 🔥 LOGIN USER (Email atau Nama)
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var input struct {
		Identifier string `json:"identifier"` // Bisa Email atau Nama
		Password   string `json:"password"`
	}

	// 📌 Parsing request JSON
	if err := c.BodyParser(&input); err != nil {
		log.Printf("[ERROR] Failed to parse request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// 📌 Cek apakah input identifier kosong
	if input.Identifier == "" || input.Password == "" {
		log.Println("[ERROR] Identifier or password cannot be empty")
		return c.Status(400).JSON(fiber.Map{"error": "Email, username, and password are required"})
	}

	// 📌 Cek apakah user ada di database berdasarkan Email atau Nama
	var user models.UserModel
	if err := ac.DB.Where("email = ? OR name = ?", input.Identifier, input.Identifier).First(&user).Error; err != nil {
		log.Printf("[ERROR] User not found: Identifier=%s", input.Identifier)
		return c.Status(401).JSON(fiber.Map{"error": "Invalid email, username, or password"})
	}

	// 📌 Cek apakah password benar
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Printf("[ERROR] Password incorrect for user: %s", user.Email)
		return c.Status(401).JSON(fiber.Map{"error": "Invalid email, username, or password"})
	}

	// 📌 Generate JWT token
	expirationTime := time.Now().Add(time.Hour * 96) // 4 hari
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": expirationTime.Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Printf("[ERROR] Failed to generate token for user: %s", user.Email)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	log.Printf("[SUCCESS] User logged in: ID=%d, Email=%s", user.ID, user.Email)
	return c.JSON(fiber.Map{"token": tokenString})
}


// 🔥 MIDDLEWARE: PROTECT ROUTES
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - No token provided"})
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - Invalid token format"})
	}

	tokenString := tokenParts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - Invalid token"})
	}

	return c.Next()
}
