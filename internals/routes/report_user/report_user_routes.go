package report_user

import (
	"arabiya-syari-fiber/internals/controllers/report_user"
	"arabiya-syari-fiber/internals/controllers/user" // Middleware Auth

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ReportUserRoutes: Register semua routes terkait laporan user
func ReportUserRoutes(app *fiber.App, db *gorm.DB) {

	// 🔒 Middleware Auth diterapkan untuk seluruh API /api/*
	api := app.Group("/api", user.AuthMiddleware(db))

	// 📚 Reading Saved Routes
	readingSavedController := report_user.NewReadingSavedController(db)
	readingSavedRoutes := api.Group("/reading_saved")
	readingSavedRoutes.Post("/save", readingSavedController.SaveReading)
	readingSavedRoutes.Delete("/unsave", readingSavedController.UnsaveReading)
	readingSavedRoutes.Get("/:user_id", readingSavedController.GetSavedReadings)

	// 🎯 User Quizzes Routes
	userQuizzesController := report_user.NewUserQuizzesController(db)
	userQuizzesRoutes := api.Group("/user_quizzes")
	userQuizzesRoutes.Post("/save", userQuizzesController.CreateOrUpdateUserQuiz)
	userQuizzesRoutes.Get("/:user_id", userQuizzesController.GetUserQuizzes)

	// ✅ User Section Quizzes Routes (BARU)
	userSectionQuizzesController := report_user.NewUserSectionQuizzesController(db)
	userSectionQuizzesRoutes := api.Group("/user_section_quizzes")
	userSectionQuizzesRoutes.Get("/:user_id", userSectionQuizzesController.GetUserSectionQuizzes)
}
