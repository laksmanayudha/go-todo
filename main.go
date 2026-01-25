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
	if len(todos) <= 0 {
		fmt.Println("No todo available")
		return
	}

	for index, todo := range todos {
		var title string = todo.Title
		var status string = TransformStatus(todo.Status)

		fmt.Printf("ID: %d | title: %v | status: %v\n", index, title, status)
	}
}

func MarkDoneTodoById(todos TodoList, id int) (TodoList, error) {
	todo, err := FindTodoById(todos, id)
	if err != nil {
		return todos, err
	}

	todo.Status = true
	updatedTodos, err := UpdateTodoById(todos, todo, id)

	return updatedTodos, nil
}

func UpdateTodoById(todos TodoList, newTodo Todo, id int) (TodoList, error) {
	if err := ValidateId(todos, id); err != nil {
		return todos, err
	}

	for index := range todos {
		if index == id {
			todos[index] = newTodo
			break
		}
	}

	return todos, nil
}

func DeleteTodoById(todos TodoList, id int) (TodoList, error) {
	if err := ValidateId(todos, id); err != nil {
		return todos, err
	}

	var newTodos TodoList = append(todos[:id], todos[id+1:]...)

	return newTodos, nil
}

func FindTodoById(todos TodoList, id int) (t Todo, err error) {
	if err := ValidateId(todos, id); err != nil {
		return t, err
	}

	for index, todo := range todos {
		if index == id {
			return todo, nil
		}
	}

	return t, errors.New("Todo not found")
}

func ValidateId(todos TodoList, id int) error {
	if id < 0 {
		return errors.New("ID must at least 0")
	}

	if id >= len(todos) {
		return errors.New("ID not found")
	}

	return nil
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
	case "done":
		var DoneCommand *flag.FlagSet = flag.NewFlagSet("done", flag.ExitOnError)
		var todoId *int = DoneCommand.Int("id", 0, "todo id")

		DoneCommand.Parse(arguments[2:])

		todoList, err := MarkDoneTodoById(todoList, *todoId)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		ShowTodos(todoList)
	case "delete":
		var DeleteCommand *flag.FlagSet = flag.NewFlagSet("delete", flag.ExitOnError)
		var todoId *int = DeleteCommand.Int("id", 0, "todo id")

		DeleteCommand.Parse(arguments[2:])

		todoList, err := DeleteTodoById(todoList, *todoId)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		ShowTodos(todoList)
	default:
		fmt.Println("Unknown command. Please provide a valid command. See available command using --help flag")
	}
}