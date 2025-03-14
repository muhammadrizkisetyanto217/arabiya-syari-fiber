package category

import (
	"arabiya-syari-fiber/internals/controllers/quizzes"
	"arabiya-syari-fiber/internals/controllers/user" // Middleware Auth

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// QuizzesRoutes: Register semua routes terkait quizzes & evaluasi
func QuizzesRoutes(app *fiber.App, db *gorm.DB) {

	// 🔒 Middleware Auth diaktifkan untuk seluruh API /api/*
	api := app.Group("/api", user.AuthMiddleware(db))

	// 📖 Reading Routes
	readingController := quizzes.NewReadingController(db)
	readingRoutes := api.Group("/readings")
	readingRoutes.Get("/", readingController.GetReadings)
	readingRoutes.Get("/:id", readingController.GetReading)
	readingRoutes.Get("/unit/:unitId", readingController.GetReadingsByUnit)
	readingRoutes.Post("/", readingController.CreateReading)
	readingRoutes.Put("/:id", readingController.UpdateReading)
	readingRoutes.Delete("/:id", readingController.DeleteReading)
	readingRoutes.Get("/:id/readingTooltips", readingController.GetReadingWithTooltips)
	readingRoutes.Get("/:id/convertTooltips", readingController.ConvertReadingWithTooltipsId)
	readingRoutes.Get("/:id/onlyTooltips", readingController.GetOnlyReadingTooltips)


	// 🔥 Section Quizzes Routes
	sectionQuizzesController := quizzes.NewSectionQuizController(db)
	sectionQuizzesRoutes := api.Group("/section-quizzes")
	sectionQuizzesRoutes.Get("/", sectionQuizzesController.GetSectionQuizzes)
	sectionQuizzesRoutes.Get("/:id", sectionQuizzesController.GetSectionQuiz)
	sectionQuizzesRoutes.Get("/unit/:unitId", sectionQuizzesController.GetSectionQuizzesByUnit)
	sectionQuizzesRoutes.Post("/", sectionQuizzesController.CreateSectionQuiz)
	sectionQuizzesRoutes.Put("/:id", sectionQuizzesController.UpdateSectionQuiz)
	sectionQuizzesRoutes.Delete("/:id", sectionQuizzesController.DeleteSectionQuiz)

	// 🧠 Quiz Routes
	quizController := quizzes.NewQuizController(db)
	quizRoutes := api.Group("/quizzes")
	quizRoutes.Get("/", quizController.GetQuizzes)
	quizRoutes.Get("/:id", quizController.GetQuiz)
	quizRoutes.Get("/section/:sectionId", quizController.GetQuizzesBySection)
	quizRoutes.Post("/", quizController.CreateQuiz)
	quizRoutes.Put("/:id", quizController.UpdateQuiz)
	quizRoutes.Delete("/:id", quizController.DeleteQuiz)

	// 📝 Quiz Questions Routes
	quizQuestionController := quizzes.NewQuizQuestionController(db)
	quizQuestionRoutes := api.Group("/quiz-questions")
	quizQuestionRoutes.Get("/", quizQuestionController.GetQuizQuestions)
	quizQuestionRoutes.Get("/:id", quizQuestionController.GetQuizQuestion)
	quizQuestionRoutes.Get("/quiz/:quizId", quizQuestionController.GetQuizQuestionsByQuizID)
	quizQuestionRoutes.Post("/", quizQuestionController.CreateQuizQuestion)
	quizQuestionRoutes.Put("/:id", quizQuestionController.UpdateQuizQuestion)
	quizQuestionRoutes.Delete("/:id", quizQuestionController.DeleteQuizQuestion)

	// 🎓 Exam Routes
	examController := quizzes.NewExamController(db)
	examRoutes := api.Group("/exams")
	examRoutes.Get("/", examController.GetExams)
	examRoutes.Get("/:id", examController.GetExam)
	examRoutes.Get("/unit/:unitId", examController.GetExamsByUnitID)
	examRoutes.Post("/", examController.CreateExam)
	examRoutes.Put("/:id", examController.UpdateExam)
	examRoutes.Delete("/:id", examController.DeleteExam)

	// 📋 Exam Questions Routes
	examQuestionController := quizzes.NewExamsQuestionController(db)
	examQuestionRoutes := api.Group("/exam-questions")
	examQuestionRoutes.Get("/", examQuestionController.GetExamsQuestions)
	examQuestionRoutes.Get("/:id", examQuestionController.GetExamsQuestion)
	examQuestionRoutes.Get("/exam/:examId", examQuestionController.GetQuestionsByExamID)
	examQuestionRoutes.Post("/", examQuestionController.CreateExamsQuestion)
	examQuestionRoutes.Put("/:id", examQuestionController.UpdateExamsQuestion)
	examQuestionRoutes.Delete("/:id", examQuestionController.DeleteExamsQuestion)

	// 🏆 Evaluation Routes
	evaluationController := quizzes.NewEvaluationController(db)
	evaluationRoutes := api.Group("/evaluations")
	evaluationRoutes.Get("/", evaluationController.GetEvaluations)
	evaluationRoutes.Get("/:id", evaluationController.GetEvaluation)
	evaluationRoutes.Get("/unit/:unitId", evaluationController.GetEvaluationsByUnitID)
	evaluationRoutes.Post("/", evaluationController.CreateEvaluation)
	evaluationRoutes.Put("/:id", evaluationController.UpdateEvaluation)
	evaluationRoutes.Delete("/:id", evaluationController.DeleteEvaluation)

	// 🎯 Evaluation Questions Routes
	evaluationQuestionController := quizzes.NewEvaluationsQuestionController(db)
	evaluationQuestionRoutes := api.Group("/evaluation-questions")
	evaluationQuestionRoutes.Get("/", evaluationQuestionController.GetEvaluationsQuestions)
	evaluationQuestionRoutes.Get("/:id", evaluationQuestionController.GetEvaluationQuestion)
	evaluationQuestionRoutes.Get("/evaluation/:evaluationId", evaluationQuestionController.GetQuestionsByEvaluationID)
	evaluationQuestionRoutes.Post("/", evaluationQuestionController.CreateEvaluationQuestion)
	evaluationQuestionRoutes.Put("/:id", evaluationQuestionController.UpdateEvaluationQuestion)
	evaluationQuestionRoutes.Delete("/:id", evaluationQuestionController.DeleteEvaluationQuestion)
}
