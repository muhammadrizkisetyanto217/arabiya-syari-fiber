package progress_user

import (
	"time"
)

type UserPointLevelRank struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `gorm:"not null;unique" json:"user_id"`
	AmountTotalQuiz int       `gorm:"default:0" json:"amount_total_quiz"`
	LevelID         uint      `gorm:"not null" json:"level_id"`
	MaxPointLevel   int       `gorm:"not null" json:"max_point_level"`
	IconURLLevel    string    `gorm:"size:255;not null" json:"icon_url_level"`
	RankID          uint      `gorm:"not null" json:"rank_id"`
	MaxLevelRank    int       `gorm:"not null" json:"max_level_rank"`
	IconURLRank     string    `gorm:"size:255;not null" json:"icon_url_rank"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (UserPointLevelRank) TableName() string {
	return "user_point_level_rank"
}
