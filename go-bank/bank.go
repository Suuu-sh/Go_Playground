package main

import (
	"fmt"

	"example.com/bank-practice/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main(){
	var accountBalance,err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
	}
	fmt.Println("Welcome to Go Bank")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentOptions()
	
		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)
	
		switch choice {
		case 1:
			fmt.Println("Your Balance is:", accountBalance)
		case 2:
			fmt.Println("How much do you want to deposit?")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	  
			if depositAmount <= 0{
			  fmt.Println("Invalid amount.")
			  // return
			  continue
			}
			accountBalance += depositAmount
			fmt.Println("New accountBalance:",accountBalance)
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
		case 3:
			fmt.Println("How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
	
			if withdrawAmount <= 0 {
				fmt.Println("You cant withdraw<")
				continue
			  }
	
			if withdrawAmount > accountBalance {
				fmt.Println("You cant withdraw>")
				continue
			  }
			accountBalance -= withdrawAmount
			fmt.Println("New accountBalance", accountBalance)
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}