package todolist

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const STORAGEDIR string = "storage"

type Storage struct {
	Filename string
}

func (storage *Storage) fullPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var storagePath string = filepath.Join(currentDir, STORAGEDIR, storage.Filename)
	_, err = os.Stat(storagePath)

	if err == nil {
		return storagePath, nil
	}

	f, err := os.Create(storagePath)
	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err = f.WriteString("[]")
	if err != nil {
		return "", err
	}

	f.Sync()

	return storagePath, nil
}

func (storage *Storage) Load() (Todolist, error) {
	storagePath, err := storage.fullPath()
	if err != nil {
		return Todolist{}, err
	}

	strTodos, err := os.ReadFile(storagePath)
	if err != nil {
		return Todolist{}, err
	}

	todos := Todolist{}
	json.Unmarshal(strTodos, &todos)

	return todos, nil
}

func (storage *Storage) Save(todos Todolist) error {
	storagePath, err := storage.fullPath()
	if err != nil {
		return err
	}

	byteTodos, err := json.Marshal(todos)
	err = os.WriteFile(storagePath, byteTodos, 0777)

	return err
}