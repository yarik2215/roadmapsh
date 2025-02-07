package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yarik2215/roadmapsh/task-tracker/tracker"
)

func newAddCmd(t *tracker.TaskTracker) *cobra.Command {
	cmd := cobra.Command{
		Use:   "add",
		Short: "Add a new task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			description := args[0]
			tracker.PrintTask(t.Add(description, tracker.TASK_STATUS_TODO))
			return nil
		},
	}
	return &cmd
}
