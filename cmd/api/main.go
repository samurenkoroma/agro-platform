package main

import (
	"log"

	"github.com/samurenkoroma/agro-platform/internal/bootstrap"
)

func main() {

	app, err := bootstrap.New()

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()

	if err != nil {
		log.Fatal(err)
	}
}
