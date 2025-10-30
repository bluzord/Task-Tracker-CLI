package task

import (
	"errors"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

func (s Status) isValid() bool {
	switch s {
	case StatusTodo, StatusInProgress, StatusDone:
		return true
	}
	return false
}

type Task struct {
	id          int
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(id int, description string) (*Task, error) {

	if description == "" {
		return nil, errors.New("description is required")
	}

	return &Task{
		id:          id,
		description: description,
		status:      StatusTodo,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}, nil
}

func (t *Task) Status() Status {
	return t.status
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}
