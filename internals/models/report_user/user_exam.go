package report_user

import (
	"time"

	"gorm.io/gorm"
)

type UserExamModel struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uint           `gorm:"not null" json:"user_id"`
	ExamID          uint           `gorm:"not null;column:exam_id" json:"exam_id"`
	Attempt         int            `gorm:"default:1;not null" json:"attempt"`
	PercentageGrade float32        `gorm:"default:0;not null" json:"percentage_grade"`
	TimeDuration    int            `gorm:"default:0;not null" json:"time_duration"`
	Point           int            `gorm:"default:0;not null" json:"point"`
	CreatedAt       time.Time      `gorm:"default:current_timestamp" json:"created_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (UserExamModel) TableName() string {
	return "user_exams"
}

func (u *UserExamModel) AfterCreate(tx *gorm.DB) error {
	return UpdateUserUnitFromExam(tx, u.UserID, u.ExamID, float64(u.PercentageGrade))
}

func (u *UserExamModel) AfterDelete(tx *gorm.DB) error {
	return CheckAndUnsetExamStatus(tx, u.UserID, u.ExamID)
}
