package service

import (
	"learn-go/internal/model"
	"learn-go/internal/repository"
)

type TaskService struct {
	taskRepo *repository.TaskRepository
}

func NewTaskService(taskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{taskRepo: taskRepo}
}

func (s *TaskService) Create(task model.Task) model.Task {
	return s.taskRepo.Create(task)
}

func (s *TaskService) GetAll() []model.Task {
	return s.taskRepo.GetAll()
}

func (s *TaskService) GetById(id int) (model.Task, bool) {
	return s.taskRepo.GetById(id)
}

func (s *TaskService) Update(id int, updatedTask model.Task) (model.Task, bool) {
	return s.taskRepo.Update(id, updatedTask)
}

func (s *TaskService) Delete(id int) bool {
	return s.taskRepo.Delete(id)
}

func (s *TaskService) DeleteAll() {
	s.taskRepo.DeleteAll()
}
