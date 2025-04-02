package progress_user

import "time"

type UserPointLog struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	UserID     uint   `gorm:"not null" json:"user_id"`
	Points     int    `gorm:"not null" json:"points"`
	SourceType string `gorm:"size:50" json:"source_type"` // "quiz", "exam", dst.
	SourceID   uint   `gorm:"not null" json:"source_id"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func (UserPointLog) TableName() string {
	return "user_point_logs"
}
