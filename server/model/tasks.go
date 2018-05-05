package tasks

import "time"

// Task : Task item
type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Tasks : List of Task
type Tasks []Task
