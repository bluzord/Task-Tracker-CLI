package main

import (
	"fmt"
	"os"
	"task-tracker/cmd"
	"task-tracker/store"
)

func main() {

	s, err := store.NewJSONStore("tasks.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cmd.HandleCommand(os.Args, s)
	if err != nil {
		fmt.Println(err)
		return
	}
}
