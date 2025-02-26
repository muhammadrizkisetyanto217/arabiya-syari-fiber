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
}
