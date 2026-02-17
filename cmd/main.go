package main

import (
	"log"

	"github.com/joseBerm2005/Backend-practice/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8000", nil)

	if err := server.Run(); err != nil {
		log.Fatal()
	}
}