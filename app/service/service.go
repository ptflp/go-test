package service

import (
	"context"
	"github.com/ptflp/go-test/app/entity"
	"github.com/ptflp/go-test/app/repository"
)

type Todoer interface {
	Create(ctx context.Context, todo *entity.Todo) error
	Complete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*entity.Todo, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, todo *entity.Todo) error
}

type TodoService struct {
	storage repository.TodoStorager
}

// NewTodoService returns a new instance of TodoService.
func NewTodoService(storage repository.TodoStorager) Todoer {
	return &TodoService{
		storage: storage,
	}
}

// Create creates a new todo.
func (s *TodoService) Create(ctx context.Context, todo *entity.Todo) error {
	return s.storage.Create(ctx, todo)
}

// Complete marks a todo as complete.
func (s *TodoService) Complete(ctx context.Context, id int) error {
	todo, err := s.storage.GetByID(ctx, id)
	if err != nil {
		return err
	}
	todo.IsCompleted = !todo.IsCompleted

	return s.storage.Update(ctx, todo)
}

// GetAll returns all todos.
func (s *TodoService) GetAll(ctx context.Context) ([]*entity.Todo, error) {
	return s.storage.GetAll(ctx)
}

// Delete deletes a todo.
func (s *TodoService) Delete(ctx context.Context, id int) error {
	return s.storage.Delete(ctx, id)
}

// Update updates a todo.
func (s *TodoService) Update(ctx context.Context, todo *entity.Todo) error {
	todoPrev, err := s.storage.GetByID(ctx, int(todo.ID))
	if err != nil {
		return err
	}
	todoPrev.Title = todo.Title

	return s.storage.Update(ctx, todoPrev)
}
