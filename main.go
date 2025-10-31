package main

import (
	"os"
	"task-tracker/cmd"
	"task-tracker/store"

	"github.com/fatih/color"
)

func main() {

	s, err := store.NewJSONStore("tasks.json")
	if err != nil {
		color.Red(err.Error())
		return
	}

	err = cmd.HandleCommand(os.Args, s)
	if err != nil {
		color.Red(err.Error())
		return
	}
}
