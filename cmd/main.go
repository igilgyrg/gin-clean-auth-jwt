package main

import (
	"github.com/igilgyrg/gin-todo/internal/app"
)

func main() {
	server := app.New()
	server.StartServer()
}
