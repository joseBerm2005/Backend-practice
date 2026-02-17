package main

// API's entry point

import (
	"log"

	"github.com/joseBerm2005/Backend-practice/cmd/api"
)

// creates server and runs the Run function
func main() {
	server := api.NewAPIServer(":8000", nil)

	if err := server.Run(); err != nil {
		log.Fatal()
	}
}