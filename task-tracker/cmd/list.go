package cmd

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
	"github.com/yarik2215/roadmapsh/task-tracker/tracker"
)

func newListCmd(t *tracker.TaskTracker) *cobra.Command {
	cmd := cobra.Command{
		Use:   "list",
		Short: "List tasks",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				status := tracker.TaskStatus(args[0])
				if !slices.Contains(tracker.TASK_STATUSES[:], status) {
					return fmt.Errorf("invalid status %s. Valid statuses are: %s", status, tracker.TASK_STATUSES)
				}
				tracker.PrintTasks(t.Filter(status))
			} else {
				tracker.PrintTasks(t.List())
			}
			return nil
		},
	}
	return &cmd
}
