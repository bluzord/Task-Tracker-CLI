package task

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
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
	ID          int       `json:"id" validate:"required,min=1"`
	Description string    `json:"description" validate:"required,min=1,max=50"`
	Status      Status    `json:"status" validate:"required"`
	CreatedAt   time.Time `json:"created_at" validate:"required"`
	UpdatedAt   time.Time `json:"updated_at" validate:"required"`
}

var validate = validator.New()

func ValidateTask(t *Task) error {
	return validate.Struct(t)
}

func NewTask(id int, description string) (*Task, error) {

	creationTime := time.Now()

	t := &Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   creationTime,
		UpdatedAt:   creationTime,
	}

	if err := ValidateTask(t); err != nil {
		return nil, fmt.Errorf("NewTask: invalid task: %w", err)
	}

	return t, nil
}
