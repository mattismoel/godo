/*
Copyright © 2024 Mattis Møl Kristensen <mattismoel@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mattismoel/gotodo/internal/model"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists all todos in the todo list.",
	Long: `Lists all todos in the todo list. Can be used to get an overview over
  the todos in the todo list, or for gaining information about the IDs of the
  different todos. For example:

  godo ls`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		todos, err := todoSrv.All(ctx)
		if err != nil {
			log.Fatalf("could not get todos from storage: %v", err)
		}

		if len(todos) <= 0 {
			fmt.Printf("It's quite empty here...\nAdd a todo with the 'add' command.\n")
			return
		}

		completed := []model.Todo{}
		notCompleted := []model.Todo{}

		for _, todo := range todos {
			if todo.Done {
				completed = append(completed, todo)
			} else {
				notCompleted = append(notCompleted, todo)
			}
		}

		fmt.Println("Not Completed:")
		fmt.Println("--------------")
		if len(notCompleted) > 0 {
			for _, todo := range notCompleted {
				fmt.Printf("[#%d] %s\n", todo.ID, todo.Text)
			}
		} else {
			fmt.Println("Nothing to see here...")
		}

		fmt.Println()
		fmt.Println()

		fmt.Println("Completed:")
		fmt.Println("----------")
		if len(completed) > 0 {
			for _, todo := range completed {
				fmt.Printf("\x1b[9m[#%d] %s\x1b[0m\n", todo.ID, todo.Text)
			}
		} else {
			fmt.Println("Nothing to see here...")
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
