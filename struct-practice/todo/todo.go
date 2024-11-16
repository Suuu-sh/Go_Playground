package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"Text"`
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

func (todo Todo) Display() {
	fmt.Printf("Your Todo has the following content: \n\n%v\n\n", todo.Text)
}

func New(todo string) (Todo, error) {
	if todo == "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: todo,
	}, nil
}
