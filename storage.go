package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Task struct {
	ID        int
	Title     string
	Done      bool
	Timestamp string
	note      string
}

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	return &Storage{filename: filename}
}

func (s *Storage) ReadTasks() ([]Task, error) {
	file, err := os.Open(s.filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		done, _ := strconv.ParseBool(record[2])
		task := Task{
			ID:        id,
			Title:     record[1],
			Done:      done,
			Timestamp: record[3],
			note:      record[4],
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
