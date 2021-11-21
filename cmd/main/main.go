package main

import (
	"IAGE-test-task/cmd/api"
)

func main() {
	app, err := api.BuildInRuntime()
	if err != nil {
		panic(err)
	}
	err = app.Start()
	if err != nil {
		panic(err)
	}
}

