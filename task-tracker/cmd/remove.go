package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yarik2215/roadmapsh/task-tracker/tracker"
)

func newRemoveCmd(t *tracker.TaskTracker) *cobra.Command {
	cmd := cobra.Command{
		Use:   "remove",
		Short: "Remove task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid task id: %s", args[0])
			}
			err = t.Remove(id)
			if err != nil {
				return err
			}
			fmt.Printf("Task %d removed\n", id)
			return nil
		},
	}
	return &cmd
}
