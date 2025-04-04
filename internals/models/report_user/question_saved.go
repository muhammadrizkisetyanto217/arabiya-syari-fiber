package report_user

import (
	"time"
)

type QuestionSavedModel struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `gorm:"not null" json:"user_id"`          // ID user yang menyimpan soal
	SourceTypeID int       `gorm:"not null" json:"source_type_id"`   // 1 = quiz, 2 = evaluation, 3 = exam
	QuestionID   uint      `gorm:"not null" json:"question_id"`      // ID soal dari sumber terkait
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"` // Waktu penyimpanan
}

// TableName mengembalikan nama tabel di database
func (QuestionSavedModel) TableName() string {
	return "question_saved"
}
