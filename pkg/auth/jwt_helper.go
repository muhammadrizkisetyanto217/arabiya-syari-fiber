package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 🔐 Ambil Secret Key dari environment variable
func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("❌ JWT_SECRET tidak ditemukan di environment! Harap pastikan file .env sudah dikonfigurasi dengan benar.")
	}
	return secret
}

// 🎟️ GenerateToken: Membuat token JWT untuk user yang login
func GenerateToken(userID uint) (string, error) {
	// ⏳ Token berlaku selama 4 hari
	expirationTime := time.Now().Add(96 * time.Hour)

	// 🎭 Claims adalah data yang akan disimpan di dalam token
	claims := jwt.MapClaims{
		"id":  userID,                // ID user yang login
		"exp": expirationTime.Unix(), // Waktu kadaluarsa token (epoch time)
	}

	// 🔐 Membuat token dengan metode HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 🔑 Menandatangani token dengan Secret Key
	tokenString, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 🔍 ValidateToken: Memeriksa apakah token JWT valid atau tidak
func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode enkripsi yang digunakan adalah HS256
		return []byte(getSecretKey()), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
