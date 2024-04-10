package controller

import "github.com/Daanooo/taski/internal/data"

type TaskController struct {
	repo *data.TaskRepository
}

func NewTaskController(repo *data.TaskRepository) *TaskController {
	return &TaskController{
		repo: repo,
	}
}

func (c TaskController) NewTask(description string) {
	task := data.Task{
		Description: description,
	}

	c.repo.Create(task)
}

func (c TaskController) GetById(id int) data.Task {
	return c.repo.GetById(id)
}
