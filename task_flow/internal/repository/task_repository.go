package repository

import "learn-go/internal/model"

type TaskRepository struct {
	tasks  []model.Task
	nextId int
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks:  make([]model.Task, 0),
		nextId: 1,
	}
}

func (r *TaskRepository) Create(task model.Task) model.Task {
	if r.nextId == 0 {
		r.nextId = 1
	}
	task.ID = r.nextId
	r.nextId++

	r.tasks = append(r.tasks, task)

	return task
}

func (r *TaskRepository) GetAll() []model.Task {
	return r.tasks
}

func (r *TaskRepository) GetById(id int) (model.Task, bool) {
	for _, task := range r.tasks {
		if task.ID == id {
			return task, true
		}
	}
	return model.Task{}, false
}

func (r *TaskRepository) Update(id int, updatedTask model.Task) (model.Task, bool) {
	for i, task := range r.tasks {
		if task.ID == id {
			updatedTask.ID = id
			r.tasks[i] = updatedTask
			return updatedTask, true
		}
	}
	return model.Task{}, false
}

func (r *TaskRepository) Delete(id int) bool {
	for i, task := range r.tasks {
		if task.ID == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return true
		}
	}
	return false
}

func (r *TaskRepository) DeleteAll() {
	r.tasks = []model.Task{}
}
