package models

import (
	"time"
	"gorm.io/gorm"
)

type QuizQuestion struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	QuestionText   string         `gorm:"type:text;not null" json:"question_text"`
	Status        string         `gorm:"type:varchar(10);not null;default:'pending';check:status IN ('active', 'pending', 'archived')" json:"status"`
	ParagraphHelp  string         `gorm:"type:text;not null" json:"paragraph_help"`
	ExplainQuestion string        `gorm:"type:text;not null" json:"explain_question"`
	AnswerText     string         `gorm:"type:text;not null" json:"answer_text"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	QuizzesID      uint           `gorm:"not null;index" json:"quizzes_id"`
}

// TableName mengatur nama tabel agar sesuai dengan skema database
func (QuizQuestion) TableName() string {
	return "quizzes_questions"
}