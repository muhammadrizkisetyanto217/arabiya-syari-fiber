package report_user

import (
	"time"

	"gorm.io/gorm"
)

type UserQuizzesModel struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID          uint      `gorm:"not null" json:"user_id"`
	QuizID          uint      `gorm:"column:quiz_id;not null" json:"quiz_id"` // Ubah ke quiz_id sesuai tabel baru
	Attempt         int       `gorm:"default:1;not null" json:"attempt"`
	PercentageGrade float32   `gorm:"default:0;not null" json:"percentage_grade"`
	TimeDuration    int       `gorm:"default:0;not null" json:"time_duration"`
	Point           int       `gorm:"default:0;not null" json:"point"`
	CreatedAt       time.Time `gorm:"default:current_timestamp" json:"created_at"`
}

// TableName memastikan GORM menggunakan tabel "user_quizzes"
func (UserQuizzesModel) TableName() string {
	return "user_quizzes"
}


// Hook GORM untuk memperbarui UserSectionQuizzesModel
func (u *UserQuizzesModel) AfterCreate(tx *gorm.DB) error {
	return UpdateUserSectionQuizzes(tx, u.UserID, u.QuizID)
}

func (u *UserQuizzesModel) AfterDelete(tx *gorm.DB) error {
	return UpdateUserSectionQuizzes(tx, u.UserID, u.QuizID)
}
