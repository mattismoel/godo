package memory

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mattismoel/gotodo/internal/model"
)

var todos = []model.Todo{
	{
		ID:        0,
		Text:      "Go grocery shopping",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        1,
		Text:      "Exercise",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type memoryStorage struct {
	todos []model.Todo
}

func New() *memoryStorage {
	return &memoryStorage{todos: todos}
}

func (s memoryStorage) AllTodos(ctx context.Context) ([]model.Todo, error) {
	return s.todos, nil
}

func (s memoryStorage) TodoByID(ctx context.Context, id int64) (model.Todo, error) {
	for _, todo := range s.todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return model.Todo{}, fmt.Errorf("could not find todo with id %d", id)
}

func (s *memoryStorage) AddTodo(ctx context.Context, todo model.Todo) (int64, error) {
	id := int64(rand.Int())
	todo.ID = id
	s.todos = append(s.todos, todo)
	return id, nil
}
