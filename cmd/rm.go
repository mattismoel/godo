/*
Copyright © 2024 Mattis Møl Kristensen <mattismoel@gmail.com>
*/
package cmd

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Removes one or multiple todos from the todo list.",
	Long: `Removes one or multiple todos from the todo list. For example:

  godo rm 1 4 2 5

  The above removes todos with ID 1, 4, 2 and 5. The IDs can be found with the
  'ls' command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatalf("could not read id as int: %v", err)
			}

			err = todoSrv.Remove(ctx, int64(id))
			if err != nil {
				log.Fatalf("could not remove todo: %v", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
