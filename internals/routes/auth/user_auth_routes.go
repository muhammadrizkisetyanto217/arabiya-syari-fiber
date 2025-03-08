package routes

import (
	controllers "arabiya-syari-fiber/internals/controllers/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes mengatur routing untuk authentication dan user
func UserRoutes(app *fiber.App, db *gorm.DB) {

	//* Jika tanpa constructor
	// userController := controllers.UserController{DB: db}

	//* Dengan constructor
	authController := controllers.NewAuthController(db)

	// 🔥 Setup AuthController
	auth := app.Group("/auth") 
	auth.Post("/register", authController.Register) // ✅ Register user baru
	auth.Post("/login", authController.Login)       // ✅ Login user

	// 🔥 Setup UserController (dengan middleware untuk proteksi API)
	userController := controllers.NewUserController(db)
	userRoutes := app.Group("/api/users", controllers.AuthMiddleware) // ✅ Proteksi semua user route
	userRoutes.Get("/", userController.GetUsers)   // ✅ Get semua users (Hanya Admin)
	userRoutes.Get("/:id", userController.GetUser) // ✅ Get satu user berdasarkan ID
	userRoutes.Put("/:id", userController.UpdateUser) // ✅ Update user
	userRoutes.Delete("/:id", userController.DeleteUser) // ✅ Hapus user
}
