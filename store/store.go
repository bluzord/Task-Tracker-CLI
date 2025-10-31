package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"task-tracker/task"
)

type Store interface {
	// AddTask adds a new task with the given description and saves changes to storage.
	AddTask(description string) error

	// GetTask returns a task by its ID.
	// Returns an error if the task does not exist.
	GetTask(id int) (*task.Task, error)

	// UpdateTask updates the description of the task with the given ID
	// and saves the updated data to storage.
	UpdateTask(id int, description string) error

	// DeleteTask removes a task by its ID and returns the deleted task.
	// Returns an error if the task does not exist.
	DeleteTask(id int) (*task.Task, error)

	// LastID returns the highest task ID in the storage.
	// Used to generate new unique IDs.
	LastID() (int, error)

	// ListAllTasks returns all tasks stored in the system.
	ListAllTasks() ([]*task.Task, error)

	// ListTasksByStatus returns all tasks filtered by the given status (todo, in-progress, done).
	ListTasksByStatus(status task.Status) ([]*task.Task, error)

	// loadTasks reads tasks from the underlying storage (e.g., JSON file)
	// and loads them into memory.
	loadTasks() error

	// saveTasks writes the current in-memory list of tasks to the underlying storage.
	saveTasks() error
}

type JSONStore struct {
	path  string
	tasks []task.Task
}

func (j *JSONStore) AddTask(description string) error {
	//TODO implement me
	panic("implement me")
}

func (j *JSONStore) GetTask(id int) (*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (j *JSONStore) UpdateTask(id int, description string) error {
	//TODO implement me
	panic("implement me")
}

func (j *JSONStore) DeleteTask(id int) (*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (j *JSONStore) LastID() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (j *JSONStore) ListAllTasks() ([]*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (j *JSONStore) ListTasksByStatus(status task.Status) ([]*task.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (j *JSONStore) loadTasks() error {
	data, err := os.ReadFile(j.path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &j.tasks)
}

func (j *JSONStore) saveTasks() error {
	//TODO implement me
	panic("implement me")
}

func NewJSONStore(path string) (*JSONStore, error) {

	if !strings.HasSuffix(path, ".json") {
		return nil, fmt.Errorf("store: path does not end with .json")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.WriteFile(path, []byte("[]"), 0644)
		if err != nil {
			return nil, err
		}
	}

	j := &JSONStore{path: path}
	err := j.loadTasks()
	if err != nil {
		return nil, err
	}

	return j, nil
}
