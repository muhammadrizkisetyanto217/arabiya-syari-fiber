package report_user

import (
	"arabiya-syari-fiber/internals/models/quizzes"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Fungsi untuk memperbarui UserSectionQuizzesModel
// UserSectionQuizzesModel menyimpan daftar kuis yang telah diselesaikan dalam suatu section
type UserSectionQuizzesModel struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       uint          `gorm:"not null" json:"user_id"`
	SectionID    uint          `gorm:"column:section_quizzes_id;not null" json:"section_id"`
	CompleteQuiz pq.Int64Array `gorm:"type:integer[]" json:"complete_quiz"`
	CreatedAt    time.Time     `gorm:"default:current_timestamp" json:"created_at"`
}




func (UserSectionQuizzesModel) TableName() string {
	return "user_section_quizzes"
}

// Fungsi untuk memperbarui UserSectionQuizzesModel
func UpdateUserSectionQuizzes(db *gorm.DB, userID uint, quizID uint) error {
	var sectionID uint

	// Menggunakan `QuizModel` dari package quizzes
	err := db.Model(&quizzes.QuizModel{}).
		Select("section_quizzes_id").
		Where("id = ?", quizID).
		First(&sectionID).Error
	if err != nil {
		return err
	}

	// Ambil semua quiz_id yang telah dikerjakan oleh user dalam section ini
	var completedQuizzes []uint
	err = db.Model(&UserQuizzesModel{}).
		Select("quiz_id").
		Where("user_id = ? AND quiz_id IN (?)",
			userID,
			db.Model(&quizzes.QuizModel{}).Select("id").Where("section_quizzes_id = ?", sectionID),
		).
		Find(&completedQuizzes).Error
	if err != nil {
		return err
	}

	// Konversi []uint ke []int64 agar kompatibel dengan pq.Int64Array
	completedQuizzesInt64 := pq.Int64Array(ConvertUintToInt64(completedQuizzes))

	// Cek apakah data sudah ada di user_section_quizzes
	var userSectionQuiz UserSectionQuizzesModel
	result := db.Where("user_id = ? AND section_quizzes_id = ?", userID, sectionID).First(&userSectionQuiz)

	if result.Error != nil && result.Error == gorm.ErrRecordNotFound {
		// Jika belum ada, buat baru
		userSectionQuiz = UserSectionQuizzesModel{
			UserID:       userID,
			SectionID:    sectionID,
			CompleteQuiz: completedQuizzesInt64,
		}
		return db.Create(&userSectionQuiz).Error
	}

	if result.Error != nil {
		return result.Error
	}

	// Jika sudah ada, update `complete_quiz`
	err = db.Model(&userSectionQuiz).Updates(map[string]interface{}{
		"complete_quiz": completedQuizzesInt64,
	}).Error

	return err
}

// ConvertUintToInt64 mengubah []uint menjadi []int64 agar kompatibel dengan pq.Int64Array
func ConvertUintToInt64(input []uint) []int64 {
	output := make([]int64, len(input))
	for i, v := range input {
		output[i] = int64(v)
	}
	return output
}
