package main

import (
	"github.com/akposieyefa/open-ai/internals"
	"github.com/akposieyefa/open-ai/routers"
)

func main() {
	internals.ConnectToDB()
	routers.ApiRoutes()
}
