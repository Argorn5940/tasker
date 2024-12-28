package main

import "fmt"

//ListTasksCmd、タスク一覧を表示するコマンドを処理
func ListTasksCmd(tm *TaskManager) error {
	//DisplayTaskeメソッドを呼び出してタスク一覧を表示
	if err := tm.DisplayTasks(); err != nil {
		return fmt.Errorf("タスク一覧の表示に失敗しました: %v", err)
	}
	return nil
}

func CompleteTaskCmd(tm *TaskManager, taskID string) error {
	if err := tm.CompleteTask(taskID); err != nil {
		return fmt.Errorf("タスクの完了処理に失敗しました: %v", err)
	}
	return nil
}

func AddTaskCmd(tm *TaskManager, title string) error {
	if err := tm.AddTask(title); err != nil {
		return fmt.Errorf("タスクの追加に失敗しました: %v", err)
	}
	fmt.Printf("新しいタスクを追加しました: %s\n", title)
	return nil
}

//HandleCommand コマンドの処理
func HandleCommand(command string, tm *TaskManager, args []string) error {
	switch command {
	case "-list":
		return ListTasksCmd(tm)
	case "-comp":
		if len(args) < 1 {
			return fmt.Errorf("タスクIDを指定してください")
		}
		taskID := args[0]
		return CompleteTaskCmd(tm, taskID)
	case "-add":
		if len(args) < 1 {
			return fmt.Errorf("タスクの説明を入力してください")
		}
		return AddTaskCmd(tm, args[0])
	case "-del":
		if len(args) < 1 {
			return fmt.Errorf("削除するタスクIDを指定してください")
		}
		return DeleteTaskCmd(tm, args[0])
	default:
		return fmt.Errorf("無効なコマンドです: %s", command)
	}
}
