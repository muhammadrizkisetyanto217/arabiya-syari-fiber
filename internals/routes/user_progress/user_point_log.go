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
	controllerRank := user_progress.NewRankLevelRequirementController(db)

	// User Point Logs
	app.Get("/api/user-point-logs/:user_id", controllerPointLogs.GetUserPointLogsByUserID)

	// User Progress (Total Quiz & Level)
	app.Get("/api/user-progress/:user_id", controllerGetUserPointLevelRank.GetUserPointLevelRankByUserID)

	// Level Point Requirement Routes
	app.Get("/api/level-point-requirements", controllerLevel.GetLevels)
	app.Get("/api/level-point-requirements/:id", controllerLevel.GetLevel)
	app.Post("/api/level-point-requirements", controllerLevel.CreateLevel)
	app.Put("/api/level-point-requirements/:id", controllerLevel.UpdateLevel)
	app.Delete("/api/level-point-requirements/:id", controllerLevel.DeleteLevel)

	// Rank Level Requirement Routes
	app.Get("/api/rank-level-requirements", controllerRank.GetAll)
	app.Get("/api/rank-level-requirements/:id", controllerRank.GetByID)
	app.Post("/api/rank-level-requirements", controllerRank.Create)
	app.Put("/api/rank-level-requirements/:id", controllerRank.Update)
	app.Delete("/api/rank-level-requirements/:id", controllerRank.Delete)
}
