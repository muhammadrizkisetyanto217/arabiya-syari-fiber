package auth

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Middleware untuk melindungi route menggunakan JWT Authentication
func AuthMiddleware(c *fiber.Ctx) error {
	// 🔍 Ambil token dari header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		// 🚫 Jika header kosong, user tidak diizinkan
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - No token provided"})
	}

	// 🛠️ Pisahkan "Bearer" dan tokennya
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		// 🚫 Jika format salah, tolak permintaan
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - Invalid token format"})
	}

	// 🎫 Ambil token dari string
	tokenString := tokenParts[1]

	// 🛡️ Validasi token
	token, err := ValidateToken(tokenString)
	if err != nil || !token.Valid {
		// 🚫 Jika token tidak valid, tolak permintaan
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - Invalid token"})
	}

	// 📌 Ambil claims (data di dalam token JWT)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// 🚫 Jika gagal mendapatkan claims, tolak permintaan
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - Invalid token claims"})
	}

	// 🆔 Ambil user ID dari token
	userID, ok := claims["id"].(float64)
	if !ok {
		// 🚫 Jika tidak bisa membaca ID, tolak permintaan
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized - User ID not found in token"})
	}

	// 📌 Simpan user ID di context agar bisa diakses di handler berikutnya
	c.Locals("user_id", uint(userID))

	// ✅ Lanjut ke handler berikutnya
	return c.Next()
}
