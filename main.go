package main

import (
	"os"
	"task-tracker/cmd"
)

func main() {
	err := cmd.HandleCommand(os.Args)
	if err != nil {
		return
	}
}
