package utils

import (
	"fmt"
	"os"
	"strconv"
)

func StoreToFile(value float64, filepath string) error {
	balanceText := fmt.Sprint(value)
	err := os.WriteFile(filepath, []byte(balanceText), 0644)

	return err
}

func ReadFromFile(filepath string) (float64, error) {
	valueText, err := os.ReadFile(filepath)

	if err != nil {
		return 0, err
	} else {
		value, err := strconv.ParseFloat(string((valueText)), 64)

		if err != nil {
			return 0, err
		}

		return value, nil
	}
}
