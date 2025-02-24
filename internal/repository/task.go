package repository

import (
	"sync"
	"todolist/internal/domain"
	"todolist/pkg/errors"
)

type TaskRepository struct {
	tasks  map[int]domain.Task
	nextID int
	mu     sync.Mutex
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make(map[int]domain.Task),
		nextID: 1,
	}
}

// Create creates a new task in the repository.
// It assigns a unique ID to the task and returns this ID.
func (r *TaskRepository) Create(task domain.Task) int {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Assign a unique ID to the task
	task.ID = r.nextID
	// Store the task in the repository
	r.tasks[r.nextID] = task
	// Increment the next ID
	r.nextID++
	// Return the ID of the newly created task
	return task.ID
}

// List returns a list of all tasks in the repository.
func (r *TaskRepository) List() ([]domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// If there are no tasks in the repository, return an empty list.
	if len(r.tasks) == 0 {
		return []domain.Task{}, nil
	}

	// Create a new list to store the tasks.
	var taskList []domain.Task

	// Iterate over the tasks in the repository and append them to the list.
	for _, task := range r.tasks {
		taskList = append(taskList, task)
	}

	// Return the list of tasks.
	return taskList, nil
}

// Update updates a task in the repository.
// It returns the updated task if it was found in the repository.
// If the task was not found, it returns an error.
func (r *TaskRepository) Update(id int, req domain.Task) (*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Get the task from the repository
	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.ErrTaskNotFound
	}

	// Update the task
	task.Title = req.Title
	task.Description = req.Description
	task.Completed = req.Completed

	// Store the updated task in the repository
	r.tasks[id] = task

	// Return the updated task
	return &task, nil
}

// Delete deletes a task from the repository.
// It returns an error if the task was not found.
func (r *TaskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Check if the task exists in the repository
	if _, exists := r.tasks[id]; exists {
		// Delete the task from the repository
		delete(r.tasks, id)
		return nil
	}
	// Return an error if the task was not found
	return errors.ErrTaskNotFound
}

// Exists checks if a task with the given ID exists in the repository.
func (r *TaskRepository) Exists(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	// Check if the task exists in the repository
	_, exists := r.tasks[id]
	return exists
}

// GetByID retrieves a task by its ID.
// It returns the task if found, otherwise returns an error.
func (r *TaskRepository) GetByID(id int) (*domain.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Get the task from the repository
	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.ErrTaskNotFound
	}

	// Return the task
	return &task, nil
}
