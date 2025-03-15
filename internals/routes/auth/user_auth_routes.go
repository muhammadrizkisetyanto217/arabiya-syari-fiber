package auth

import (
	"arabiya-syari-fiber/internals/controllers/user"
	"arabiya-syari-fiber/internals/controllers/category"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes mengatur routing untuk authentication dan user
func UserRoutes(app *fiber.App, db *gorm.DB) {

	//* Jika tanpa constructor
	// userController := controllers.UserController{DB: db}

	//* Dengan constructor
	authController := user.NewAuthController(db)

	// 🔥 Setup AuthController
	auth := app.Group("/auth")
	auth.Post("/register", authController.Register) // ✅ Register user baru
	auth.Post("/login", authController.Login)       // ✅ Login user

	auth.Post("/forgot-password/check", authController.CheckSecurityAnswer) // validasi email dan jawaban keamanan
	auth.Post("/forgot-password/reset", authController.ResetPassword)       // reset password setelah validasi berhasil
	
	// 🔥 Setup Category
	categoryController := category.NewCategoryController(db)
	categoryRoutes := app.Group("/api/category") // ✅ Proteksi semua category route
	categoryRoutes.Get("/", categoryController.GetCategories)            // ✅ Get semua category (Hanya Admin)
	categoryRoutes.Get("/:id", categoryController.GetCategory)            // ✅ Get satu category berdasarkan ID
	categoryRoutes.Post("/", categoryController.CreateCategory)           // ✅ Buat category baru
	categoryRoutes.Put("/:id", categoryController.UpdateCategory)        // ✅ Update category
	categoryRoutes.Delete("/:id", categoryController.DeleteCategory)     // ✅ Hapus category

	// 🔥 Setup AuthController with middleware
	protectedRoutes := app.Group("/api/auth", user.AuthMiddleware(db))
	protectedRoutes.Post("/logout", authController.Logout)                  // ✅ Logout User
	protectedRoutes.Post("/change-password", authController.ChangePassword) // ✅ Ganti Password User

	// 🔥 Setup UserController (dengan middleware untuk proteksi API)
	userController := user.NewUserController(db)
	userRoutes := app.Group("/api/users", user.AuthMiddleware(db)) // ✅ Proteksi semua user route
	userRoutes.Get("/", userController.GetUsers)                   // ✅ Get semua users (Hanya Admin)
	userRoutes.Get("/:id", userController.GetUser)                 // ✅ Get satu user berdasarkan ID
	userRoutes.Put("/:id", userController.UpdateUser)              // ✅ Update user
	userRoutes.Delete("/:id", userController.DeleteUser)           // ✅ Hapus user
}
