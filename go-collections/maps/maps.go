package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"google": "https://google.com",
		"amazon": "https://aws.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["amazon"])
	fmt.Println(websites["amazon web service"])

	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "google")
	fmt.Println(websites)
}
