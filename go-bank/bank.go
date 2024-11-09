package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"
func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil{
		return 1000, errors.New("failed to find file")
	}

	balanceText := string(data)

	balance ,err :=strconv.ParseFloat(balanceText,64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}
func writeBalanceToFile(balance float64){
    balanceText :=fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile,[]byte(balanceText), 0644)
}
func main(){
	var accountBalance,err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------")
	}
	fmt.Println("Welcome to Go Bank")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. check balace")
		fmt.Println("2. deposit money")
		fmt.Println("3. withdraw money")
		fmt.Println("4. exit")
	
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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Goodbye")
			return
		}
	}
}