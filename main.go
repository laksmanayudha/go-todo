package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Todo struct {
	Title string
	Status bool
}

type TodoList []Todo

func AddTodo(todos TodoList, title string) (TodoList, error) {
	if (title == "") {
		return nil, errors.New("Title required")
	}

	var newTodo Todo = Todo{Title: title}
	return append(todos, newTodo), nil
}

func TransformStatus(status bool) string {
	if status { return "Done" }
	return "Pending"
}

func ShowTodos(todos TodoList) {
	for index, todo := range todos {
		var num int = index + 1
		var title string = todo.Title
		var status string = TransformStatus(todo.Status)

		fmt.Printf("%d. title: %v | status: %v\n", num, title, status)
	}
}

func main() {
	// check len of arguments
	var arguments []string = os.Args
	if len(arguments) <= 1 {
		fmt.Println("Please provide a command. Use --help to see available commands")
		os.Exit(1)
	}

	var todoList TodoList = TodoList{
		Todo{Title: "Learning Golang"},
		Todo{Title: "Create POS Application"},
		Todo{Title: "Update Todo CLI Program"},
	}

	var command string = arguments[1]
	switch command {
	case "add":
		// create add subcommand
		var AddCommand *flag.FlagSet = flag.NewFlagSet("add", flag.ExitOnError)
		var title *string = AddCommand.String("title", "", "add a todo title")

		AddCommand.Parse(arguments[2:])

		todoList, err := AddTodo(todoList, *title)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		ShowTodos(todoList)
	case "list":
		ShowTodos(todoList)
	default:
		fmt.Println("Unknown command. Please provide a valid command. See available command using --help flag")
	}
}