package main

import (
	"log"

	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/database"
	"github.com/codescalersinternships/Secret-note-API-SPA--Basem/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDatabase()

	r := gin.Default()

	routes.RegisterRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
