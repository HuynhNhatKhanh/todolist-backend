package usecase

import (
	"strings"
	"todolist/internal/domain"
	"todolist/internal/repository"
	"todolist/pkg/errors"
)

type TaskUseCase struct {
	repo *repository.TaskRepository
}

// NewTaskUseCase creates a new instance of TaskUseCase.
// It initializes the use case with the provided task repository.
func NewTaskUseCase(repo *repository.TaskRepository) *TaskUseCase {
	// Return a new TaskUseCase with the injected repository.
	return &TaskUseCase{repo: repo}
}

// CreateTask creates a new task in the repository.
// It returns the ID of the newly created task or an error if the title is empty.
func (u *TaskUseCase) CreateTask(title, description string) (int, error) {
	// Check if the title is empty.
	if strings.TrimSpace(title) == "" {
		return 0, errors.ErrTitleRequired
	}

	// Create a new task with the provided title, description, and completed flag.
	task := domain.Task{
		Title:       title,
		Description: description,
		Completed:   false,
	}

	// Create the task in the repository and return its ID.
	return u.repo.Create(task), nil
}

func (u *TaskUseCase) ListTasks() ([]domain.Task, error) {
	tasks, err := u.repo.List()
	if err != nil {
		return nil, errors.ErrDatabase
	}
	if tasks == nil {
		return []domain.Task{}, nil
	}
	return tasks, nil
}

// UpdateTask updates a task by its ID.
// It returns an error if the task does not exist.
func (u *TaskUseCase) UpdateTask(id int, req domain.TaskUpdate) (*domain.Task, error) {
	// Check if the task exists in the repository.
	if !u.repo.Exists(id) {
		return nil, errors.ErrTaskNotFound
	}

	// Get the existing task from the repository.
	existingTask, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Check if the title is empty.
	if req.Title != nil && strings.TrimSpace(*req.Title) == "" {
		return nil, errors.ErrTitleRequired
	}

	task := domain.Task{
		ID:          id,
		Title:       existingTask.Title,
		Description: existingTask.Description,
		Completed:   existingTask.Completed,
	}

	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = *req.Description
	}
	if req.Completed != nil {
		task.Completed = *req.Completed
	}

	// Update the task in the repository.
	return u.repo.Update(id, task)
}

// DeleteTask deletes a task by its ID.
// It returns an error if the task does not exist.
func (u *TaskUseCase) DeleteTask(id int) error {
	// Check if the task exists in the repository.
	if !u.repo.Exists(id) {
		return errors.ErrTaskNotFound
	}
	// Delete the task from the repository.
	return u.repo.Delete(id)
}
