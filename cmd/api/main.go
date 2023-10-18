package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ptflp/go-test/app/handler"
	"github.com/ptflp/go-test/app/repository"
	"github.com/ptflp/go-test/app/service"
	"github.com/ptflp/go-test/internal/router"
	"github.com/ptflp/go-test/internal/server"
	"log"
	"os"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")

	var storageTodos repository.TodoStorager
	switch dbDriver {
	case "inmemory":
		storageTodos = repository.NewInMemory()
	case "sqlite":
		storageTodos = repository.NewSqliteTodo()
	}

	appPort := os.Getenv("APP_PORT")

	serviceTodos := service.NewTodoService(storageTodos)
	handleTodos := handler.NewTodosHandler(serviceTodos)

	r := router.PrepareRoutes(handleTodos)
	fmt.Println("Server is running on port", appPort)
	httpServer := server.NewServer(fmt.Sprintf(":%s", appPort), server.WithRoutes(r))
	err = httpServer.Start()
	if err != nil {
		os.Exit(2)
	}
}
