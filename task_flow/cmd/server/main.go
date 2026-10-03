package main

import (
	"learn-go/internal/handler"
	"learn-go/internal/repository"
	"learn-go/internal/service"
	"log"
	"net/http"
)

func main() {
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

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
