package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yarik2215/roadmapsh/task-tracker/tracker"
)

func newRootCmd() *cobra.Command {
	t := tracker.NewTracker(tracker.NewFileStorage("./tasks.json"))
	root := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI tool for managing tasks",
	}

	root.AddCommand(
		newAddCmd(&t),
		newGetCmd(&t),
		newListCmd(&t),
		newUpdateCmd(&t),
		newRemoveCmd(&t),
		newMarkCmd(&t),
	)

	return root
}

func Execute() {
	cmd := newRootCmd()
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
