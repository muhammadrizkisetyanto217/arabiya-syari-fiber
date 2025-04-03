package progress_user

import (
    "time"
)

type RankLevelRequirement struct {
    ID          uint      `gorm:"primaryKey" json:"id"`
    NameRank    string    `gorm:"size:100;not null" json:"name_rank"`
    MaxLevel    int       `gorm:"not null" json:"max_level"`
    IconURL     string    `gorm:"size:255;not null" json:"icon_url"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (RankLevelRequirement) TableName() string {
    return "rank_level_requirement"
}
