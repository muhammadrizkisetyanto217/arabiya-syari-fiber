package user

import (
	// "errors"
	"log"
	// "os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	config "arabiya-syari-fiber/internals/configs"
	"arabiya-syari-fiber/internals/models/auth"
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

	if err := c.BodyParser(&input); err != nil {
		log.Printf("[ERROR] Failed to parse request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request format"})
	}

	// Validasi input
	if err := input.Validate(); err != nil {
		log.Printf("[ERROR] Validation failed: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	// Hash password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[ERROR] Failed to hash password: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "Failed to secure password"})
	}
	input.Password = string(passwordHash)

	// Simpan user ke database
	if err := ac.DB.Create(&input).Error; err != nil {
		log.Printf("[ERROR] Failed to save user to database: %v", err)
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return c.Status(400).JSON(fiber.Map{"error": "Email already registered"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Failed to register user"})
	}

	log.Printf("[SUCCESS] User registered: ID=%d, Email=%s", input.ID, input.Email)
	return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
}

// 🔥 LOGIN USER
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var input struct {
		Identifier string `json:"identifier"` // Bisa Email atau Nama
		Password   string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		log.Printf("[ERROR] Failed to parse request body: %v", err)
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Cek user berdasarkan Email atau Nama
	var user models.UserModel
	if err := ac.DB.Where("email = ? OR name = ?", input.Identifier, input.Identifier).First(&user).Error; err != nil {
		log.Printf("[ERROR] User not found: Identifier=%s", input.Identifier)
		return c.Status(401).JSON(fiber.Map{"error": "Invalid email, username, or password"})
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Printf("[ERROR] Password incorrect for user: %s", user.Email)
		return c.Status(401).JSON(fiber.Map{"error": "Invalid email, username, or password"})
	}

	// Generate JWT token
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

// 🔥 LOGOUT USER
func (ac *AuthController) Logout(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - No token provided"})
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - Invalid token format"})
	}
	tokenString := tokenParts[1]

	// Cek apakah token sudah ada di blacklist
	var existingToken auth.TokenBlacklist
	if err := ac.DB.Where("token = ?", tokenString).First(&existingToken).Error; err == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token already blacklisted"})
	}

	// Tambahkan token ke blacklist
	blacklistToken := auth.TokenBlacklist{
		Token:     tokenString,
		ExpiredAt: time.Now().Add(96 * time.Hour), // Sesuai waktu expired token
	}

	if err := ac.DB.Create(&blacklistToken).Error; err != nil {
		log.Printf("[ERROR] Failed to blacklist token: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to logout"})
	}

	return c.JSON(fiber.Map{"message": "Logged out successfully"})
	
}


// 🔥 CHANGE PASSWORD
func (ac *AuthController) ChangePassword(c *fiber.Ctx) error {
	// Ambil token dari request
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - No token provided"})
	}

	// Pisahkan "Bearer" dan tokennya
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - Invalid token format"})
	}
	tokenString := tokenParts[1]

	// Validasi JWT Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - Invalid token"})
	}

	// Ambil User ID dari JWT Claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - Invalid token claims"})
	}

	userID := uint(claims["id"].(float64))

	// Parsing request body
	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}

	// Validasi input
	if input.OldPassword == "" || input.NewPassword == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Both old and new passwords are required"})
	}

	// Cek apakah user ada di database
	var user models.UserModel
	if err := ac.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Cek apakah password lama benar
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Old password is incorrect"})
	}

	// Cek apakah password baru sama dengan yang lama
	if input.OldPassword == input.NewPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "New password must be different from old password"})
	}

	// Hash password baru
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash new password"})
	}

	// Update password di database
	if err := ac.DB.Model(&user).Update("password", string(newHashedPassword)).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update password"})
	}

	return c.JSON(fiber.Map{"message": "Password changed successfully"})
}



// 🔥 Middleware untuk proteksi route
func AuthMiddleware(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - No token provided"})
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - Invalid token format"})
		}

		tokenString := tokenParts[1]

		// Cek blacklist token
		var existingToken auth.TokenBlacklist
		if err := db.Where("token = ?", tokenString).First(&existingToken).Error; err == nil {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - Token is blacklisted"})
		}

		// Validasi JWT token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})

		if err != nil || !token.Valid {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - Invalid token"})
		}

		// Cek Expired
		claims := token.Claims.(jwt.MapClaims)
		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return c.Status(401).JSON(fiber.Map{"error": "Unauthorized - Token expired"})
		}

		return c.Next()
	}
}
