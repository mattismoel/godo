package model

import (
	"fmt"
	"strings"
	"time"
)

type Todo struct {
	ID          int64
	Text        string
	Done        bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt time.Time
}

func NewTodo(text string) (Todo, error) {
	var todo Todo
	if strings.TrimSpace(text) == "" {
		return Todo{}, fmt.Errorf("no text provided for todo...")
	}

	todo.Text = text
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	return todo, nil
}
