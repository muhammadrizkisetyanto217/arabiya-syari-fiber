package progress_user

import "time"

type UserPointLog struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	Points      int       `gorm:"not null"`
	Description string    `gorm:"size:255;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (UserPointLog) TableName() string {
	return "user_point_logs"
}
