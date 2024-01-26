package main

import (
	"bank/utils"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const balanceFilePath = "./balance.txt"

func main() {

	fmt.Printf("Welcome to your bank %v", randomdata.FirstName(randomdata.Male))

	fmt.Println("Choose What do you want to do")

	balance, err := utils.ReadFromFile(balanceFilePath)

	if err != nil {
		panic(err)
	}

	var userOption int

	utils.RenderOptions()

	fmt.Scan(&userOption)

	for {
		switch userOption {
		case 1:
			fmt.Println("Your current balance is: ", balance)
		case 2:
			utils.DepositMoney(&balance)
		case 3:
			utils.WithDrawMoney(&balance)
		case 4:
			fmt.Println("Thanks for choosing us")
			fmt.Println("Have a good day")
			return
		default:
			fmt.Println("Please choose a valid option")
		}

		utils.RenderOptions()
		fmt.Scan((&userOption))
	}

}
