package user_progress

import (
	"arabiya-syari-fiber/internals/controllers/user_progress"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserPointLogRoutes(app *fiber.App, db *gorm.DB) {
	controllerPointLogs := user_progress.NewUserPointLogController(db)
	controllerGetUserPointLevelRank := user_progress.NewUserPointLevelRankController(db)
	controllerLevel := user_progress.NewLevelPointRequirementController(db)

	// User Point Logs
	app.Get("/api/user_point_logs/:user_id", controllerPointLogs.GetUserPointLogsByUserID)

	// User Progress (Total Quiz & Level)
	app.Get("/api/user_progress/:user_id", controllerGetUserPointLevelRank.GetUserPointLevelRankByUserID)

	// Level Point Requirement Routes
	app.Get("/api/level_point_requirements", controllerLevel.GetLevels)
	app.Get("/api/level_point_requirements/:id", controllerLevel.GetLevel)
	app.Post("/api/level_point_requirements", controllerLevel.CreateLevel)
	app.Put("/api/level_point_requirements/:id", controllerLevel.UpdateLevel)
	app.Delete("/api/level_point_requirements/:id", controllerLevel.DeleteLevel)
}
