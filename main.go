package main

import (
	"arabiya-syari-fiber/database"
	// "arabiya-syari-fiber/models"
	"arabiya-syari-fiber/routes"
	"log"

	// "myapp/database"
	// "myapp/models"
	// "myapp/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Koneksi ke database
	database.ConnectDB()

	// Auto-migrate model ke database
	// database.DB.AutoMigrate(&models.User{})

	app := fiber.New()

	// Middleware
	app.Use(logger.New()) // Logging request
	app.Use(cors.New(cors.Config{
    AllowOrigins: "*", // Mengizinkan semua origin
    // AllowCredentials: true, // Jangan gunakan ini bersamaan dengan "*"
}))
	// Setup routes
	routes.SetupUserRoutes(app)

	// Start server
	port := ":8080"
	log.Println("Server running on port", port)
	log.Fatal(app.Listen(port))
}
