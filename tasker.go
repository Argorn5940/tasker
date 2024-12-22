package main

import (
	"os"
	"strconv"

	"github.com/aquasecurity/table"
)

type TaskManger struct {
	tasks []Task
}

func NewTaskManger() *TaskManger {
	return &TaskManger{
		tasks: make([]Task, 0),
	}
}

func (tm *TaskManger) DisplyTasks() error {
	t := table.New(os.Stdout)
	t.SetHeaders("ID", "タスク", "状態", "作成日時", "メモ")

	for _, task := range tm.tasks {
		status := "✖"
		if task.Done {
			status = "✔"
		}
		t.AddRow(
			strconv.Itoa(task.ID),
			task.Title,
			status,
			task.Timestamp,
			task.note,
		)
	}
	t.Render()
	return nil
}
