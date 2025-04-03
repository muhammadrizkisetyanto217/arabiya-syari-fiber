package report_user_quiz

import (
	"arabiya-syari-fiber/internals/models/progress_user"
	"arabiya-syari-fiber/internals/models/report_user"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserQuizzesController struct {
	DB *gorm.DB
}

func NewUserQuizzesController(db *gorm.DB) *UserQuizzesController {
	return &UserQuizzesController{DB: db}
}

func getAdditionalPoint(attempt int) int {
	switch attempt {
	case 1:
		return 20
	case 2:
		return 40
	default:
		return 10
	}
}

func (uqc *UserQuizzesController) CreateOrUpdateUserQuiz(c *fiber.Ctx) error {
	globalStart := time.Now()
	log.Println("[START] CreateOrUpdateUserQuiz")

	var input report_user.UserQuizzesModel
	if err := c.BodyParser(&input); err != nil {
		log.Println("[ERROR] Body parsing failed:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	tx := uqc.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var userQuiz report_user.UserQuizzesModel
	err := tx.Where("user_id = ? AND quiz_id = ?", input.UserID, input.QuizID).First(&userQuiz).Error

	if err == gorm.ErrRecordNotFound {
		attempt := 1
		addPoint := getAdditionalPoint(attempt)

		userQuiz = report_user.UserQuizzesModel{
			UserID:          input.UserID,
			QuizID:          input.QuizID,
			Attempt:         attempt,
			PercentageGrade: input.PercentageGrade,
			TimeDuration:    input.TimeDuration,
			Point:           addPoint,
		}

		if err := tx.Create(&userQuiz).Error; err != nil {
			log.Println("[ERROR] Gagal insert user_quizzes:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save user quiz result"})
		}

		logUserPoint(tx, input.UserID, addPoint, input.QuizID)
		incrementAmountTotalQuiz(tx, input.UserID, addPoint)

		log.Printf("[DONE] Quiz baru disimpan dalam %.2fs", time.Since(globalStart).Seconds())
		return c.JSON(fiber.Map{
			"message":          "User quiz result saved successfully",
			"attempt":          attempt,
			"percentage_grade": input.PercentageGrade,
			"point":            addPoint,
		})
	}

	if err != nil {
		log.Println("[ERROR] Query record gagal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to process request"})
	}

	userQuiz.Attempt += 1
	addPoint := getAdditionalPoint(userQuiz.Attempt)
	userQuiz.Point += addPoint

	if input.PercentageGrade > userQuiz.PercentageGrade {
		userQuiz.PercentageGrade = input.PercentageGrade
		userQuiz.TimeDuration = input.TimeDuration
	}

	if err := tx.Model(&userQuiz).Updates(userQuiz).Error; err != nil {
		log.Println("[ERROR] Update user_quizzes gagal:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user quiz result"})
	}

	logUserPoint(tx, userQuiz.UserID, addPoint, userQuiz.QuizID)
	incrementAmountTotalQuiz(tx, userQuiz.UserID, addPoint)

	log.Printf("[DONE] Quiz updated dalam %.2fs", time.Since(globalStart).Seconds())
	return c.JSON(fiber.Map{
		"message":          "User quiz result updated successfully",
		"attempt":          userQuiz.Attempt,
		"percentage_grade": userQuiz.PercentageGrade,
		"point":            userQuiz.Point,
	})
}

