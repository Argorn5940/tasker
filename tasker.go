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

func (tm *TaskManager) AddTask(title string) error {
	//新しいタスクIDを生成（既存タスクの数+1)
	newID := fmt.Sprintf("%d", len(tm.tasks)+1)

	//新しいタスクを作成
	task := Task{
		ID:            newID,
		Title:         title,
		Completed:     false,
		CompletedDate: "",
	}
	//タスクをリストに追加
	tm.tasks = append(tm.tasks, task)
	return nil
}

func (tm *TaskManager) UpdateTask(taskID string, newTitle string) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == taskID {
			tm.tasks[i].Title = newTitle
			return nil
		}
	}
	return fmt.Errorf("タスクが見つかりません: %s", taskID)
}

func (tm *TaskManager) DeleteTask(taskID string) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == taskID {
			//タスクを削除（スライスから要素を削除)
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("タスクが見つかりません: %s", taskID)
}
