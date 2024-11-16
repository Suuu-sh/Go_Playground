package main

import (
	"fmt"

	"example.com/practice-struct-practice/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	fmt.Println(userNote)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if value == "" {
		return ""
	}

	return value
}
