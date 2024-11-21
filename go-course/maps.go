package main

import "fmt"

type Producst struct {
	id    string
	title string
	price float64
}

func main() {
	// websites := []string{"http", "https"}

	websites := map[string]string{
		"Google":     "https://ggogle",
		"Amazon web": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])
	websites["Linkedin"] = "https://linkedin"
	fmt.Println(websites["Linkedin"])
	delete(websites, "Google")
	fmt.Println(websites)
}
