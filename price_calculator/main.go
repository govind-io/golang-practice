package main

import (
	"fmt"
	"price_calculators/price"
)

func main() {

	taxRates := [5]float64{0.5, 0.2, 0.3, 0.15, 0.11}

	prices := [len(taxRates)]*price.TaxIncludedPricesJob{}

	for index, rate := range taxRates {
		prices[index] = price.NewTaxIncludedPricesJob(rate)
	}

	for _, priceing := range prices {
		fmt.Println(priceing)
	}

}
