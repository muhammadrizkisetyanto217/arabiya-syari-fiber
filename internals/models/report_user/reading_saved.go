package report_user

import (
	"time"
)

type ReadingSavedModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`    // User yang menyimpan reading
	ReadingID uint           `gorm:"not null" json:"reading_id"` // Reading yang disimpan
	CreatedAt time.Time      `gorm:"default:current_timestamp" json:"created_at"`
}

// TableName mengembalikan nama tabel di database
func (ReadingSavedModel) TableName() string {
	return "reading_saved"
}
