package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID            string
	Title         string
	Completed     bool
	CompletedDate string
}

type Storage struct {
	filename string
}

func NewStorage(filename string) *Storage {
	return &Storage{
		filename: filename,
	}
}

func (s *Storage) ReadTasks() ([]Task, error) {
	if _, err := os.Stat(s.filename); os.IsNotExist(err) {
		file, err := os.Create(s.filename)
		if err != nil {
			return nil, fmt.Errorf("failed to create CSV file: %v", err)
		}
		file.Close()
		return []Task{}, nil
	}

	file, err := os.Open(s.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %v", err)
	}

	var tasks []Task
	for _, record := range records {
		if len(record) < 3 {
			continue
		}
		completed, _ := strconv.ParseBool(record[2])
		task := Task{
			ID:            record[0],
			Title:         record[1],
			Completed:     completed,
			CompletedDate: "",
		}
		if len(record) > 3 {
			task.CompletedDate = record[3]
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (s *Storage) WriteTasks(tasks []Task) error {
	file, err := os.Create(s.filename)
	if err != nil {
		return fmt.Errorf("failed to create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, task := range tasks {
		record := []string{
			task.ID,
			task.Title,
			strconv.FormatBool(task.Completed),
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write CSV record: %v", err)
		}
	}
	return nil
}
