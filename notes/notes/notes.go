package notes

import (
	"encoding/json"
	"fmt"
	utils "notes/Utils"
	"strings"
	"time"
)

type Note struct {
	Title       string    `json:"title"`
	Created_at  time.Time `json:"created_at"`
	Description string    `json:"description"`
}

func New() *Note {
	title, titleErr := utils.GetUserInput("Enter note title", "Invalid note title")

	if titleErr != nil {
		fmt.Println("Invalid note title input")
		return &Note{}
	}

	description, descErr := utils.GetUserInput("Enter note description", "Invalid note description")

	if descErr != nil {
		fmt.Println("Invalid note description input")
		return &Note{}
	}

	return &Note{
		Title:       title,
		Description: description,
		Created_at:  time.Now(),
	}
}

func (note *Note) GetValues() (string, string, time.Time) {
	return note.Title, note.Description, note.Created_at
}

func (note *Note) Display() {
	fmt.Printf("Your Note: %v\nContent: %v \n", note.Title, note.Description)
}

func (note *Note) Save() {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	noteData, err := json.Marshal(note)

	if err != nil {
		fmt.Println("Unable to parse note data into json")
		return
	}

	utils.WriteToFile(string(noteData[:]), fileName, "notes")

}
