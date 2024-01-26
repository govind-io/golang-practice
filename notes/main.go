package main

import (
	"notes/notes"
	"notes/todos"
)

type outputable interface {
	Display()
	Save()
}

func main() {
	newNote := notes.New()

	todo := todos.New()

	output(todo)
	output(newNote)
}

func output(struc outputable) {
	struc.Display()
	struc.Save()

}
