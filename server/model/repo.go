package model

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

// TasksList : Return all existing tasks
func (r *Repo) TasksList() Tasks {
	return r.tasks
}

// InitialRepo : Give us some seed data
func InitialRepo() *Repo {
	r := &Repo{}
	r.CreateTask(Task{Name: "Import bot loan interest"})
	r.CreateTask(Task{Name: "Host meetup"})
	return r
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
