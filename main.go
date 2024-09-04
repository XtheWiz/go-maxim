package main

import (
	"fmt"

	"example.com/bank/banks"

	"github.com/Pallinder/go-randomdata"
)

func main() {
	fmt.Println(randomdata.PhoneNumber())
	banks.Bank()
}
