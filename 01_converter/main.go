package main

import "fmt"

func main() {
	const rateUSDtoEURO float64 = 0.95
	const rateUSDtoRUB float64 = 80
	var rateEUROtoRUB float64 = rateUSDtoRUB / rateUSDtoEURO
	fmt.Print(rateEUROtoRUB)

}