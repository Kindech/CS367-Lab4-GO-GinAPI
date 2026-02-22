package main

import (
	"log"

	"go-api-gin/config"
	"go-api-gin/handlers"
	"go-api-gin/repositories"
	"go-api-gin/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Config Database
	db := config.ConnectDB()
	defer db.Close()

	// 2. Setup Layers (Dependency Injection)
	repo := repositories.NewStudentRepository(db)
	service := services.NewStudentService(repo)
	handler := handlers.NewStudentHandler(service)

	// 3. Setup Router
	r := gin.Default()

	r.GET("/students", handler.GetStudents)
	r.GET("/students/:id", handler.GetStudentByID)
	r.POST("/students", handler.CreateStudent)
	r.PUT("/students/:id", handler.UpdateStudent)
	r.DELETE("/students/:id", handler.DeleteStudent)

	// Start Server
	log.Println("Server is running on port 8080...")
	r.Run(":8080")
}
