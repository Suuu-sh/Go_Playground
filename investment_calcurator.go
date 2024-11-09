package main

import (
	"fmt"
	"math"
)

func main() {
  const inflationRate = 2.5
  var investmentAmount float64
  var years float64
  expectedReturnRate := 5.5


  outputText("Investment Amoiunt: ")
  fmt.Scan(&investmentAmount)

  outputText("ReturenRate: ")
  fmt.Scan(&expectedReturnRate)

  outputText("Years: ")
  fmt.Scan(&years)

  futureValue,futureRealValue :=calculateFutureValue(investmentAmount,expectedReturnRate,years)


//   futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

//   fmt.Println("Future Value: ",futureValue)
//   formattedFV := fmt.Sprintf("Future Value: %.0f",futureValue)

  fmt.Printf("Future Value: %.0f,%.0f",futureValue,futureRealValue)
}

func outputText (text string){
  fmt.Print(text)
}

func calculateFutureValue(investmentAmount, expectedReturenRate, years float64) (fv float64,rfv float64){
  fv = investmentAmount * math.Pow(1+expectedReturenRate/100, years)
  rfv = 12
  return fv,rfv
}

