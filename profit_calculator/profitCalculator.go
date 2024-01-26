package main

import "fmt"

func main() {
	var revenue, expense, taxRate float64

	revenue = getUserInput("Revenue")

	expense = getUserInput("Expense")

	taxRate = getUserInput("Tax Rate")

	ebt := revenue - expense

	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit

	fmt.Println("Profit: ", profit)

	fmt.Println("EBT: ", ebt)

	fmRatio := fmt.Sprintf("Ratio: %.2f \n", ratio)

	fmt.Println(fmRatio)

}

func getUserInput(text string) float64 {
	var userInput float64

	fmt.Printf("%v: ", text)

	if _, err := fmt.Scan(&userInput); err != nil {
		fmt.Println(("Please enter a valid input"))
		return 0
	}

	return userInput
}
