package models

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type UsersProfile struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	UserID       uint           `gorm:"not null;index" json:"user_id"`
	DonationName string         `gorm:"size:50" json:"donation_name"`
	FullName     string         `gorm:"size:50" json:"full_name"`
	DateOfBirth *time.Time `json:"date_of_birth" time_format:"2006-01-02"`
	Gender       Gender         `gorm:"size:10" json:"gender"`
	PhoneNumber  string         `gorm:"size:20" json:"phone_number"`
	Bio          string         `gorm:"size:300" json:"bio"`
	Location     string         `gorm:"size:50" json:"location"`
	Occupation   string         `gorm:"size:20" json:"occupation"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// Pastikan tabel bernama `users_profile`
func (UsersProfile) TableName() string {
	return "users_profile"
}
