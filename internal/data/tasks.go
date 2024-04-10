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
	stmt, err := t.db.Prepare("select * from tasks where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	task := &Task{}
	for rows.Next() {
		if err := rows.Scan(&task.ID, &task.Description, &task.Completed); err != nil {
			log.Fatal(err)
		}
	}

	return *task
}
