package quizzes

import (
	"time"

	"gorm.io/gorm"
)

type SectionQuizModel struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	NameQuizzes      string         `gorm:"size:50;not null" json:"name_quizzes"`
	Status           string         `gorm:"size:10;default:'pending';check:status IN ('active', 'pending', 'archived')" json:"status"`
	MaterialsQuizzes string         `gorm:"type:text;not null" json:"materials_quizzes"`
	IconURL          string         `gorm:"size:100" json:"icon_url"`
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	CreatedBy        uint           `gorm:"not null;constraint:OnDelete:CASCADE" json:"created_by"`
	UnitID           uint           `gorm:"not null;constraint:OnDelete:CASCADE" json:"unit_id"`


}

func (SectionQuizModel) TableName() string {
	return "section_quizzes"
}
