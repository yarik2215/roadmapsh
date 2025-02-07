package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/yarik2215/roadmapsh/task-tracker/tracker"
)

func newUpdateCmd(t *tracker.TaskTracker) *cobra.Command {
	cmd := cobra.Command{
		Use:   "mark",
		Short: "Change task status",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid task id: %s", args[1])
			}
			description := args[0]
			task, err := t.Update(id, description)
			if err != nil {
				return err
			}
			tracker.PrintTask(task)
			return nil
		},
	}
	return &cmd
}
