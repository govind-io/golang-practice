package app

import (
	"api/src/event"
	appUtil "api/src/utils"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App *fiber.App
}

func NewServer() *Server {
	return &Server{
		App: fiber.New(),
	}
}

func (app *Server) RunServer() {
	app.App.Mount("/event", event.GetEventApp())

	port, err := strconv.ParseInt(appUtil.GetEnv("port"), 10, 64)

	if err != nil {
		port = 3000
	}

	if port == 0 {
		port = 3001
	}

	fmt.Println(port)

	err = app.App.Listen(fmt.Sprintf(":%v", port))

	if err != nil {
		log.Fatal(err)
	}
}
