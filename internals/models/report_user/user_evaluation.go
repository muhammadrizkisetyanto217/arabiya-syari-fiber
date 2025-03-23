package report_user

import (
	"time"

	"gorm.io/gorm"
)

type UserEvaluationModel struct {
	ID              uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uint           `gorm:"not null" json:"user_id"`
	EvaluationID    uint           `gorm:"not null;column:evaluation_id" json:"evaluation_id"`
	Attempt         int            `gorm:"default:1;not null" json:"attempt"`
	PercentageGrade float32        `gorm:"default:0;not null" json:"percentage_grade"`
	TimeDuration    int            `gorm:"default:0;not null" json:"time_duration"`
	Point           int            `gorm:"default:0;not null" json:"point"`
	CreatedAt       time.Time      `gorm:"default:current_timestamp" json:"created_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (UserEvaluationModel) TableName() string {
	return "user_evaluations"
}

func (u *UserEvaluationModel) AfterCreate(tx *gorm.DB) error {
	return UpdateUserUnitFromEvaluation(tx, u.UserID, u.EvaluationID)
}

func (u *UserEvaluationModel) AfterDelete(tx *gorm.DB) error {
	return CheckAndUnsetEvaluationStatus(tx, u.UserID, u.EvaluationID)
}