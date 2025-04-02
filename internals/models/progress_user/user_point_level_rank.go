package progress_user

import (
    "time"
)

type UserPointLevelRank struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    UserID          uint      `gorm:"not null;unique" json:"user_id"`
    AmountTotalQuiz int       `gorm:"default:0" json:"amount_total_quiz"`
    UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (UserPointLevelRank) TableName() string {
    return "user_point_level_rank"
}
