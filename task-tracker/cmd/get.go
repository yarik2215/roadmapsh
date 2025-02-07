package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yarik2215/roadmapsh/task-tracker/tracker"
)

func newGetCmd(t *tracker.TaskTracker) *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "Get task",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid task id: %s", args[0])
			}
			task, err := t.Get(id)
			if err != nil {
				return err
			}
			tracker.PrintTask(task)
			return nil
		},
	}
	return &cmd
}
