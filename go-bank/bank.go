package main

import "fmt"

func main(){
	var accountBalance float64 = 1000 
	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. check balace")
	fmt.Println("2. deposit money")
	fmt.Println("3. withdraw money")
	fmt.Println("4. exit")

	var choice int
	fmt.Println("Your choice: ")
	fmt.Scan(&choice)

	fmt.Printf("Your choice %v",choice)

	if choice == 1 {
      fmt.Println("Your Balance is:", accountBalance)
	} else if choice == 2 {
	  fmt.Println("How much do you want to deposit?")
	  var depositAmount float64
	  fmt.Scan(&depositAmount)
	  accountBalance += depositAmount
	  fmt.Println("New accountBalance:",accountBalance)
	} else if choice == 3 {
		fmt.Println("How much do you want to withdraw?")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("New accountBalance", accountBalance)
	} else {
		fmt.Println("Goodbye")
	}

}