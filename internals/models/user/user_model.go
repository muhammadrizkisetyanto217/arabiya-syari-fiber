package models

import (
	"time"
)

type UserModel struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"size:50;not null" json:"name"`
	Email     string     `gorm:"size:255;unique;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	GoogleID  *string    `gorm:"size:255;unique" json:"google_id,omitempty"`
	Role      string     `gorm:"type:varchar(20);not null;default:'user'" json:"role"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName memastikan nama tabel sesuai dengan skema database
func (UserModel) TableName() string {
	return "users"
}