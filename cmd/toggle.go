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

var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggles the 'done' state of one or multiple todos",
	Long: `Toggles the 'done' state of one or multiple todos. For example:

  godo toggle 1 2 4 6

  The above toggles the 'done' state of the todos with IDs 1, 2, 4 and 6.
  The IDs can be found with the 'ls' command.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatalf("could not read id as int: %v", err)
			}

			err = todoSrv.ToggleDone(ctx, int64(id))
			if err != nil {
				log.Fatalf("could not toggle done: %v", err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}
