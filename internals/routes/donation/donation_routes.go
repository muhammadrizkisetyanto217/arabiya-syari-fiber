package routes

import (
	controllers "arabiya-syari-fiber/internals/controllers/donation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DonationRoutes(app *fiber.App, db *gorm.DB) {
	donationLevelsController := controllers.NewDonationLevelsController(db)

	donationLevelsGroup := app.Group("/api/donation-levels")

	// Donation Levels API
	donationLevelsGroup.Get("/", donationLevelsController.GetAll)
	donationLevelsGroup.Get("/:id", donationLevelsController.GetByID)
	donationLevelsGroup.Post("/", donationLevelsController.Create)
	donationLevelsGroup.Put("/:id", donationLevelsController.Update)
	donationLevelsGroup.Delete("/:id", donationLevelsController.Delete)
	

	donationStatsController := controllers.NewUserDonationLogsController(db)

	api := app.Group("/api/user-donation-logs")

	// Donation Stats API
	api.Get("/", donationStatsController.GetAll)
	api.Get("/:id", donationStatsController.GetByID)
	api.Post("/", donationStatsController.Create)
	api.Put("/:id", donationStatsController.Update)
	api.Delete("/:id", donationStatsController.Delete)

}
