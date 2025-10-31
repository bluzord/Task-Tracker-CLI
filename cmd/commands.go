package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"task-tracker/store"
	"task-tracker/task"
)

type Command string

const (
	CommandAdd            Command = "add"
	CommandUpdate         Command = "update"
	CommandDelete         Command = "delete"
	CommandList           Command = "list"
	CommandMarkInProgress Command = "mark-in-progress"
	CommandMarkDone       Command = "mark-done"
)

func printUsage() {
	fmt.Println(`Usage:
  task-cli add [description]              Add a new task
  task-cli update [id] [description]      Update a task
  task-cli delete [id]                    Delete a task
  task-cli list                           List all tasks
  task-cli list [done|todo|in-progress]	  List tasks by status
  task-cli mark-in-progress [id]          Mark task as in progress
  task-cli mark-done [id]                 Mark task as done`)
}

func HandleCommand(args []string, s store.Store) error {

	if len(args) < 2 {
		printUsage()
		return fmt.Errorf("cmd: invalid number of arguments")
	}

	switch Command(args[1]) {
	case CommandAdd:
		handleAdd(args[1:], s)
	case CommandUpdate:
		handleUpdate(args[1:], s)
	case CommandDelete:
		handleDelete(args[1:], s)
	case CommandList:
		handleList(args[1:], s)
	case CommandMarkInProgress:
		handleMarkInProgress(args[1:], s)
	case CommandMarkDone:
		handleMarkDone(args[1:], s)
	default:
		printUsage()
	}
	return nil
}

func parseID(arg string) (int, error) {
	id, err := strconv.Atoi(arg)
	if err != nil {
		return 0, fmt.Errorf("cmd: invalid ID: %s (must be integer)", arg)
	}
	return id, nil
}

func handleAdd(args []string, s store.Store) {
	if len(args) < 2 {
		printUsage()
		return
	}

	description := strings.Join(args[1:], " ")
	fmt.Println("Add description:", description)
}

func handleUpdate(args []string, s store.Store) {
	if len(args) < 3 {
		printUsage()
		return
	}

	id, err := parseID(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	description := strings.Join(args[2:], " ")

	fmt.Printf("Update %d description: %s\n", id, description)
}

func handleDelete(args []string, s store.Store) {
	if len(args) != 2 {
		printUsage()
		return
	}

	id, err := parseID(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Deleted id:", id)
}

func handleList(args []string, s store.Store) {
	if len(args) == 1 {
		fmt.Println("List all tasks")
		return
	}
	if len(args) == 2 {
		switch task.Status(args[1]) {
		case task.StatusDone:
			fmt.Println("List done tasks")
		case task.StatusInProgress:
			fmt.Println("List in-progress tasks")
		case task.StatusTodo:
			fmt.Println("List todo tasks")
		default:
			printUsage()
		}
		return
	}
	printUsage()
}

func handleMarkInProgress(args []string, s store.Store) {
	if len(args) != 2 {
		printUsage()
		return
	}

	id, err := parseID(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mark In Progress task", id)
}

func handleMarkDone(args []string, s store.Store) {
	if len(args) != 2 {
		printUsage()
		return
	}

	id, err := parseID(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Mark Done task", id)
}
