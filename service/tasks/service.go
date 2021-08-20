package tasks

import (
	"log"

	"github.com/Boobuh/golang-school-project/dal"
)

type UseCase struct {
	repo   dal.Repository
	logger *log.Logger
}

func NewUseCase(repo dal.Repository, logger *log.Logger) *UseCase {
	return &UseCase{repo: repo, logger: logger}
}

func (c *UseCase) GetTasks() ([]dal.Task, error) {
	return c.repo.GetTasks()
}

func (c *UseCase) GetTask(projectID, columnID, taskID int) (*dal.ExtendedTask, error) {
	return c.repo.GetTask(taskID)
}

func (c *UseCase) CreateTask(task *dal.Task) error {
	task = &dal.Task{ColumnID: task.ColumnID}
	return c.repo.CreateTask(task)
}

func (c *UseCase) DeleteTask(projectID, columnID, taskID int) error {
	return c.repo.DeleteTask(projectID, columnID, taskID)
}

func (c *UseCase) UpdateTask(task *dal.Task) error {
	_, err := c.repo.GetTask(task.ID)
	if err != nil {

		return err
	}
	err = c.repo.CreateTask(task)
	return err
}

func (c *UseCase) GetAllByColumnID(columnID int) ([]dal.ExtendedTask, error) {
	column, err := c.repo.GetColumn(columnID)
	if err != nil {
		return nil, err
	}
	return column.Tasks, nil
}

//=======================================================================================//
