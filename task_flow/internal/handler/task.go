package handler

import (
	"encoding/json"
	"learn-go/internal/model"
	"log"
	"net/http"
	"time"
)

var tasks []model.Task
var nextId = 1

func Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateTask(w, r)

	case http.MethodGet:
		GetTasks(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	task.ID = nextId
	nextId++

	task.CreatedAt = time.Now()

	task.UpdatedAt = time.Now()

	task.Status = "pending"

	tasks = append(tasks, task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := model.APIResponse{
		Status:  "success",
		Message: "Tasks fetched successfully",
		Data:    tasks,
	}

	json.NewEncoder(w).Encode(response)
}
