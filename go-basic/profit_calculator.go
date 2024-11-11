package main

import (
	"errors"
	"fmt"
	"os"
)
func main(){
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue,err1  := getUserInput("Revenue: ")

	expenses,err2 := getUserInput("Expenses: ")

	taxRate,err3  := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil{
		fmt.Println(err1)
		return
	}

	ebt, profit, ratio := calculate_profit(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func getUserInput(text string)(float64,error){
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)


	if userInput <= 0 {
      return 0, errors.New("value must be apositive")
	}

	return userInput, nil
}
func calculate_profit (revenue, expenses, taxRate float64)(float64,float64,float64){
	ebt := revenue - expenses
	profit := ebt * (1-taxRate/100)
	ratio := ebt / profit

	WriteToFile(ebt,profit,ratio)

	return ebt, profit, ratio
}

func WriteToFile(ebt,profit,ratio float64){
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt,profit,ratio)
	os.WriteFile("results.txt",[]byte(results),0644)
}