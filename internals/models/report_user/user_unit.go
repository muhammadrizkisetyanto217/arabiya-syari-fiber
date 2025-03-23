package report_user

import (
	"time"

	"gorm.io/gorm"
)

type UserUnitModel struct {
	ID                    uint           `gorm:"primaryKey" json:"id"`
	UserID                uint           `gorm:"not null;index" json:"user_id"`
	UnitID                uint           `gorm:"not null;index" json:"unit_id"`
	IsReading             bool           `gorm:"default:false" json:"is_reading"`
	IsEvaluation          bool           `gorm:"default:false" json:"is_evaluation"`
	IsCompleteSectionQuiz bool           `gorm:"column:is_complete_section_quizzes;default:false" json:"is_complete_section_quizzes"`
	TotalSectionQuizzes   int            `gorm:"default:0" json:"total_section_quizzes"`
	GradeExam             float64        `json:"grade_exam"`
	GradeResult           float64        `json:"grade_result"`
	CreatedAt             time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt             gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

// TableName untuk override nama tabel default
func (UserUnitModel) TableName() string {
	return "user_units"
}
