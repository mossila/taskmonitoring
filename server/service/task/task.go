package task

import (
	"time"
)

// Task -
type Task struct {
	Name   string
	Status Status
	Due    time.Time
}

// Status -
type Status int

const (
	Fail Status = iota + 1
	Success
	Waiting
)

func (status Status) String() string {
	name := [...]string{
		"Fail",
		"Success",
		"Waiting",
	}
	return name[status]
}
