package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userName := make([]string, 2, 5)
	userName[0] = "Julie"

	userName = append(userName, "Max")
	userName = append(userName, "Manuel")

	fmt.Println(userName)

	// making map
	courseRatings := make(floatMap, 3)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.5

	// type alias
	fmt.Println(courseRatings)
	courseRatings.output()

	// for loop - slice
	for index, value := range userName {
		fmt.Println(index, value)
	}

	// for loop - map
	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
