package main

import (
	"log"
	"os"

	"caselog/internal/db"
	"caselog/internal/handler"
	"caselog/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	projectRepository := repository.NewProjectRepository(database)
	projectHandler := handler.NewProjectHandler(projectRepository)

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", projectHandler.Index)
	r.GET("/projects/new", projectHandler.New)
	r.POST("/projects", projectHandler.Create)
	r.GET("/projects/:id", projectHandler.Show)

	// r.Run(":8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run("0.0.0.0:" + port)
}