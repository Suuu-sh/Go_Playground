package main

import "fmt"
func main(){
	age := 42

	agePointer := &age

	fmt.Println("Age", *agePointer)
	adultYears := getAdultYears(agePointer)
	fmt.Println(adultYears)
}

func getAdultYears(age *int) int  {
	return *age - 18
}