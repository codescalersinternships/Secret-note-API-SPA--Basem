package main

import (
	"log"
	"os"

	controllers "github.com/codescalersinternships/Secret-note-API-SPA--Basem/controller"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/database"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/routes"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Secret Note API
// @version 1.0
// @description This is a simple API for creating and managing secret notes.

// @host localhost:8080

func main() {
	database := database.NewDB()
	database.ConnectDatabase()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed loading .env file: %v", err)
	}

	frontendBaseURL := os.Getenv("FRONTEND_BASE_URL")
	port := os.Getenv("PORT")

	r := gin.Default()
	r.Use(utils.RateLimiterMiddleware(1, 5))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendBaseURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	noteController := controllers.NewNoteController(database)
	userController := controllers.NewUserController(database)
	routes.RegisterRoutes(r, noteController, userController)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
