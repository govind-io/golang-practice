package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(inputText, errorText string) (string, error) {

	fmt.Println(inputText)

	reader := bufio.NewReader(os.Stdin)

	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(errorText)
		return "", err
	}

	userInput = strings.TrimSuffix(userInput, "\n")

	userInput = strings.TrimSuffix(userInput, "\r")

	return userInput, err

}

func WriteToFile(text, filename, filetype string) {
	filepath := fmt.Sprintf("./data/%v/%v.json", filetype, filename)
	err := os.WriteFile(filepath, []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}
