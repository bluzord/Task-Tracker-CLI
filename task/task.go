package task

import (
	"fmt"
	"time"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in-progress"
	StatusDone       Status = "done"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusTodo, StatusInProgress, StatusDone:
		return true
	}
	return false
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(id int, description string) (*Task, error) {

	if id < 1 {
		return nil, fmt.Errorf("task: invalid id (must be greater than 0)")
	}

	if description == "" {
		return nil, fmt.Errorf("task: description is required")
	}

	if len(description) > 50 {
		return nil, fmt.Errorf("task: description is too long (must be less or equal to 50)")
	}

	return &Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}
