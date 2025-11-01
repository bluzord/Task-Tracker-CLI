package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"tracker/task"
)

type Store interface {
	// AddTask adds a new task with the given description and saves changes to storage.
	// Returns an ID of a new task.
	AddTask(description string) (int, error)

	// UpdateTask updates the description of the task with the given ID
	// and saves the updated data to storage.
	UpdateTask(id int, description string) error

	// DeleteTask removes a task by its ID and returns the deleted task.
	// Returns an error if the task does not exist.
	DeleteTask(id int) (*task.Task, error)

	// LastID returns the highest task ID in the storage.
	// Used to generate new unique IDs.
	LastID() int

	// ListAllTasks returns all tasks stored in the system.
	ListAllTasks() []task.Task

	// ListTasksByStatus returns all tasks filtered by the given status (todo, in-progress, done).
	ListTasksByStatus(status task.Status) []task.Task

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

var statusOrder = map[task.Status]int{
	task.StatusTodo:       1,
	task.StatusInProgress: 2,
	task.StatusDone:       3,
}

func (j *JSONStore) AddTask(description string) (int, error) {
	id := j.LastID() + 1

	t, err := task.NewTask(id, description)
	if err != nil {
		return -1, err
	}

	j.tasks = append(j.tasks, *t)

	err = j.saveTasks()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func (j *JSONStore) UpdateTask(id int, description string) error {
	if len(description) < 1 || len(description) > 50 {
		return fmt.Errorf("store: description length must be between 1 and 50")
	}
	for i, t := range j.tasks {
		if t.ID == id {
			j.tasks[i].Description = description
			j.tasks[i].UpdatedAt = time.Now()

			return j.saveTasks()
		}
	}

	return fmt.Errorf("store: task with id [%d] not found", id)
}

func (j *JSONStore) DeleteTask(id int) (*task.Task, error) {
	idx := -1

	for i, t := range j.tasks {
		if t.ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		return nil, fmt.Errorf("store: task with id [%d] not found", id)
	}

	t := j.tasks[idx]
	j.tasks = append(j.tasks[:idx], j.tasks[idx+1:]...)

	err := j.saveTasks()
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (j *JSONStore) LastID() int {

	if len(j.tasks) == 0 {
		return 0
	}

	maxID := j.tasks[0].ID
	for _, t := range j.tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID
}

func (j *JSONStore) ListAllTasks() []task.Task {
	sort.Slice(j.tasks, func(i, k int) bool {
		si := statusOrder[j.tasks[i].Status]
		sk := statusOrder[j.tasks[k].Status]

		if si != sk {
			return si < sk
		}

		return j.tasks[i].UpdatedAt.After(j.tasks[k].UpdatedAt)
	})
	return j.tasks
}

func (j *JSONStore) ListTasksByStatus(status task.Status) []task.Task {
	result := make([]task.Task, 0)
	for _, t := range j.tasks {
		if t.Status == status {
			result = append(result, t)
		}
	}
	sort.Slice(result, func(i, k int) bool {
		return j.tasks[i].UpdatedAt.After(j.tasks[k].UpdatedAt)
	})
	return result
}

func (j *JSONStore) loadTasks() error {
	data, err := os.ReadFile(j.path)
	if err != nil {
		return fmt.Errorf("store: failed to read tasks: %w", err)
	}

	err = json.Unmarshal(data, &j.tasks)
	if err != nil {
		return fmt.Errorf("store: failed to unmarshal tasks: %w", err)
	}

	for _, t := range j.tasks {
		if err := task.ValidateTask(&t); err != nil {
			return fmt.Errorf("store: failed to validate task [%d]: %v", t.ID, err)
		}
	}
	return nil
}

func (j *JSONStore) saveTasks() error {
	arr, err := json.Marshal(j.tasks)
	if err != nil {
		return fmt.Errorf("store: failed to marshal tasks: %w", err)
	}
	return os.WriteFile(j.path, arr, 0644)
}

func NewJSONStore(path string) (*JSONStore, error) {

	if !strings.HasSuffix(path, ".json") {
		return nil, fmt.Errorf("store: path does not end with .json")
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.WriteFile(path, []byte("[]"), 0644)
		if err != nil {
			return nil, fmt.Errorf("store: failed to save tasks: %w", err)
		}
	}

	j := &JSONStore{path: path}
	err := j.loadTasks()
	if err != nil {
		return nil, err
	}

	return j, nil
}
