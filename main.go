package main

import (
	"lbms/app"
	"log"
)

func main() {
	server, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
