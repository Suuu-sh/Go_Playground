package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/practice-struct-practice/note"
	"example.com/practice-struct-practice/todo"
)

type saver interface {
	Save() error
}

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	todoText := getTodoData()
	todo, err := todo.New(todoText)

	todo.Display()
	err = saveData(todo)

	if err != nil {
		return
	}
	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		return
	}
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	return nil
}

func getTodoData() string {
	return getUserInput("Todo text:")
}
func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
