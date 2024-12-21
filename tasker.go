package main

import "time"

type Task struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Tasker []Task

func (tasker *Tasker) Add(title string) {
	task := Task{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	*tasker = append(*tasker, task)
}
