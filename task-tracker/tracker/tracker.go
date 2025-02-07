package tracker

import (
	"fmt"
	"slices"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

var TASK_STATUSES = [...]TaskStatus{TASK_STATUS_TODO, TASK_STATUS_DONE, TASK_STATUS_IN_PROGRESS}

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskStorage interface {
	Load(dst *[]Task)
	Save(src []Task)
}

type TaskTracker struct {
	tasks   []Task
	storage TaskStorage
}

func NewTracker(storage TaskStorage) TaskTracker {
	return TaskTracker{tasks: make([]Task, 0), storage: storage}
}

func (t *TaskTracker) load() {
	// copy(t.tasks, t.storage.Load())
	t.storage.Load(&t.tasks)
}

func (t *TaskTracker) save() {
	t.storage.Save(t.tasks)
}

func (t *TaskTracker) Add(description string, status TaskStatus) Task {
	t.load()
	id := 0
	if len(t.tasks) > 0 {
		id = t.tasks[len(t.tasks)-1].Id
	}
	task := Task{
		Id: id + 1, Description: description, Status: status, CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}
	t.tasks = append(t.tasks, task)
	t.storage.Save(t.tasks)
	return task
}

func (t *TaskTracker) Remove(id int) error {
	i, ok := slices.BinarySearchFunc(t.tasks, id, func(e Task, t int) int { return t - e.Id })
	if !ok {
		return fmt.Errorf("Task #%d not found", id)
	}
	t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
	t.save()
	return nil
}

func (t *TaskTracker) ChangeStatus(id int, status TaskStatus) (Task, error) {
	t.load()
	i, ok := slices.BinarySearchFunc(t.tasks, id, func(e Task, t int) int { return t - e.Id })
	if !ok {
		return Task{}, fmt.Errorf("Task #%d not found", id)
	}
	t.tasks[i].Status = status
	t.save()
	return t.tasks[i], nil
}

func (t *TaskTracker) Update(id int, description string) (Task, error) {
	t.load()
	i, ok := slices.BinarySearchFunc(t.tasks, id, func(e Task, t int) int { return t - e.Id })
	if !ok {
		return Task{}, fmt.Errorf("Task #%d not found", id)
	}
	t.tasks[i].Description = description
	t.save()
	return t.tasks[i], nil
}

func (t *TaskTracker) Get(id int) (Task, error) {
	t.load()
	i, ok := slices.BinarySearchFunc(t.tasks, id, func(e Task, t int) int { return t - e.Id })
	if !ok {
		return Task{}, fmt.Errorf("Task #%d not found", id)
	}
	return t.tasks[i], nil
}

func (t *TaskTracker) List() []Task {
	t.load()
	return t.tasks
}

func (t *TaskTracker) Filter(status TaskStatus) []Task {
	t.load()
	arr := make([]Task, 0)
	for _, task := range t.tasks {
		if task.Status == status {
			arr = append(arr, task)
		}
	}
	return arr
}
