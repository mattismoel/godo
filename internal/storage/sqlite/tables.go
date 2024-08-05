package sqlite

import (
	"database/sql"
	"fmt"
)

func (s sqliteStorage) createTodosTable(tx *sql.Tx) error {
	query := `
  CREATE TABLE IF NOT EXISTS todos (
    id                INTEGER     PRIMARY KEY     NOT NULL,
    text              TEXT                        NOT NULL,
    done              BOOLEAN                     NOT NULL,
    created_at        INTEGER                     NOT NULL,
    updated_at        INTEGER                     NOT NULL,
    completed_at      INTEGER     DEFAULT(0)      NOT NULL
  )`

	_, err := tx.Exec(query)
	if err != nil {
		return fmt.Errorf("could not execute query: %v", err)
	}

	return nil
}
