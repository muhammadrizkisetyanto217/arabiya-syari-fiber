package routes

import (
	controllers "arabiya-syari-fiber/internals/controllers/donation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DonationRoutes(app *fiber.App, db *gorm.DB) {
	donationLevelsController := controllers.NewDonationLevelsController(db)

	api := app.Group("/api/donation-levels")

	// Donation Levels API
	api.Get("/", donationLevelsController.GetAll)
	api.Get("/:id", donationLevelsController.GetByID)
	api.Post("/", donationLevelsController.Create)
	api.Put("/:id", donationLevelsController.Update)
	api.Delete("/:id", donationLevelsController.Delete)
}
