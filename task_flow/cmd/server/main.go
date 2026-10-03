package main

import (
	"context"
	"learn-go/internal/database"
	"learn-go/internal/handler"
	"learn-go/internal/repository"
	"learn-go/internal/service"
	"log"
	"net/http"
)

func main() {
	databaseURL := "postgres://macmini@localhost:5432/taskflow"
	
	conn, err := database.Connect(databaseURL)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer conn.Close(context.Background())

	taskRepo := repository.NewTaskRepository()

	taskService := service.NewTaskService(taskRepo)

	taskHandler := handler.NewTaskHandler(taskService)

	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/tasks", taskHandler.Tasks)
	http.HandleFunc("GET /tasks/{id}", taskHandler.GetTaskById)
	http.HandleFunc("PUT /tasks/{id}", taskHandler.UpdateTask)
	http.HandleFunc("DELETE /tasks/{id}", taskHandler.DeleteTask)
	http.HandleFunc("DELETE /tasks", taskHandler.DeleteAllTask)

	log.Println("Starting server on port http://localhost:8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
