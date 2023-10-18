package repository

import (
	"context"
	"github.com/ptflp/go-test/app/entity"
)

type unimplemented struct{}

func (unimplemented) Create(ctx context.Context, todo *entity.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (unimplemented) GetAll(ctx context.Context) ([]*entity.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (unimplemented) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}

func (unimplemented) Update(ctx context.Context, todo *entity.Todo) error {
	//TODO implement me
	panic("implement me")
}

func (unimplemented) GetByID(ctx context.Context, id int) (*entity.Todo, error) {
	//TODO implement me
	panic("implement me")
}
