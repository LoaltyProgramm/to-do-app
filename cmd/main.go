package main

import (
	"go1f/pkg/server"
	"log"
)

func main() {
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
