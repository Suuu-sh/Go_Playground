package main

import "fmt"

func main() {
	userNames := make([]string, 2, 10)
	fmt.Println(cap(userNames))
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuxl")

	courseRatings := make(map[string]float64, 3)
	courseRatings["go"] = 5.4
	courseRatings["go3"] = 5.4
	courseRatings["go4"] = 5.4
	courseRatings["go5"] = 5.4
	courseRatings["go53"] = 2.4

	fmt.Println(courseRatings)

	for index, value := range userNames {
		fmt.Println(index)
		fmt.Println(value)
	}

	for index, value := range courseRatings {
		fmt.Println(index)
		fmt.Println(value)
	}
}
