package progress_user

import "time"

type UserPointLog struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"not null"`
	Points     int       `gorm:"not null"`
	SourceType string    `gorm:"size:50"`   // "quiz", "exam", dst.
	SourceID   uint
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}

func (UserPointLog) TableName() string {
	return "user_point_logs"
}
