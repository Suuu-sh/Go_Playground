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

// type displayer interface {
// 	Display()
// }

// type outputtable interface {
// 	Save() error
// 	Display()
// }

type outputtable interface {
	Save() error
	Display()
}

func main() {
	printSomething(1)
	title, content := getNoteData()
	todoText := getUserInput("Todo text:")

	userNote, err := note.New(title, content)

	if err != nil {
		return
	}
	todo, err := todo.New(todoText)

	if err != nil {
		return
	}

	todo.Display()
	err = outputData(todo)

	if err != nil {
		return
	}
	err = outputData(userNote)

	if err != nil {
		return
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func printSomething(value interface{}) {
	fmt.Println(value)
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	return nil
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
