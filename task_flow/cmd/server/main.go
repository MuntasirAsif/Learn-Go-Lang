package main

import (
	"learn-go/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", handler.Health)
	http.HandleFunc("/tasks", handler.Tasks)
	http.HandleFunc("GET /tasks/{id}", handler.GetTaskById)
	http.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
	http.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)
	http.HandleFunc("DELETE /tasks", handler.DeleteAllTask)

	log.Println("Starting server on port http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
