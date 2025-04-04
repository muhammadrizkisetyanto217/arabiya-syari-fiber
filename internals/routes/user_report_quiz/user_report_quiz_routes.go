package report_user

import (
	"arabiya-syari-fiber/internals/controllers/user" // Middleware Auth
	report_user_quiz "arabiya-syari-fiber/internals/controllers/user_report_quiz"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ReportUserRoutes: Register semua routes terkait laporan user
func ReportUserRoutes(app *fiber.App, db *gorm.DB) {

	// 🔒 Middleware Auth diterapkan untuk seluruh API /api/*
	api := app.Group("/api", user.AuthMiddleware(db))

	// 📦 User Units Routes
	userUnitController := report_user_quiz.NewUserUnitController(db)
	userUnitRoutes := api.Group("/user-units")
	userUnitRoutes.Get("/:user_id", userUnitController.GetUserUnitsByUserID)

	// 📖 User Reading Routes
	userReadingController := report_user_quiz.NewUserReadingController(db)
	userReadingRoutes := api.Group("/user-readings")
	userReadingRoutes.Post("/", userReadingController.Create)
	userReadingRoutes.Get("/", userReadingController.GetAll)

	// 📚 Reading Saved Routes
	readingSavedController := report_user_quiz.NewReadingSavedController(db)
	readingSavedRoutes := api.Group("/reading-saved")
	readingSavedRoutes.Post("/save", readingSavedController.SaveReading)
	readingSavedRoutes.Delete("/unsave", readingSavedController.UnsaveReading)
	readingSavedRoutes.Get("/:user_id", readingSavedController.GetSavedReadings)

	// 🎯 User Quizzes Routes
	userQuizzesController := report_user_quiz.NewUserQuizzesController(db)
	userQuizzesRoutes := api.Group("/user-quizzes")
	userQuizzesRoutes.Post("/save", userQuizzesController.CreateOrUpdateUserQuiz)
	userQuizzesRoutes.Get("/:user_id", userQuizzesController.GetUserQuizzes)

	// ✅ User Section Quizzes Routes
	userSectionQuizzesController := report_user_quiz.NewUserSectionQuizzesController(db)
	userSectionQuizzesRoutes := api.Group("/user-section-quizzes")
	userSectionQuizzesRoutes.Get("/:user_id", userSectionQuizzesController.GetUserSectionQuizzes)

	// 📝 User Evaluation Routes
	userEvaluationController := report_user_quiz.NewUserEvaluationController(db)
	userEvaluationRoutes := api.Group("/user-evaluations")
	userEvaluationRoutes.Post("/", userEvaluationController.CreateOrUpdateUserEvaluation)
	userEvaluationRoutes.Get("/:user_id", userEvaluationController.GetByUserID)

	// 📌 Question Saved Routes
	questionSavedController := report_user_quiz.NewQuestionSavedController(db)
	questionSavedRoutes := api.Group("/question-saved")
	questionSavedRoutes.Post("/", questionSavedController.Create)
	questionSavedRoutes.Get("/:user_id", questionSavedController.GetByUserID)
	questionSavedRoutes.Get("/detail/:id", questionSavedController.GetByID)
	questionSavedRoutes.Put("/:id", questionSavedController.Update)
	questionSavedRoutes.Delete("/:id", questionSavedController.Delete)

}
