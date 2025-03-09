package routes

import (
	"arabiya-syari-fiber/internals/controllers/donation"
	authControllers "arabiya-syari-fiber/internals/controllers/user" // Middleware Auth

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// DonationRoutes: Register semua routes terkait donasi
func DonationRoutes(app *fiber.App, db *gorm.DB) {

	// 🔒 Middleware Auth diterapkan untuk seluruh API /api/*
	api := app.Group("/api", authControllers.AuthMiddleware(db))

	// 🎯 Donation Levels Routes
	donationLevelsController := donation.NewDonationLevelsController(db)
	donationLevelsRoutes := api.Group("/donation-levels")
	donationLevelsRoutes.Get("/", donationLevelsController.GetAll)
	donationLevelsRoutes.Get("/:id", donationLevelsController.GetByID)
	donationLevelsRoutes.Post("/", donationLevelsController.Create)
	donationLevelsRoutes.Put("/:id", donationLevelsController.Update)
	donationLevelsRoutes.Delete("/:id", donationLevelsController.Delete)

	// 💰 User Donation Logs Routes
	donationLogsController := donation.NewUserDonationLogsController(db)
	donationLogsRoutes := api.Group("/user-donation-logs")
	donationLogsRoutes.Get("/", donationLogsController.GetAll)
	donationLogsRoutes.Get("/:id", donationLogsController.GetByID)
	donationLogsRoutes.Post("/", donationLogsController.Create)
	donationLogsRoutes.Put("/:id", donationLogsController.Update)
	donationLogsRoutes.Delete("/:id", donationLogsController.Delete)
}
