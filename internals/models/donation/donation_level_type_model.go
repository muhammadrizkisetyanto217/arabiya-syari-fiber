package donation

import (
	"time"

	"gorm.io/gorm"
)

type DonationLevelsTypeModel struct {
    ID           uint           `gorm:"primaryKey" json:"id"`
    Name         string         `gorm:"size:50;unique;not null" json:"name"`
    Description  string         `gorm:"size:250;not null" json:"description"`
    MinAmount    int            `gorm:"not null" json:"min_amount"`
    PrizeDiamond int            `gorm:"default:10;not null" json:"prize_diamond"`
    IconURL      string         `gorm:"size:255" json:"icon_url"`
    CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (DonationLevelsTypeModel) TableName() string {
	return "donation_levels_type"
}