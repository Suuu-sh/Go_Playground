package main

import "fmt"

func main() {
	numbers := []int{1, 10, 10}
	sum := sumup(1, 19, 12)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

func sumup(firstNumber int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
