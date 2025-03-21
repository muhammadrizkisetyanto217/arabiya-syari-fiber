package quizzes

import (
	"time"

	"gorm.io/gorm"
)

// Evaluation struct merepresentasikan tabel evaluations di database
type EvaluationModel struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	NameEvaluation string        `gorm:"size:50;not null" json:"name_evaluation" validate:"required,max=50"`
	Status        string         `gorm:"type:varchar(10);default:'pending';check:status IN ('active', 'pending', 'archived')" json:"status" validate:"required,oneof=active pending archived"`
	Point         int            `gorm:"not null;default:30" json:"point" validate:"gte=0"`
	TotalQuestion *int           `json:"total_question,omitempty" validate:"omitempty,gte=0"`
	IconURL       *string        `gorm:"size:100" json:"icon_url,omitempty" validate:"omitempty,url"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	UnitID        uint           `json:"unit_id" validate:"required"`
	CreatedBy     uint           `json:"created_by" validate:"required"`
}

// TableName mengatur nama tabel agar sesuai dengan skema database
func (EvaluationModel) TableName() string {
	return "evaluations"
}
