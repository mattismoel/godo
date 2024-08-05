/*
Copyright © 2024 Mattis Møl Kristensen <mattismoel@gmail.com>
*/
package cmd

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mattismoel/gotodo/internal/service"
	"github.com/mattismoel/gotodo/internal/storage/sqlite"
	"github.com/spf13/cobra"
)

var (
	todoSrv *service.TodoService
)

var rootCmd = &cobra.Command{
	Use:   "gotodo",
	Short: "A simple todo list CLI application.",
}

func Execute() {
	sqliteCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sqliteStore, err := sqlite.New(sqliteCtx, "database.db")
	if err != nil {
		log.Fatalf("could not create sqlite instance: %v", err)
	}

	todoSrv = service.NewTodoService(sqliteStore)

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
