package main

import (
	"IAGE-test-task/cmd/api"
	"fmt"
)

func main() {
	fmt.Println("App started...")
	app, err := api.BuildInRuntime()
	if err != nil {
		panic(err)
	}
	err = app.Start()
	if err != nil {
		panic(err)
	}
}

