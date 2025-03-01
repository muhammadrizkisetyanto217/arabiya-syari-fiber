package category

import (
	controllers "arabiya-syari-fiber/internals/controllers/quizzes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func QuizzesRoutes(app *fiber.App, db *gorm.DB) {

	readingController := controllers.NewReadingController(db)

	// Reading routes
	readingGroup := app.Group("/api/readings")
	readingGroup.Get("/", readingController.GetReadings)
	readingGroup.Get("/:id", readingController.GetReading)
	readingGroup.Get("/unit/:unitId", readingController.GetReadingsByUnit)
	readingGroup.Post("/", readingController.CreateReading)
	readingGroup.Put("/:id", readingController.UpdateReading)
	readingGroup.Delete("/:id", readingController.DeleteReading)

	sectionQuizzesController := controllers.NewSectionQuizController(db)

	// Section Quizzes routes
	sectionQuizzesGroup := app.Group("/api/section-quizzes")
	sectionQuizzesGroup.Get("/", sectionQuizzesController.GetSectionQuizzes)
	sectionQuizzesGroup.Get("/:id", sectionQuizzesController.GetSectionQuiz)
	sectionQuizzesGroup.Get("/unit/:unitId", sectionQuizzesController.GetSectionQuizzesByUnit)
	sectionQuizzesGroup.Post("/", sectionQuizzesController.CreateSectionQuiz)
	sectionQuizzesGroup.Put("/:id", sectionQuizzesController.UpdateSectionQuiz)
	sectionQuizzesGroup.Delete("/:id", sectionQuizzesController.DeleteSectionQuiz)


		// Tambahkan routing untuk quizzes
	quizController := controllers.NewQuizController(db)

	// Quizzes routes
	quizGroup := app.Group("/api/quizzes")
	quizGroup.Get("/", quizController.GetQuizzes)
	quizGroup.Get("/:id", quizController.GetQuiz)
	quizGroup.Get("/unit/:unitId", quizController.GetQuizzesByUnit)
	quizGroup.Post("/", quizController.CreateQuiz)
	quizGroup.Put("/:id", quizController.UpdateQuiz)
	quizGroup.Delete("/:id", quizController.DeleteQuiz)
}
