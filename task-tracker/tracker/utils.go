package tracker

import "fmt"

func PrintTask(task Task) {
	fmt.Printf(" * %d %s %s | %s", task.Id, task.Description, task.Status, task.UpdatedAt)
}

func PrintTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf(" * %d %s %s | %s\n", t.Id, t.Description, t.Status, t.UpdatedAt.Format("2006-01-02 15:04"))
	}
}
