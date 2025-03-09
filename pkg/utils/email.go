package utils

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

// Konfigurasi SMTP Gmail
const (
	SMTPHost     = "smtp.gmail.com"
	SMTPPort     = 587
	SMTPUsername = "muhammadrizkisetyanto217@gmail.com"  // Ganti dengan email Gmail kamu
	SMTPPassword = "umcx rsxt ercw llpe"      // Gunakan App Password dari Google
)

// Kirim email reset password
func SendResetPasswordEmail(toEmail, resetLink string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", SMTPUsername)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "Reset Password Anda")
	m.SetBody("text/html", fmt.Sprintf(`
		<h3>Reset Password</h3>
		<p>Klik link berikut untuk mereset password Anda:</p>
		<a href="%s">%s</a>
		<p>Link ini berlaku selama 15 menit.</p>
	`, resetLink, resetLink))

	d := gomail.NewDialer(SMTPHost, SMTPPort, SMTPUsername, SMTPPassword)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("gagal mengirim email: %v", err)
	}
	return nil
}
