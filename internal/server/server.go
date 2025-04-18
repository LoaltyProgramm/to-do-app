package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartServer() error {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	http.Handle("/", http.FileServer(http.Dir("../web")))

	log.Printf("Server start on port: %v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		return fmt.Errorf("server is not work: %v", err)
	}

	return nil
}