package main

import (
	"os"
	"task-tracker/cmd"
)

func main() {
	if len(os.Args) < 2 {
		cmd.PrintUsage()
		return
	}

	cmd.HandleCommand(os.Args[1:])
}
