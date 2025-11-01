package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"tracker/store"
	"tracker/task"

	"github.com/fatih/color"
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

func printTasks(tasks []task.Task) {
	if len(tasks) == 0 {
		color.Yellow("No tasks found")
		return
	}

	for _, t := range tasks {

		var statusString, labelString, dateString string

		switch t.Status {
		case task.StatusDone:
			statusString = color.GreenString("(%s)", t.Status)
		case task.StatusTodo:
			statusString = color.YellowString("(%s)", t.Status)
		case task.StatusInProgress:
			statusString = color.CyanString("(%s)", t.Status)
		}

		if t.UpdatedAt.After(t.CreatedAt) {
			labelString = color.HiBlackString("| updated")
			dateString = color.HiBlackString(t.UpdatedAt.Format("2006-01-02 15:04:05"))
		} else {
			labelString = color.HiBlackString("| created")
			dateString = color.HiBlackString(t.CreatedAt.Format("2006-01-02 15:04:05"))
		}

		fmt.Printf(
			"%-17s %-25s %-50s %s %s\n",
			color.HiMagentaString("[%d]", t.ID),
			statusString,
			t.Description,
			labelString,
			dateString,
		)
	}
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
	id, err := s.AddTask(description)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf(
		"%s %s %s\n",
		color.GreenString("Task added:"),
		color.HiMagentaString("[%d]", id),
		description,
	)
}

func handleUpdate(args []string, s store.Store) {
	if len(args) < 3 {
		printUsage()
		return
	}

	id, err := parseID(args[1])
	if err != nil {
		color.Red(err.Error())
		return
	}

	description := strings.Join(args[2:], " ")

	err = s.UpdateTask(id, description)
	if err != nil {
		color.Red(err.Error())
		return
	}

	fmt.Printf(
		"%s %s %s\n",
		color.GreenString("Task updated:"),
		color.HiMagentaString("[%d]", id),
		description,
	)
}

func handleDelete(args []string, s store.Store) {
	if len(args) != 2 {
		printUsage()
		return
	}

	id, err := parseID(args[1])
	if err != nil {
		color.Red(err.Error())
		return
	}

	t, err := s.DeleteTask(id)
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Printf(
		"%s %s %s\n",
		color.GreenString("Task deleted:"),
		color.HiMagentaString("[%d]", t.ID),
		t.Description,
	)
}

func handleList(args []string, s store.Store) {
	if len(args) == 1 {
		printTasks(s.ListAllTasks())
		return
	}
	if len(args) == 2 {
		switch task.Status(args[1]) {
		case task.StatusDone:
			printTasks(s.ListTasksByStatus(task.StatusDone))
		case task.StatusInProgress:
			printTasks(s.ListTasksByStatus(task.StatusInProgress))
		case task.StatusTodo:
			printTasks(s.ListTasksByStatus(task.StatusTodo))
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
		color.Red(err.Error())
		return
	}

	err = s.ChangeStatus(id, task.StatusInProgress)
	if err != nil {
		color.Red(err.Error())
		return
	}

	fmt.Printf(
		"%s %s %s\n",
		color.GreenString("Task marked:"),
		color.HiMagentaString("[%d]", id),
		color.CyanString("(in-progress)"),
	)
}

func handleMarkDone(args []string, s store.Store) {
	if len(args) != 2 {
		printUsage()
		return
	}

	id, err := parseID(args[1])
	if err != nil {
		color.Red(err.Error())
		return
	}

	err = s.ChangeStatus(id, task.StatusDone)
	if err != nil {
		color.Red(err.Error())
		return
	}

	fmt.Printf(
		"%s %s %s\n",
		color.GreenString("Task marked:"),
		color.HiMagentaString("[%d]", id),
		color.GreenString("(done)"),
	)
}
