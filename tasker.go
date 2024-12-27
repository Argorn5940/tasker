package main

import (
	"fmt"
	"os"
	"time"

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

func (tm *TaskManager) CompleteTask(taskID string) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == taskID {
			tm.tasks[i].Completed = true
			tm.tasks[i].CompletedDate = time.Now().Format("2006-01-02-15-04")
			return nil
		}
	}
	return fmt.Errorf("タスクが見つかりません:%s", taskID)
}
