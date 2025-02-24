package main

import (
	"todolist/config"
	"todolist/internal/delivery"
	"todolist/internal/repository"
	"todolist/internal/usecase"
	"todolist/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin engine.
	r := gin.Default()

	// Enable CORS for the API.
	r.Use(config.CorsConfig())

	// Register error handler middleware.
	r.Use(middleware.ErrorHandler())

	// Create a new task repository.
	repo := repository.NewTaskRepository()

	// Create a new task use case with the repository.
	useCase := usecase.NewTaskUseCase(repo)

	// Register the task handler with the use case.
	delivery.NewTaskHandler(r, useCase)

	// Run the server.
	r.Run(":8080")
}
