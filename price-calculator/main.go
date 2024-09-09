package main

import (
	"fmt"

	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdd := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdd, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("could not process the job")
			fmt.Println(err)
		}
	}
}
