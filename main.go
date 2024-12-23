package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	// 実行ファイルのディレクトリを取得
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// CSVファイルのパスを実行ファイルと同じディレクトリに設定
	csvPath := filepath.Join(dir, "tasks.csv")
	storage := NewStorage(csvPath)
	tm := NewTaskManager()

	// CSVファイルからタスクを読み込む
	tasks, err := storage.ReadTasks()
	if err != nil {
		log.Printf("Error reading tasks from file %s: %v", csvPath, err)
		tasks = []Task{}
	}
	tm.tasks = tasks

	if len(os.Args) < 2 {
		log.Fatal("コマンドを指定してください（例：list）")
		return
	}

	switch os.Args[1] {
	case "list":
		if err := tm.DisplayTasks(); err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatal("無効なコマンドです。'list'を使用してください")
	}

	if err := storage.WriteTasks(tm.tasks); err != nil {
		log.Printf("Warning: Failed to save tasks: %v", err)
	}
}