package prices

import (
	"bufio"
	"fmt"
	"os"

	"example.com/practice-project/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("price.txt")
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()

	if err != nil {
		fmt.Println("Error")
		file.Close()
		return
	}
	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		file.Close()
	}

	job.InputPrices = prices
	file.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.3f", price)] = fmt.Sprintf("%.3f", taxIncludedPrice)
	}

	fmt.Println(job.TaxRate)
	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
