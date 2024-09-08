package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])

	prices[1] = 9.99

	prices = append(prices, 5.99, 12.99, 3.99, 29.99, 11.99)
	fmt.Println(prices)

	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}

// func main() {
// 	var productNames = [4]string{"A Book"}

// 	prices := [4]float64{10.99, 9.99, 4.99, 7.99}

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)
// 	fmt.Println(prices)
// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:3]
// 	fmt.Println(featuredPrices)
// }
