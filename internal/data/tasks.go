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

func (t TaskRepository) Create(task Task) {
	stmt, err := t.db.Prepare("insert into tasks (description, completed) values (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.Description, task.Completed)
	if err != nil {
		log.Fatal(err)
	}
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
