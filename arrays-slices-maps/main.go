package main

import "fmt"

func main() {
	userNames := make([]string, 2, 10)
	fmt.Println(cap(userNames))
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuxl")

}
