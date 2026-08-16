package handler

import (
	"encoding/json"
	"learn-go/internal/model"
	"log"
	"net/http"
	"strconv"
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
	response := model.APIResponse{
		Status:  "success",
		Message: "Task created successfully",
		Data:    task,
	}
	json.NewEncoder(w).Encode(response)
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

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal("Error parsing ID: ", err)
	}

	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			response := model.APIResponse{
				Status:  "success",
				Message: "Task fetched successfully",
				Data:    task,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Fatal("Error parsing ID: ", err)
	}

	var updatedTask model.Task

	err = json.NewDecoder(r.Body).Decode(&updatedTask)

	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	updatedTask.ID = id
	updatedTask.UpdatedAt = time.Now()

	for i, task := range tasks {
		if task.ID == id {
			tasks[i] = updatedTask
			w.Header().Set("Content-Type", "application/json")
			response := model.APIResponse{
				Status:  "success",
				Message: "Task updated successfully",
				Data:    updatedTask,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	http.Error(w, "Task not found", http.StatusNotFound)
}
