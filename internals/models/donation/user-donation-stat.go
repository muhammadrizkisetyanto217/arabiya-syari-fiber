package models

import (
	"time"
)

type UserDonationLog struct {
	ID             int       `json:"id"`
	Amount         int       `json:"amount"`
	DonatableType  string    `json:"donatable_type" gorm:"default:pending"` // Tambah default
	Prayer         string    `json:"prayer,omitempty"`
	TotalAmount    int       `json:"total_amount"`
	TotalQuestion  int       `json:"total_question"`
	Diamond        int       `json:"diamond" gorm:"default:30"` // Default 30
	DonatableID    string    `json:"donatable_id"`
	DonatedAt      time.Time `json:"donated_at" gorm:"default:CURRENT_TIMESTAMP"`
	UserID         int       `json:"user_id"`
}
