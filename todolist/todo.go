package todolist

import (
	"errors"
	"fmt"
)

type Todo struct {
	Title  string
	Status bool
}

func (todo *Todo) StatusDesc() string {
	switch todo.Status {
	case true:
		return "Done"
	default:
		return "Pending"
	}
}

type Todolist struct {
	Data []Todo
}

func (todolist *Todolist) Add(title string) (*Todo, error) {
	if title == "" {
		return nil, errors.New("Title required")
	}

	var todo Todo = Todo{Title: title}
	todolist.Data = append(todolist.Data, todo)
	return &todo, nil
}

func (todolist *Todolist) Show() {
	if len(todolist.Data) <= 0 {
		fmt.Println("No todo available");
		return
	}

	for index, todo := range todolist.Data {
		var title string = todo.Title
		var status string = todo.StatusDesc()

		fmt.Printf("ID: %d | Title: %v | Status: %v\n", index, title, status)
	}
}

func (todolist *Todolist) FindById(id int) (t *Todo, err error) {
	if err := todolist.ValidateId(id); err != nil {
		return t, err
	}

	for index, todo := range todolist.Data {
		if index == id {
			return &todo, nil
		}
	}

	return t, errors.New("Todo not found")
}

func (todolist *Todolist) DeleteById(id int) error {
	if err := todolist.ValidateId(id); err != nil {
		return err
	}

	todolist.Data = append(todolist.Data[:id], todolist.Data[id+1:]...)
	return nil
}

func (todolist *Todolist) UpdateById(id int, newTodo Todo) error {
	if err := todolist.ValidateId(id); err != nil {
		return err
	}

	for index := range todolist.Data {
		if index == id {
			todolist.Data[index] = newTodo
			break
		}
	}

	return nil
}

func (todolist *Todolist) MarkDoneById(id int) error {
	todo, err := todolist.FindById(id)
	if err != nil {
		return err
	}

	todo.Status = true;
	return nil
}

func (todolist *Todolist) ValidateId(id int) error {
	if id < 0 {
		return errors.New("ID must at least 0")
	}

	if id >= len(todolist.Data) {
		return errors.New("ID not found")
	}

	return nil
}

func New() Todolist {
	return Todolist{Data: []Todo{}}
}