package main

import (
	"flag"
	"fmt"
	"todo/todolist"
)

func runApp() error {
	arguments, err := todolist.GetArguments()
	todolist.CheckErr(err)

	var storage todolist.Storage = todolist.Storage{Filename: "todos.json"}
	todos, err := storage.Load()
	todolist.CheckErr(err)

	var command string = arguments[1]
	var subCommand []string = arguments[2:]

	switch command {
	case "add":
		todolist.HandleAddCommand(subCommand, todos, storage)
	case "list":
		todolist.HandleListCommand(todos)
	case "done":
		todolist.HandleMarkDoneCommand(subCommand, todos, storage)
	case "delete":
		todolist.HandleDeleteCommand(subCommand, todos, storage)
	case "--help":
		flag.Usage = func ()  {
			fmt.Println("Available command are add, done, list, and delete. Use --help those command to see params")
		}
		flag.Usage()
	default:
		fmt.Println("Unknown command. Please provide a valid command. See available command using --help flag")
	}

	return nil
}

func main() {
	err := runApp()
	todolist.CheckErr(err)
}