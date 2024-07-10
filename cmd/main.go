package main

import (
	"os"

	"github.com/akposieyefa/open-ai/database/seeders"
	"github.com/akposieyefa/open-ai/internals"
	"github.com/akposieyefa/open-ai/routers"
)

func main() {

	internals.ConnectToDB()

	arg := os.Args[1]

	if arg == "seeders" {
		seeders.Run()
	} else if arg == "server" {
		routers.ApiRoutes()
	}

}
