package handler

import (
	"encoding/json"
	"learn-go/internal/model"
	"learn-go/internal/service"
	"net/http"
	"strconv"
	"time"
)

type TaskHandler struct {
	taskService *service.TaskService
}

func NewTaskHandler(taskService *service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTask(w, r)

	case http.MethodGet:
		h.GetTasks(w, r)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.Status = "pending"

	createdTask := h.taskService.Create(task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := model.APIResponse{
		Status:  "success",
		Message: "Task created successfully",
		Data:    createdTask,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tasks := h.taskService.GetAll()
	response := model.APIResponse{
		Status:  "success",
		Message: "Tasks fetched successfully",
		Data:    tasks,
	}

	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, found := h.taskService.GetById(id)
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := model.APIResponse{
		Status:  "success",
		Message: "Task fetched successfully",
		Data:    task,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask model.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask.UpdatedAt = time.Now()

	task, found := h.taskService.Update(id, updatedTask)
	if !found {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := model.APIResponse{
		Status:  "success",
		Message: "Task updated successfully",
		Data:    task,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	deleted := h.taskService.Delete(id)
	if !deleted {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := model.APIResponse{
		Status:  "success",
		Message: "Task deleted successfully",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *TaskHandler) DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	h.taskService.DeleteAll()

	w.Header().Set("Content-Type", "application/json")
	response := model.APIResponse{
		Status:  "success",
		Message: "All tasks deleted successfully",
	}

	json.NewEncoder(w).Encode(response)
}
