package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"go1f/pkg/api"
)

func StartServer() error {
	port := os.Getenv("TODO_PORT")

	api.Init()
	
	http.Handle("/", http.FileServer(http.Dir("../web")))

	log.Printf("Server start on port: %v", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		return fmt.Errorf("Server is not work: %v", err)
	}

	return nil
}