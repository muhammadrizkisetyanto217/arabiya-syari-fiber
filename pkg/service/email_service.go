package services

import (
	"fmt"
	"net/smtp"
	"os"
)

// Kirim OTP ke email
func SendOTP(email, otp string) error {
	from := os.Getenv("SMTP_EMAIL") // Ambil dari .env
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Konfigurasi SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Format pesan email
	message := []byte(fmt.Sprintf("Subject: Reset Password OTP\n\nYour OTP code is: %s", otp))

	// Kirim email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, message)
	if err != nil {
		return err
	}
	return nil
}
