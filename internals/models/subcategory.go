package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Subcategory struct {
	ID                           uint           `json:"id" gorm:"primaryKey"`
	Name                         string         `json:"name"`
	Status                       string         `json:"status" gorm:"type:VARCHAR(10);check:status IN ('active', 'pending', 'archived')"`
	DescriptionLong              string         `json:"description_long"`
	GreatTotalThemesOrLevels     int            `json:"great_total_themes_or_levels"`
	TotalThemesOrLevels          int            `json:"total_themes_or_levels"`
	CompletedTotalThemesOrLevels int            `json:"completed_total_themes_or_levels"`
	UpdateNews                   json.RawMessage `json:"update_news" gorm:"type:jsonb"`
	ImageURL                     string         `json:"image_url"`
	CreatedAt                    time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                    time.Time      `json:"updated_at"`
	DeletedAt                    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	CategoriesID                 int            `json:"category_id"`
}

func (s *Subcategory) BeforeCreate(tx *gorm.DB) error {
	s.CreatedAt = time.Now()
	return nil
}