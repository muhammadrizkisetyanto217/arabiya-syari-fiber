package user

import (
	"arabiya-syari-fiber/internals/controllers/user" // ✅ Hanya satu alias

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes: Register semua routes terkait Users Profile
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Inisialisasi controller
	usersProfileController := user.NewUsersProfileController(db)

	// 🔒 Gunakan Middleware Auth untuk melindungi semua route `/api/*`
	api := app.Group("/api", user.AuthMiddleware(db))

	// 🎯 Users Profile Routes
	usersProfileRoutes := api.Group("/users-profiles")
	usersProfileRoutes.Get("/", usersProfileController.GetProfiles)
	usersProfileRoutes.Get("/:id", usersProfileController.GetProfile)
	usersProfileRoutes.Post("/", usersProfileController.CreateProfile)
	usersProfileRoutes.Put("/:id", usersProfileController.UpdateProfile)
	usersProfileRoutes.Delete("/:id", usersProfileController.DeleteProfile)
}
