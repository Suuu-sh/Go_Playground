package main

import "fmt"
func main(){
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	expenses = getUserInput("Expenses: ")
	taxRate = getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculate_profit(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func getUserInput(text string)float64{
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}
func calculate_profit (revenue, expenses, taxRate float64)(ebt, profit, ratio float64){
	ebt = revenue - expenses
	profit = ebt * (1-taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}