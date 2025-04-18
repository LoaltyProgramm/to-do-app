package main

import (
	"log"

	"github.com/LoaltyProgramm/to-do-app/internal/config"
	"github.com/LoaltyProgramm/to-do-app/internal/db"
	"github.com/LoaltyProgramm/to-do-app/internal/handlers"
	"github.com/LoaltyProgramm/to-do-app/internal/repository"
	"github.com/LoaltyProgramm/to-do-app/internal/server"
	"github.com/LoaltyProgramm/to-do-app/internal/service"
)

func main() {
	cfg := config.Config{}
	if err := cfg.GetEnv(); err != nil {
		log.Fatal(err)
	}
	log.Println("Read env file complited")

	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	log.Println("Initialisation db file complited")

	DB := repository.NewTaskRepository(db.DB)
	taskService := service.NewTaskService(DB)
	taskHandlers := handlers.NewTaskHandlers(taskService)
	log.Println("Service task complted")

	taskHandlers.InitHandler()
	log.Println("Handlers start complited")

	log.Println("Server start complited")
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}	
