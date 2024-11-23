package main

import (
	"fmt"
)

type TransformFn func(int) int
type anotherFn func(int, []string, map[string][]int) ([]int, string)

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{3, 1, 3, 4}

	//copyが生成されないようにpointer &を使用
	// doubled := transformNumbers(&numbers, double)
	// tripled := transformNumbers(&numbers, triple)

	// fmt.Println(doubled)
	// fmt.Println(tripled)

	transformd := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})
	fmt.Println(transformd)

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	transformerFn1 := transformFn(&numbers)
	transformerFn2 := transformFn(&moreNumbers)

}

func transformFn(numbers *[]int) TransformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNumbers(numbers *[]int, transform TransformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

// func double(number int) int {
// 	return number * 2
// }

// func triple(number int) int {
// 	return number * 3
// }
