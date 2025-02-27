package main

import (
	"arabiya-syari-fiber/internals/database"
	// "arabiya-syari-fiber/models"
	"arabiya-syari-fiber/internals/routes"
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
    AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
    AllowHeaders: "Origin, Content-Type, Accept, Authorization",
    // AllowCredentials: true,
}))
	// Setup routes
	routes.SetupRoutes(app, database.DB)

	// Start server
	port := ":8080"
	log.Println("Server running on port", port)
	log.Fatal(app.Listen(port))
}
