package routes

import (
	"arabiya-syari-fiber/internals/controllers/donation"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DonationRoutes(app *fiber.App, db *gorm.DB) {
	donationLevelsController := controllers.NewDonationLevelsController(db)

	api := app.Group("/api")

	// Donation Levels API
	api.Get("/donation-levels", donationLevelsController.GetAll)
	api.Get("/donation-levels/:id", donationLevelsController.GetByID)
	api.Post("/donation-levels", donationLevelsController.Create)
	api.Put("/donation-levels/:id", donationLevelsController.Update)
	api.Delete("/donation-levels/:id", donationLevelsController.Delete)
}
