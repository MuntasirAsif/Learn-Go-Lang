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

func sendJSON(w http.ResponseWriter, statusCode int, status bool, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(model.APIResponse{
		Status:     status,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}

func (h *TaskHandler) Tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateTask(w, r)

	case http.MethodGet:
		h.GetTasks(w, r)

	default:
		sendJSON(w, http.StatusMethodNotAllowed, false, "Method not allowed", nil)
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, false, "Error parsing JSON: "+err.Error(), nil)
		return
	}

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.Status = "pending"

	createdTask := h.taskService.Create(task)

	sendJSON(w, http.StatusCreated, true, "Task created successfully", createdTask)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.taskService.GetAll()
	sendJSON(w, http.StatusOK, true, "Tasks fetched successfully", tasks)
}

func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, false, "Invalid task ID", nil)
		return
	}

	task, found := h.taskService.GetById(id)
	if !found {
		sendJSON(w, http.StatusNotFound, false, "task not available for this id", nil)
		return
	}

	sendJSON(w, http.StatusOK, true, "Task fetched successfully", task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, false, "Invalid task ID", nil)
		return
	}

	var updatedTask model.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, false, "Error parsing JSON: "+err.Error(), nil)
		return
	}

	updatedTask.UpdatedAt = time.Now()

	task, found := h.taskService.Update(id, updatedTask)
	if !found {
		sendJSON(w, http.StatusNotFound, false, "task not available for this id", nil)
		return
	}

	sendJSON(w, http.StatusOK, true, "Task updated successfully", task)
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		sendJSON(w, http.StatusBadRequest, false, "Invalid task ID", nil)
		return
	}

	deleted := h.taskService.Delete(id)
	if !deleted {
		sendJSON(w, http.StatusNotFound, false, "task not available for this id", nil)
		return
	}

	sendJSON(w, http.StatusOK, true, "Task deleted successfully", nil)
}

func (h *TaskHandler) DeleteAllTask(w http.ResponseWriter, r *http.Request) {
	h.taskService.DeleteAll()
	sendJSON(w, http.StatusOK, true, "All tasks deleted successfully", nil)
}
