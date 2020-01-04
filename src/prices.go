package main

import "fmt"

// FinalPrice - calculates the final price in dollars and real
func FinalPrice(costPrice float64) (dollarPrice, realPrice float64) {
	profitFactor := 1.33
	conversionRate := 2.34

	dollarPrice = costPrice * profitFactor
	realPrice = dollarPrice * conversionRate

	return
}

func main() {
	dollarPrice, realPrice := FinalPrice(34.99)

	fmt.Printf("Final price in dollars: %.2f\n"+"Final price in real: %.2f\n", dollarPrice, realPrice)
}
