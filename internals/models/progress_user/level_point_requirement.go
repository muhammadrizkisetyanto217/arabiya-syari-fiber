package progress_user

import (
    "time"
)

type LevelPointRequirement struct {
    ID            uint      `gorm:"primaryKey" json:"id"`
    NameLevel     string    `gorm:"size:100;not null" json:"name_level"`
    MinPointLevel int       `gorm:"not null" json:"min_point_level"`
    IconURL       string    `gorm:"size:255;not null" json:"icon_url"`
    CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (LevelPointRequirement) TableName() string {
    return "level_point_requirement"
}
