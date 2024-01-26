package utils

import (
	"errors"
	"fmt"
)

func GetUserInput(inputText, errText string) (float64, error) {
	userInput := 0.0

	fmt.Println(inputText)

	_, err := fmt.Scan(&userInput)

	if err != nil || userInput <= 0 {
		fmt.Println(errText)
		return 0, errors.New(errText)
	}

	return userInput, nil

}
