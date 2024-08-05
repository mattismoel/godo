/*
Copyright © 2024 Mattis Møl Kristensen <mattismoel@gmail.com>
*/
package cmd

import (
	"context"
	"log"
	"time"

	"github.com/mattismoel/gotodo/internal/model"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds one or multiple todos to the todo list.",
	Long: `Adds one or multiple todos to the todo list. For example:

  godo add "First todo"
  godo add "Second todo" "Third todo" "Fourth todo"

  The above first adds a todo "First todo", afterwhich three more todos are added,
  "Second todo", "Third todo" and "Fourth todo"`,

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		actionCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for _, arg := range args {
			todo, err := model.NewTodo(arg)
			if err != nil {
				log.Fatalf("could not create todo: %v", err)
			}

			_, err = todoSrv.Add(actionCtx, todo)
			if err != nil {
				log.Fatalf("could not add todo: %v", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
