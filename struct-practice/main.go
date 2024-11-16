package main

import (
	"errors"
	"fmt"
)

func main() {
	getNoteData()
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note title:")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	content, err := getUserInput("Note content:")

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return title, content, err
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if value == "" {
		return "", errors.New("empty")
	}

	return value, nil
}
