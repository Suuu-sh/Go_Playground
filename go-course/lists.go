package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

type TemperatureData struct {
	day1 float64
}

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	prices[1] = 11.99

	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices)
}

// func main() {
// 	var productNames [4]string
// 	prices := [4]float64{10.99, 9.3, 34.4, 29.0}

// 	productNames[2] = "carpet"

// 	fmt.Println(prices[2])
// 	fmt.Println(productNames)

// 	// 左にスライスするとcapacityも減る
// 	featurePrices := prices[:2]
// 	featurePrices[0] = 199.99
// 	highlightedPrices := featurePrices[:3]
// 	fmt.Println(featurePrices)
// 	fmt.Println(prices)

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(featurePrices), cap(featurePrices))
// }
