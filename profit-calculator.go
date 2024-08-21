package main

import "fmt"

func ProfitCalculator() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expense: ")
	fmt.Scan(&expense)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("Earn Before Tax(EBT):", ebt)
	fmt.Println("Profit:", profit)
	fmt.Println("Ratio:", ratio)
}
