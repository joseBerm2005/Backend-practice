package main

import ("github.com/joseBerm2005/Backend-practice")

func main() {
	server := api.NewAPIServer(":8000", nil)

	if err := server.Run(); err != nil {
		log
	}
}