package repository

import "learn-go/internal/model"

type TaskRepository struct {
	tasks  []model.Task
	nextId int
}

func (r *TaskRepository) Create(task model.Task) model.Task {
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
