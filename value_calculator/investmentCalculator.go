package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 1

	var investment, year, expectedInterest float64

	fmt.Print("Enter Investment amount: ")
	fmt.Scan(&investment)

	fmt.Print("Enter Investment years: ")
	fmt.Scan(&year)

	fmt.Print("Enter Expected interest rate: ")
	fmt.Scan(&expectedInterest)

	futureValue := investment * math.Pow((1+(expectedInterest/100)), year)

	futureValue = futureValue / math.Pow((1+(inflationRate/100)), year)

	fmt.Println(futureValue)

}
