package main

import (
	"fmt"
	"struct/user"
)

func main() {

	firstName := getUserInput("Enter First name")
	lastName := getUserInput("Enter Last name")
	birthDate := getUserInput("Enter Birth Date")

	appUser := user.New(firstName, lastName, birthDate)

	appUser.OutPutUserData()
	appUser.ClearUserName()
	appUser.OutPutUserData()

}

func getUserInput(inputText string) string {
	fmt.Println(inputText)

	userInput := ""

	fmt.Scan(&userInput)

	return userInput

}
