package main

import (
	"api/src/app"
	"api/src/db"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	//Loading ENV variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	_, err = db.NewDbConnect()

	if err != nil {
		log.Fatal(err)
		return
	}

	app := app.NewServer()

	app.RunServer()

}
