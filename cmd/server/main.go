package main

import (
	"os"

	chess "github.com/utatti/chess"
)

func main() {
	var port string
	if envPort, ok := os.LookupEnv("PORT"); ok {
		port = envPort
	} else {
		port = "8080"
	}
	server := chess.NewServer()
	server.Start(":" + port)
}
