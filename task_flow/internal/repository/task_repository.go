package repository

import (
	"context"
	"errors"
	"learn-go/internal/model"
	"log"

	"github.com/jackc/pgx/v5"
)

type TaskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) Create(task model.Task) model.Task {
	err := r.db.QueryRow(
		context.Background(),
		"INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id",
		task.Title, task.Description, task.Status,
	).Scan(&task.ID)

	if err != nil {
		log.Println("Error inserting task:", err)
		return model.Task{}
	}

	return task
}

func (r *TaskRepository) GetAll() []model.Task {
	rows, err := r.db.Query(context.Background(), "SELECT id, title, description, status FROM tasks")

	if err != nil {
		log.Println("Error querying tasks:", err)
		
		return []model.Task{}
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			log.Println("Error scanning task:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		log.Println("Rows iteration error:", err)
	}

	if tasks == nil {
		tasks = []model.Task{}
	}

	return tasks
}

func (r *TaskRepository) GetById(id int) (model.Task, bool) {
	var task model.Task
	err := r.db.QueryRow(context.Background(), "SELECT id, title, description, status FROM tasks WHERE id = $1", id).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Task{}, false
		}
		log.Println("Error getting task by ID:", err)
		return model.Task{}, false
	}
	return task, true
}

func (r *TaskRepository) Update(id int, updatedTask model.Task) (model.Task, bool) {
	var task model.Task
	err := r.db.QueryRow(context.Background(), "UPDATE tasks SET title = $1, description = $2, status = $3 WHERE id = $4 RETURNING id, title, description, status", updatedTask.Title, updatedTask.Description, updatedTask.Status, id).Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Task{}, false
		}
		log.Println("Error updating task:", err)
		return model.Task{}, false
	}
	return task, true
}

func (r *TaskRepository) Delete(id int) bool {
	return false
}

func (r *TaskRepository) DeleteAll() {
}
