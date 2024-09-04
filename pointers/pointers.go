package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age:", *agePointer)

	getAdultYear(agePointer)
	fmt.Println("Get Adult:", age)
}

func getAdultYear(age *int) {
	*age = *age - 18
}
