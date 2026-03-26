package todolist

import (
	"errors"
	"flag"
	"log"
	"os"
)

func HandleAddCommand(arguments []string, todos Todolist, storage Storage) {
	var AddCommand *flag.FlagSet = flag.NewFlagSet("add", flag.ExitOnError)
		var title *string = AddCommand.String("title", "", "add a todo title")

		AddCommand.Parse(arguments)

		_, err := todos.Add(*title)
		CheckErr(err)

		todos.Show()
		storage.Save(todos)
}

func HandleListCommand(todos Todolist) {
	todos.Show()
}

func HandleMarkDoneCommand(arguments []string, todos Todolist, storage Storage) {
	var DoneCommand *flag.FlagSet = flag.NewFlagSet("done", flag.ExitOnError)
		var todoId *int = DoneCommand.Int("id", 0, "todo id")

		DoneCommand.Parse(arguments)

		err := todos.MarkDoneById(*todoId)
		CheckErr(err)

		todos.Show()
		storage.Save(todos)
}

func HandleDeleteCommand(arguments []string, todos Todolist, storage Storage) {
	var DeleteCommand *flag.FlagSet = flag.NewFlagSet("delete", flag.ExitOnError)
		var todoId *int = DeleteCommand.Int("id", 0, "todo id")

		DeleteCommand.Parse(arguments[2:])

		err := todos.DeleteById(*todoId)
		CheckErr(err)

		todos.Show()
		storage.Save(todos)
}

func GetArguments() ([]string, error) {
	var arguments []string = os.Args

	if len(arguments) <= 1 {
		return []string{}, errors.New("Please provide a command")
	}

	return arguments, nil
}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}