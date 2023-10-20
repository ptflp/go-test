package repository

import (
	"context"
	"fmt"
	"github.com/ptflp/go-test/app/entity"
	"sync"
)

type InMemory struct {
	unimplemented
	data []*entity.Todo
	sync.Mutex
}

func NewInMemory() TodoStorager {
	return &InMemory{
		data: make([]*entity.Todo, 0, 100),
	}
}

func (r *InMemory) Create(ctx context.Context, todo *entity.Todo) error {
	r.Lock()
	defer r.Unlock()
	todo.ID = int64(len(r.data) + 1)
	r.data = append(r.data, todo)
	if len(r.data) > 99 {
		r.data = r.data[:3]
	}

	return nil
}

func (r *InMemory) GetAll(ctx context.Context) ([]*entity.Todo, error) {
	r.Lock()
	defer r.Unlock()

	return r.data, nil
}

func (r *InMemory) Delete(ctx context.Context, id int) error {
	r.Lock()
	defer r.Unlock()
	for i, todo := range r.data {
		if todo.ID == int64(id) {
			r.data = append(r.data[:i], r.data[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("todo not found")
}

func (r *InMemory) Update(ctx context.Context, todo *entity.Todo) error {
	r.Lock()
	defer r.Unlock()
	for i, t := range r.data {
		if t.ID == todo.ID {
			r.data[i] = todo
			return nil
		}
	}

	return fmt.Errorf("todo not found")
}

// GetByID returns a todo by id.
func (r *InMemory) GetByID(ctx context.Context, id int) (*entity.Todo, error) {
	r.Lock()
	defer r.Unlock()
	for _, todo := range r.data {
		if todo.ID == int64(id) {
			return todo, nil
		}
	}

	return nil, fmt.Errorf("todo not found")
}
