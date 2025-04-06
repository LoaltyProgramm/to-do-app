package server

import (
	"log"
	"net/http"
	"os"
	"fmt"
)

func StartServer() error {
	port := os.Getenv("TODO_PORT")

	http.Handle("/", http.FileServer(http.Dir("../web")))

	log.Printf("Server start on port: %v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		return fmt.Errorf("Server is not work: %v", err)
	}

	return nil
}