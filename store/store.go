package store

import (
	"fmt"
	"strings"
	"task-tracker/task"
)

type Store interface {
	AddTask(description string) error
	GetTask(id int) (*task.Task, error)
	UpdateTask(id int, description string) error
	DeleteTask(id int) (*task.Task, error)
	LastID() (int, error)
	ListAllTasks() ([]*task.Task, error)
	ListTasksByStatus(status task.Status) ([]*task.Task, error)
	loadTasks() error
	saveTasks() error
}

type JSONStore struct {
	path  string
	tasks []task.Task
}

func (j JSONStore) AddTask(description string) error {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) GetTask(id int) (*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) UpdateTask(id int, description string) error {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) DeleteTask(id int) (*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) LastID() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) ListAllTasks() ([]*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) ListTasksByStatus(status task.Status) ([]*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) loadTasks() error {
	//TODO implement me
	panic("implement me")
}

func (j JSONStore) saveTasks() error {
	//TODO implement me
	panic("implement me")
}

func NewJSONStore(path string) (*JSONStore, error) {

	if !strings.HasSuffix(path, ".json") {
		return nil, fmt.Errorf("store: path does not end with .json")
	}

	return &JSONStore{path: path}, nil
}
