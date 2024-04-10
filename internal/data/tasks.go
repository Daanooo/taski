package data

import (
	"database/sql"
	"log"
)

type Task struct {
	ID          int
	Description string
	Completed   int
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t TaskRepository) Create(task Task) error {
	stmt, err := t.db.Prepare("insert into tasks (description, completed) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Description, task.Completed)
	if err != nil {
		return err
	}

	return nil
}

func (t TaskRepository) GetAll() ([]Task, error) {
	query := "select * from tasks"
	var tasks []Task

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task Task

		if err := rows.Scan(&task.ID, &task.Description, &task.Completed); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (t TaskRepository) GetById(id int) Task {
	query := "select * from tasks where id = ?"
	var task Task

	err := t.db.QueryRow(query, id).Scan(&task.ID, &task.Description, &task.Completed)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No task found for id %d", id)
		} else {
			log.Fatal(err)
		}
	}

	return task
}
