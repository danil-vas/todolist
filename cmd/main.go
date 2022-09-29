package main

import (
	"log"
	"todolist"
	"todolist/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todolist.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}
}
