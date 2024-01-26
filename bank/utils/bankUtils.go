package utils

import (
	"errors"
	"fmt"
)

const balanceFilePath = "./balance.txt"

func DepositMoney(balance *float64) error {

	depositInputText := "Enter Deposti Amount"
	depositInputError := "Invalid Deposit Amount"

	depositedAmount, err := GetUserInput(depositInputText, depositInputError)

	if err != nil {
		return err
	}

	*balance = *balance + depositedAmount

	StoreToFile(*balance, balanceFilePath)

	return nil

}

func WithDrawMoney(balance *float64) error {
	withdrawInputText := "Enter Withdraw Amount"
	withdrawInputError := "Invalid Withdraw Amount"

	withdrawnAmount, err := GetUserInput(withdrawInputText, withdrawInputError)

	if err != nil {
		return err
	}

	updatedBalance := *balance - withdrawnAmount

	if updatedBalance < 0 {
		fmt.Println("Not enough balance")
		return errors.New("not enough balance")
	}

	*balance = updatedBalance

	StoreToFile(*balance, balanceFilePath)

	return nil
}
