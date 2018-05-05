package model

import (
	"fmt"
)

// Repo : repo for Tasks
type Repo struct {
	currentID int
	tasks     Tasks
}

// CreateTask : add new task to DB
func (r *Repo) CreateTask(t Task) Task {
	r.currentID++
	t.ID = r.currentID
	r.tasks = append(r.tasks, t)
	return t
}

// InitialRepo : Give us some seed data
func InitialRepo() *Repo {
	r := &Repo{}
	r.CreateTask(Task{Name: "Import bot loan interest"})
	r.CreateTask(Task{Name: "Host meetup"})
	return r
}

// TasksList : Return all existing tasks
func (r *Repo) TasksList() Tasks {
	return r.tasks
}

// TaskByID : Task by ID
func (r *Repo) TaskByID(id int) Task {
	for _, t := range r.tasks {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return Task{}
}

// UpdateTask : Update task by id
func (r *Repo) UpdateTask(task Task) Task {
	found := false
	for i, t := range r.tasks {
		if t.ID == task.ID {
			r.tasks[i] = task
			found = true
		}
	}
	if !found {
		err := fmt.Errorf("Not found id:%v", task.ID)
		panic(err)
	}
	return task
}
