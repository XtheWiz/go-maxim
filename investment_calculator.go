package main

import (
	"fmt"
	"math"
)

func InvestmentCalculator() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	// fmt.Println("Future Value (adjusted for inflation):", futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	fmt.Print(formattedFV, formattedRFV)
}
