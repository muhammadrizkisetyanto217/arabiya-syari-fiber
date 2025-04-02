package progress_user

import (
    "time"
)

type UserPointLevelRank struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    UserID          uint      `gorm:"not null;unique" json:"user_id"`
    AmountTotalQuiz int       `gorm:"default:0" json:"amount_total_quiz"`
    LevelID         uint      `gorm:"not null" json:"level_id"`
    MinPointLevel   int       `gorm:"not null" json:"min_point_level"`
    IconURL         string    `gorm:"size:255;not null" json:"icon_url"`
    CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (UserPointLevelRank) TableName() string {
    return "user_point_level_rank"
}
