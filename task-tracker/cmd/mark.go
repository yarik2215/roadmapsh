package cmd

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yarik2215/roadmapsh/task-tracker/tracker"
)

func newMarkCmd(t *tracker.TaskTracker) *cobra.Command {
	cmd := cobra.Command{
		Use:   "mark",
		Short: "Change task status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid task id: %s", args[1])
			}
			status := tracker.TaskStatus(args[0])
			if !slices.Contains(tracker.TASK_STATUSES[:], status) {
				return fmt.Errorf("invalid status %s. Valid statuses are: %s", status, tracker.TASK_STATUSES)
			}
			task, err := t.ChangeStatus(id, status)
			if err != nil {
				return err
			}
			tracker.PrintTask(task)
			return nil
		},
	}
	return &cmd
}
