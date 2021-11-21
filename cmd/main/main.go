package main

import (
	"IAGE-test-task/cmd/api"
	"log"
)

func main() {
	log.Println("App started...")
	app, err := api.BuildInRuntime()
	if err != nil {
		panic(err)
	}
	err = app.Start()
	if err != nil {
		panic(err)
	}
}

