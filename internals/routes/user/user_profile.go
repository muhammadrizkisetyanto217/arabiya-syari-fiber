package user

import (
	controllers "arabiya-syari-fiber/internals/controllers/user"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Inisialisasi controller
	usersProfileController := controllers.NewUsersProfileController(db)

	// Grouping routes untuk Users Profile
	usersProfileRoutes := app.Group("/users-profiles")
	usersProfileRoutes.Get("/", usersProfileController.GetProfiles)
	usersProfileRoutes.Get("/:id", usersProfileController.GetProfile)
	usersProfileRoutes.Post("/", usersProfileController.CreateProfile)
	usersProfileRoutes.Put("/:id", usersProfileController.UpdateProfile)
	usersProfileRoutes.Delete("/:id", usersProfileController.DeleteProfile)

	// Log semua route yang terdaftar
	// Log semua route yang terdaftar
	for _, routes := range app.Stack() {
		for _, route := range routes {
			log.Println("Method:", route.Method, "Path:", route.Path)
		}
	}

}
