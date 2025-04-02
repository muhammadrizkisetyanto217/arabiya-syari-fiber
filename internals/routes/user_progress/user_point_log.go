package user_progress

import (
	"arabiya-syari-fiber/internals/controllers/user_progress" // sesuaikan path sesuai struktur project kamu
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserPointLogRoutes(app *fiber.App, db *gorm.DB) {
	controllerPointLogs := progress_user.NewUserPointLogController(db)

	// GET /api/user_point_logs/:user_id
	app.Get("/api/user_point_logs/:user_id", controllerPointLogs.GetUserPointLogsByUserID)

	controllerGetUserPointLevelRank := progress_user.NewUserPointLevelRankController(db)
	// GET /api/user_progress/:user_id
	app.Get("/api/user_progress/:user_id", controllerGetUserPointLevelRank.GetUserPointLevelRankByUserID)
}
