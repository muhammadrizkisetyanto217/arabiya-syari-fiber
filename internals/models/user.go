package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"size:50;not null" json:"name"`
	Email     string     `gorm:"size:255;unique;not null" json:"email"`
	Password  string     `gorm:"not null" json:"-"`
	GoogleID  *string    `gorm:"size:255;unique" json:"google_id,omitempty"`
	Role      string     `gorm:"type:varchar(20);not null;default:'user'" json:"role"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

// BeforeCreate hook untuk memastikan CreatedAt diatur saat record dibuat
func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate hook untuk memastikan UpdatedAt diatur saat record diupdate
// func (u *User) BeforeUpdate(tx *gorm.DB) error {
// 	u.UpdatedAt = time.Now()
// 	return nil
// }