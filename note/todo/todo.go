package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("invalid input: text")
	}

	return Todo{
		Text: text,
	}, nil
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (note Todo) Save() error {
	fileName := "todo.json"

	todoContentJSON, err := json.Marshal(note)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, todoContentJSON, 0o644)
}
