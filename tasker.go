package main

import (
	"os"

	"github.com/aquasecurity/table"
)

type TaskManager struct {
	tasks []Task
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make([]Task, 0),
	}
}

func (tm *TaskManager) DisplayTasks() error {
	t := table.New(os.Stdout)
	t.SetHeaders("ID", "タスク", "状態", "完了日")

	for _, task := range tm.tasks {
		status := "✖"
		if task.Completed {
			status = "✔"
		}
		t.AddRow(
			task.ID,
			task.Title,
			status,
			task.CompletedDate,
		)
	}
	t.Render()
	return nil
}
