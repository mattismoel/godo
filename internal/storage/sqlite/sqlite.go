package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/mattismoel/gotodo/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteStorage struct {
	db *sql.DB
}

func New(ctx context.Context, dbPath string) (*sqliteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("could not open sqlite database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	store := &sqliteStorage{db: db}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %v", err)
	}

	defer tx.Rollback()

	err = store.createTodosTable(tx)
	if err != nil {
		return nil, fmt.Errorf("could not create todo table: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("could not commit transaction: %v", err)
	}

	return store, nil
}

func (s sqliteStorage) AllTodos(ctx context.Context) ([]model.Todo, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("could not begin transaction: %v", err)
	}

	defer tx.Rollback()

	query := `
  SELECT
    id,
    text,
    done,
    created_at,
    updated_at,
    completed_at
  FROM todos`

	rows, err := tx.Query(query)
	if err != nil {
		return nil, fmt.Errorf("could not query for todos: %v", err)
	}

	defer rows.Close()

	todos := []model.Todo{}
	for rows.Next() {
		var todo model.Todo
		var createdAtUnix, updatedAtUnix, completedAtUnix int64

		err = rows.Scan(&todo.ID, &todo.Text, &todo.Done, &createdAtUnix, &updatedAtUnix, &completedAtUnix)
		if err != nil {
			return nil, fmt.Errorf("could not scan todo into struct: %v", err)
		}

		todo.CreatedAt = time.Unix(createdAtUnix, 0)
		todo.UpdatedAt = time.Unix(updatedAtUnix, 0)
		todo.CompletedAt = time.Unix(completedAtUnix, 0)

		todos = append(todos, todo)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("could not commit transaction: %v", err)
	}

	return todos, nil
}

func (s sqliteStorage) TodoByID(ctx context.Context, id int64) (model.Todo, error) {
	return model.Todo{}, nil
}

func (s sqliteStorage) AddTodo(ctx context.Context, todo model.Todo) (int64, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("could not begin transaction: %v", err)
	}

	defer tx.Rollback()
	query := `
  INSERT INTO todos (
    text,
    done,
    created_at,
    updated_at
  ) VALUES (?, ?, ?, ?)`

	res, err := tx.Exec(query,
		todo.Text,
		false,
		time.Now().Unix(),
		time.Now().Unix(),
	)

	if err != nil {
		return 0, fmt.Errorf("could not exec query: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("could not get id of inserted todo: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("could not commit transaction: %v", err)
	}

	return id, nil
}

func (s sqliteStorage) RemoveTodoByID(ctx context.Context, id int64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	defer tx.Rollback()

	query := "DELETE FROM todos WHERE id = ?"
	_, err = tx.Exec(query, id)
	if err != nil {
		return fmt.Errorf("could not execute query: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}

func (s sqliteStorage) ToggleDoneTodoByID(ctx context.Context, id int64, done bool) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}

	defer tx.Rollback()

	query := `
  UPDATE todos
  SET
    done = ?,
    completed_at = ?
  WHERE id = ?`

	_, err = tx.Exec(query, done, time.Now().Unix(), id)
	if err != nil {
		return fmt.Errorf("could not execute query: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("could not commit transaction: %v", err)
	}

	return nil
}

func (s sqliteStorage) IsTodoDone(ctx context.Context, id int64) (bool, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return false, fmt.Errorf("could not begin transaction: %v", err)
	}

	defer tx.Rollback()

	query := "SELECT done FROM todos WHERE id = ?"

	var done bool
	err = tx.QueryRow(query, id).Scan(&done)
	if err != nil {
		return false, fmt.Errorf("could not scan 'done' status into bool: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return false, fmt.Errorf("could not commit transrction: %v", err)
	}

	return done, nil
}
