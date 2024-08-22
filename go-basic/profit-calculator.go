package main

import (
	"errors"
	"fmt"
	"os"
)

func storeResult(ebt, profit, ratio float64) {
	results := fmt.Sprintf(`
	EBT: %.1f
	Profit: %.1f
	Ratio: %.3f
	`, ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
}

func ProfitCalculator() {
	var revenue float64
	var expense float64
	var taxRate float64
	var err error

	revenue, err = getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Revenue value is not greater than 0")
		panic("Can't continue the process, bye 😎😎😎")
	}

	expense, err = getUserInput("Expense: ")
	if err != nil {
		fmt.Println("Expense value is not greater than 0")
		panic("Can't continue the process, bye 😎😎😎")
	}

	taxRate, err = getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println("Tax rate value is not greater than 0")
		panic("Can't continue the process, bye 😎😎😎")
	}

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

	storeResult(ebt, profit, ratio)
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

func getUserInput(text string) (float64, error) {
	var userInput float64

	fmt.Print(text)
	fmt.Scan(&userInput)

	var err error
	if userInput <= 0 {
		err = errors.New("invalid input, the number have to greater than 0")
	}

	return userInput, err
}
