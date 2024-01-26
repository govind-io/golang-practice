package todos

import (
	"encoding/json"
	"fmt"
	utils "notes/Utils"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func New() *Todo {
	userInout, err := utils.GetUserInput("Enter the to do text", "Invalid todo text")

	if err != nil {
		return &Todo{}
	}

	return &Todo{
		Text: userInout,
	}
}

func (todo *Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo *Todo) Save() {
	todoText := strings.ReplaceAll(todo.Text, " ", "_")
	todoText = strings.ToLower(todoText)

	todoJsonContent, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Unable to parse todo data into json")
		return
	}

	utils.WriteToFile(string(todoJsonContent[:]), todoText, "todos")
}
