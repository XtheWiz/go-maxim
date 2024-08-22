package main

import (
	"fmt"
)

func ProfitCalculator() {
	var revenue float64
	var expense float64
	var taxRate float64

	revenue = getUserInput("Revenue: ")
	expense = getUserInput("Expense: ")
	taxRate = getUserInput("Tax Rate: ")

	// fmt.Print("Expense: ")
	// fmt.Scan(&expense)

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	// ebt := revenue - expense
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	ebt, profit, ratio := calculateFinancials(revenue, expense, taxRate)

	printCurrency("Earn Before Tax(EBT)", ebt)
	printCurrency("Profit", profit)
	printCurrency("Ratio", ratio)
}

func printCurrency(text string, number float64) {
	fmt.Printf("%s: %.2f\n", text, number)
}

func calculateFinancials(revenue, expense, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(text string) float64 {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}
