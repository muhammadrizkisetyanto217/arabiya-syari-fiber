package report_user

import (
	"errors"

	"gorm.io/gorm"
)

//////////////////////////////////////////////////////////
// === BAGIAN UNTUK USER READING ===
//////////////////////////////////////////////////////////

func UpdateUserUnitFromReading(db *gorm.DB, userID uint, readingID uint) error {
	var unitID uint

	err := db.Table("readings").
		Select("unit_id").
		Where("id = ?", readingID).
		Scan(&unitID).Error

	if err != nil {
		return err
	}
	if unitID == 0 {
		return errors.New("unit_id not found for reading_id")
	}

	var userUnit UserUnitModel
	result := db.Where("user_id = ? AND unit_id = ?", userID, unitID).First(&userUnit)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		userUnit = UserUnitModel{
			UserID:    userID,
			UnitID:    unitID,
			IsReading: true,
		}
		return db.Create(&userUnit).Error
	} else if result.Error != nil {
		return result.Error
	}

	return db.Model(&userUnit).Update("is_reading", true).Error
}

func CheckAndUnsetUserUnitReadingStatus(db *gorm.DB, userID uint, readingID uint) error {
	var unitID uint
	err := db.Table("readings").
		Select("unit_id").
		Where("id = ?", readingID).
		Scan(&unitID).Error
	if err != nil || unitID == 0 {
		return err
	}

	var count int64
	err = db.Table("user_readings").
		Joins("JOIN readings ON readings.id = user_readings.reading_id").
		Where("user_readings.user_id = ? AND readings.unit_id = ?", userID, unitID).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return db.Model(&UserUnitModel{}).
			Where("user_id = ? AND unit_id = ?", userID, unitID).
			Update("is_reading", false).Error
	}

	return nil
}

//////////////////////////////////////////////////////////
// === BAGIAN UNTUK USER EVALUATION ===
//////////////////////////////////////////////////////////

func UpdateUserUnitFromEvaluation(db *gorm.DB, userID, evaluationID uint) error {
	var unitID uint

	err := db.Table("evaluations").
		Select("unit_id").
		Where("id = ?", evaluationID).
		Scan(&unitID).Error
	if err != nil || unitID == 0 {
		return err
	}

	var userUnit UserUnitModel
	result := db.Where("user_id = ? AND unit_id = ?", userID, unitID).First(&userUnit)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		userUnit = UserUnitModel{
			UserID:       userID,
			UnitID:       unitID,
			IsEvaluation: true,
		}
		return db.Create(&userUnit).Error
	} else if result.Error != nil {
		return result.Error
	}

	return db.Model(&userUnit).Update("is_evaluation", true).Error
}

func CheckAndUnsetEvaluationStatus(db *gorm.DB, userID, evaluationID uint) error {
	var unitID uint
	err := db.Table("evaluations").
		Select("unit_id").
		Where("id = ?", evaluationID).
		Scan(&unitID).Error
	if err != nil || unitID == 0 {
		return err
	}

	var count int64
	err = db.Table("user_evaluations").
		Joins("JOIN evaluations ON evaluations.id = user_evaluations.evaluation_id").
		Where("user_evaluations.user_id = ? AND evaluations.unit_id = ?", userID, unitID).
		Count(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		return db.Model(&UserUnitModel{}).
			Where("user_id = ? AND unit_id = ?", userID, unitID).
			Update("is_evaluation", false).Error
	}

	return nil
}



//////////////////////////////////////////////////////////
// === BAGIAN UNTUK USER QUIZ ===
//////////////////////////////////////////////////////////

func UpdateUserUnitFromExam(db *gorm.DB, userID, examID uint, grade float64) error {
	var unitID uint

	err := db.Table("exams").
		Select("unit_id").
		Where("id = ?", examID).
		Scan(&unitID).Error
	if err != nil || unitID == 0 {
		return err
	}

	var userUnit UserUnitModel
	result := db.Where("user_id = ? AND unit_id = ?", userID, unitID).First(&userUnit)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		userUnit = UserUnitModel{
			UserID:    userID,
			UnitID:    unitID,
			GradeExam: grade,
		}
		return db.Create(&userUnit).Error
	} else if result.Error != nil {
		return result.Error
	}

	// Update GradeExam-nya
	return db.Model(&userUnit).Update("grade_exam", grade).Error
}


func CheckAndUnsetExamStatus(db *gorm.DB, userID, examID uint) error {
	var unitID uint
	err := db.Table("exams").
		Select("unit_id").
		Where("id = ?", examID).
		Scan(&unitID).Error
	if err != nil || unitID == 0 {
		return err
	}

	var count int64
	err = db.Table("user_exams").
		Joins("JOIN exams ON exams.id = user_exams.exam_id").
		Where("user_exams.user_id = ? AND exams.unit_id = ?", userID, unitID).
		Count(&count).Error
	if err != nil {
		return err
	}

	// Jika tidak ada lagi exam → kosongkan GradeExam
	if count == 0 {
		return db.Model(&UserUnitModel{}).
			Where("user_id = ? AND unit_id = ?", userID, unitID).
			Update("grade_exam", 0).Error
	}

	return nil
}
