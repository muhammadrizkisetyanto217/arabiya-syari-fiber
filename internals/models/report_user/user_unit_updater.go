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
