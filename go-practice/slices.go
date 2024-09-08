package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"gaming", "reading", "learning"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	fmt.Println(hobbies[1:])

	// 3)
	favHobbies := hobbies[:2]
	fmt.Println(favHobbies)

	// 4)
	fmt.Println(cap(favHobbies))

	favHobbies = favHobbies[1:3]
	fmt.Println(favHobbies)
	fmt.Println(cap(favHobbies))

	// 5)
	goals := []string{"finish all!", "make prototype!"}
	fmt.Println(goals)

	// 6)
	goals[1] = "create app!"
	fmt.Println(goals)
	goals = append(goals, "good on golang!")
	fmt.Println(goals)

	// 7)
	products := []Product{
		{"prd-1", "First Product", 12.99},
		{"prd-2", "Second Product", 39.99},
	}
	fmt.Println(products)

	products = append(products, Product{"prd-3", "Third Product", 1.99})
	fmt.Println(products)

}
