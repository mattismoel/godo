package storage

import (
	"context"

	"github.com/mattismoel/gotodo/internal/model"
)

type Storage interface {
	AllTodos(context.Context) ([]model.Todo, error)
	TodoByID(context.Context, int64) (model.Todo, error)
	AddTodo(context.Context, model.Todo) (int64, error)
	RemoveTodoByID(context.Context, int64) error
	ToggleDoneTodoByID(context.Context, int64, bool) error
	IsTodoDone(context.Context, int64) (bool, error)
}
