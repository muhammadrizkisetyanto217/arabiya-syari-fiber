package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Ambil Secret Key dari environment
var secretKey = os.Getenv("JWT_SECRET")

func init() {
	fmt.Println("🔐 JWT_SECRET:", secretKey) // Debugging apakah secret key terload
}

// GenerateToken membuat JWT token baru untuk user
func GenerateToken(userID uint) (string, error) {
	expirationTime := time.Now().Add(96 * time.Hour) // Token berlaku 4 hari
	claims := jwt.MapClaims{
		"id":  userID,
		"exp": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken memeriksa validitas JWT token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
}

// Middleware untuk memeriksa autentikasi JWT
func AuthMiddleware(c *fiber.Ctx) error {
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
	token, err := ValidateToken(tokenString)
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - Invalid token"})
	}

	return c.Next()
}
