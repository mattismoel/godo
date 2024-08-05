package service

import (
	"context"
	"fmt"

	"github.com/mattismoel/gotodo/internal/model"
	"github.com/mattismoel/gotodo/internal/storage"
)

type TodoService struct {
	store storage.Storage
}

func NewTodoService(store storage.Storage) *TodoService {
	return &TodoService{store: store}
}

func (s TodoService) All(ctx context.Context) ([]model.Todo, error) {
	todos, err := s.store.AllTodos(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get all todos from storage: %v", err)
	}

	return todos, nil
}

func (s TodoService) ByID(ctx context.Context, id int64) (model.Todo, error) {
	todo, err := s.store.TodoByID(ctx, id)
	if err != nil {
		return model.Todo{}, fmt.Errorf("could not get todo with id %d from storage: %v", id, err)
	}

	return todo, nil
}

func (s TodoService) Add(ctx context.Context, todo model.Todo) (int64, error) {
	id, err := s.store.AddTodo(ctx, todo)
	if err != nil {
		return 0, fmt.Errorf("could not add todo to storage: %v", err)
	}

	return id, nil
}

func (s TodoService) Remove(ctx context.Context, id int64) error {
	err := s.store.RemoveTodoByID(ctx, id)
	if err != nil {
		return fmt.Errorf("could not remove todo with id %d from storage: %v", id, err)
	}

	return nil
}

func (s TodoService) ToggleDone(ctx context.Context, id int64) error {
	isAlreadyDone, err := s.store.IsTodoDone(ctx, id)
	if err != nil {
		return fmt.Errorf("could not get 'done' status for todo with id %d: %v", id, err)
	}

	err = s.store.ToggleDoneTodoByID(ctx, id, !isAlreadyDone)
	if err != nil {
		return fmt.Errorf("could not complete todo in store: %v", err)
	}

	return nil
}
