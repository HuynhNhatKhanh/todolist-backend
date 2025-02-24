package delivery

import (
	"net/http"
	"strconv"
	"todolist/internal/domain"
	"todolist/internal/usecase"
	"todolist/pkg/errors"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	useCase *usecase.TaskUseCase
}

// NewTaskHandler creates a new TaskHandler and registers its methods with the
// given Gin engine.
func NewTaskHandler(r *gin.Engine, useCase *usecase.TaskUseCase) {
	handler := &TaskHandler{useCase: useCase}

	// Create a new API group for tasks.
	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks", handler.ListTasks)
	r.PUT("/tasks/:id", handler.UpdateTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)
}

// CreateTask handles the creation of a new task.
// It binds the JSON request to a TaskCreate struct, validates it,
// and then calls the use case to create the task.
func (h *TaskHandler) CreateTask(c *gin.Context) {
	// Bind the JSON request body to a TaskCreate struct
	var req domain.TaskCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(errors.ErrInvalidRequest)
		return
	}

	// Call the use case to create a new task
	id, err := h.useCase.CreateTask(req.Title, req.Description)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// ListTasks handles the request to list all tasks.
// It retrieves the list of tasks from the use case and returns them as a JSON response.
func (h *TaskHandler) ListTasks(c *gin.Context) {
	// Retrieve the list of tasks from the use case
	tasks, err := h.useCase.ListTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// UpdateTask handles the request to update a task.
// It extracts the task ID from the URL, binds the JSON request body to a TaskRequest struct,
// validates it, and then calls the use case to update the task.
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	// Extract the task ID from the URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errors.ErrInvalidTaskID)
		return
	}

	// Bind the JSON request body to a TaskUpdate struct
	var req domain.TaskUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Call the use case to update the task
	updatedTask, err := h.useCase.UpdateTask(id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask handles the request to delete a task.
// It extracts the task ID from the URL, validates it, and then calls the use case to delete the task.
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	// Extract the task ID from the URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(errors.ErrInvalidTaskID)
		return
	}

	// Call the use case to delete the task
	if err := h.useCase.DeleteTask(id); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}
