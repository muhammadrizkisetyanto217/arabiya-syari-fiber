package main

import (
	"log"

	"arabiya-syari-fiber/internals/configs"
	"arabiya-syari-fiber/internals/database"
	"arabiya-syari-fiber/internals/middleware"
	"arabiya-syari-fiber/internals/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// ✅ Load environment variables
	configs.LoadEnv()

	// ✅ Koneksi ke database
	database.ConnectDB()

	// ✅ Inisialisasi Fiber App
	app := fiber.New()

	// ✅ Setup Middleware
	middleware.SetupMiddleware(app)

	// ✅ Setup Routes
	routes.SetupRoutes(app, database.DB)

	// ✅ Ambil PORT dari .env atau default 8080
	port := configs.GetEnv("PORT")
	serverAddress := ":" + port

	// ✅ Start Server
	log.Println("🚀 Server running on", serverAddress)
	log.Fatal(app.Listen(serverAddress))
}