func (uqc *UserQuizzesController) GetUserQuizzes(c *fiber.Ctx) error {
	start := time.Now()
	log.Println("[START] GetUserQuizzes")

	userID, err := c.ParamsInt("user_id")
	if err != nil {
		log.Println("[ERROR] Invalid user_id:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var userQuizzes []report_user.UserQuizzesModel
	if err := uqc.DB.Where("user_id = ?", userID).Find(&userQuizzes).Error; err != nil {
		log.Println("[ERROR] Gagal ambil data user_quizzes:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user quiz results"})
	}

	log.Printf("[DONE] GetUserQuizzes dalam %.2fms", time.Since(start).Seconds()*1000)
	return c.JSON(userQuizzes)
}

func logUserPoint(tx *gorm.DB, userID uint, points int, quizID uint) {
	err := tx.Create(&progress_user.UserPointLog{
		UserID:     userID,
		Points:     points,
		SourceType: "quiz",
		SourceID:   quizID,
	}).Error

	if err != nil {
		log.Printf("[ERROR] Gagal insert user_point_log: %v", err)
	} else {
		log.Printf("[SUCCESS] Logged user point for user_id=%d, quiz_id=%d, points=%d", userID, quizID, points)
	}
}

// func incrementAmountTotalQuiz(tx *gorm.DB, userID uint, point int) {
// 	var rank progress_user.UserPointLevelRank
// 	if err := tx.Where("user_id = ?", userID).First(&rank).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			rank = progress_user.UserPointLevelRank{
// 				UserID:          userID,
// 				AmountTotalQuiz: point,
// 			}
// 			if err := tx.Create(&rank).Error; err != nil {
// 				log.Println("[ERROR] Gagal create user_point_level_rank:", err)
// 			} else {
// 				log.Println("[INFO] Created new user_point_level_rank")
// 			}
// 		} else {
// 			log.Println("[ERROR] Fetch rank error:", err)
// 		}
// 	} else {
// 		if err := tx.Model(&rank).
// 			Update("amount_total_quiz", gorm.Expr("amount_total_quiz + ?", point)).Error; err != nil {
// 			log.Println("[ERROR] Gagal update total point:", err)
// 		} else {
// 			log.Println("[INFO] Updated existing user_point_level_rank (incremented)")
// 		}
// 	}
// }

// func incrementAmountTotalQuiz(tx *gorm.DB, userID uint, point int) {
// 	var rank progress_user.UserPointLevelRank
// 	if err := tx.Where("user_id = ?", userID).First(&rank).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			// Ambil level default (id=1)
// 			var level progress_user.LevelPointRequirement
// 			if err := tx.First(&level, 1).Error; err != nil {
// 				log.Println("[ERROR] Gagal ambil level ID 1:", err)
// 				return
// 			}

// 			rank = progress_user.UserPointLevelRank{
// 				UserID:          userID,
// 				AmountTotalQuiz: point,
// 				LevelID:         level.ID,
// 				MaxPointLevel:   level.MaxPointLevel,
// 				IconURL:         level.IconURL,
// 			}

// 			if err := tx.Create(&rank).Error; err != nil {
// 				log.Println("[ERROR] Gagal create user_point_level_rank:", err)
// 			} else {
// 				log.Println("[INFO] Created new user_point_level_rank dengan level ID 1")
// 			}
// 		} else {
// 			log.Println("[ERROR] Fetch rank error:", err)
// 		}
// 	} else {
// 		// Sudah ada, cukup update point saja
// 		if err := tx.Model(&rank).
// 			Update("amount_total_quiz", gorm.Expr("amount_total_quiz + ?", point)).Error; err != nil {
// 			log.Println("[ERROR] Gagal update total point:", err)
// 		} else {
// 			log.Println("[INFO] Updated existing user_point_level_rank (incremented)")
// 		}
// 	}
// }

// func incrementAmountTotalQuiz(tx *gorm.DB, userID uint, point int) {
// 	var rank progress_user.UserPointLevelRank
// 	if err := tx.Where("user_id = ?", userID).First(&rank).Error; err != nil {
// 		if err == gorm.ErrRecordNotFound {
// 			// Ambil level default (id=1)
// 			var level progress_user.LevelPointRequirement
// 			if err := tx.First(&level, 1).Error; err != nil {
// 				log.Println("[ERROR] Gagal ambil level ID 1:", err)
// 				return
// 			}

// 			rank = progress_user.UserPointLevelRank{
// 				UserID:          userID,
// 				AmountTotalQuiz: point,
// 				LevelID:         level.ID,
// 				MaxPointLevel:   level.MaxPointLevel,
// 				IconURL:         level.IconURL,
// 			}

// 			if err := tx.Create(&rank).Error; err != nil {
// 				log.Println("[ERROR] Gagal create user_point_level_rank:", err)
// 			} else {
// 				log.Println("[INFO] Created new user_point_level_rank dengan level ID 1")
// 			}
// 		} else {
// 			log.Println("[ERROR] Fetch rank error:", err)
// 		}
// 	} else {
// 		// Update jumlah point
// 		if err := tx.Model(&rank).
// 			Update("amount_total_quiz", gorm.Expr("amount_total_quiz + ?", point)).Error; err != nil {
// 			log.Println("[ERROR] Gagal update total point:", err)
// 			return
// 		}

// 		// Refresh data rank setelah update point
// 		if err := tx.Where("user_id = ?", userID).First(&rank).Error; err != nil {
// 			log.Println("[ERROR] Gagal ambil data rank setelah update:", err)
// 			return
// 		}

// 		// Cek apakah perlu upgrade level
// 		if rank.AmountTotalQuiz > rank.MaxPointLevel {
// 			var nextLevel progress_user.LevelPointRequirement
// 			if err := tx.
// 				Where("max_point_level > ?", rank.MaxPointLevel).
// 				Order("max_point_level ASC").
// 				First(&nextLevel).Error; err != nil {
// 				log.Println("[INFO] Tidak ada level lebih tinggi, tetap di level sekarang")
// 				return
// 			}

// 			// Update ke level berikutnya
// 			if err := tx.Model(&rank).Updates(map[string]interface{}{
// 				"level_id":        nextLevel.ID,
// 				"max_point_level": nextLevel.MaxPointLevel,
// 				"icon_url":        nextLevel.IconURL,
// 			}).Error; err != nil {
// 				log.Println("[ERROR] Gagal upgrade level:", err)
// 			} else {
// 				log.Printf("[INFO] User ID %d naik level ke %s", userID, nextLevel.NameLevel)
// 			}
// 		} else {
// 			log.Println("[INFO] Updated existing user_point_level_rank (incremented)")
// 		}
// 	}
// }

func incrementAmountTotalQuiz(tx *gorm.DB, userID uint, point int) {
	var rank progress_user.UserPointLevelRank
	if err := tx.Where("user_id = ?", userID).First(&rank).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Ambil level dan rank default
			var level progress_user.LevelPointRequirement
			var rankReq progress_user.RankLevelRequirement

			if err := tx.First(&level, 1).Error; err != nil {
				log.Println("[ERROR] Gagal ambil level ID 1:", err)
				return
			}
			if err := tx.First(&rankReq, 1).Error; err != nil {
				log.Println("[ERROR] Gagal ambil rank ID 1:", err)
				return
			}

			// Inisialisasi rank user baru
			rank = progress_user.UserPointLevelRank{
				UserID:          userID,
				AmountTotalQuiz: point,
				LevelID:         level.ID,
				MaxPointLevel:   level.MaxPointLevel,
				IconURLLevel:    level.IconURL,
				RankID:          rankReq.ID,
				MaxLevelRank:    rankReq.MaxLevel,
				IconURLRank:     rankReq.IconURL,
			}

			if err := tx.Create(&rank).Error; err != nil {
				log.Println("[ERROR] Gagal create user_point_level_rank:", err)
			} else {
				log.Println("[INFO] Created new user_point_level_rank (Level 1, Rank 1)")
			}
			return
		} else {
			log.Println("[ERROR] Fetch rank error:", err)
			return
		}
	}

	// Tambah poin
	if err := tx.Model(&rank).
		Update("amount_total_quiz", gorm.Expr("amount_total_quiz + ?", point)).Error; err != nil {
		log.Println("[ERROR] Gagal update total point:", err)
		return
	}

	// Ambil ulang data rank setelah update
	if err := tx.Where("user_id = ?", userID).First(&rank).Error; err != nil {
		log.Println("[ERROR] Gagal ambil data rank setelah update:", err)
		return
	}

	// Cek upgrade level
	needUpdate := false
	var nextLevel progress_user.LevelPointRequirement
	if rank.AmountTotalQuiz > rank.MaxPointLevel {
		if err := tx.Where("max_point_level > ?", rank.MaxPointLevel).
			Order("max_point_level ASC").First(&nextLevel).Error; err == nil {
			rank.LevelID = nextLevel.ID
			rank.MaxPointLevel = nextLevel.MaxPointLevel
			rank.IconURLLevel = nextLevel.IconURL
			needUpdate = true
			log.Printf("[INFO] User ID %d naik LEVEL ke %s", userID, nextLevel.NameLevel)
		}
	}

	// Cek upgrade rank
	var nextRank progress_user.RankLevelRequirement
	if int(rank.LevelID) > rank.MaxLevelRank {
		if err := tx.Where("max_level > ?", rank.MaxLevelRank).
			Order("max_level ASC").First(&nextRank).Error; err == nil {
			rank.RankID = nextRank.ID
			rank.MaxLevelRank = nextRank.MaxLevel
			rank.IconURLRank = nextRank.IconURL
			needUpdate = true
			log.Printf("[INFO] User ID %d naik RANK ke %s", userID, nextRank.NameRank)
		}
	}

	// Simpan perubahan jika ada
	if needUpdate {
		if err := tx.Save(&rank).Error; err != nil {
			log.Println("[ERROR] Gagal update rank/level:", err)
		}
	} else {
		log.Println("[INFO] Updated existing user_point_level_rank (incremented)")
	}
}
