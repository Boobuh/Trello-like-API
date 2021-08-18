package columns

import (
	"errors"
	"log"

	"github.com/Boobuh/golang-school-project/dal"
)

func NewUseCase(repo dal.Repository, logger *log.Logger) *UseCase {
	return &UseCase{repo: repo, logger: logger}
}

type UseCase struct {
	repo   dal.Repository
	logger *log.Logger
}

type Service interface {
	GetColumns() ([]dal.Column, error)
	GetProjectColumn(projectID, columnID int) (*dal.ExtendedColumn, error)
	CreateColumn(column *dal.Column) error
	DeleteColumn(projectID, columnID int) error
	UpdateColumn(updatedColumn *dal.Column) error
	GetAllByProjectID(projectID int) ([]dal.ExtendedColumn, error)
}

//=======================================================================================//

func (c *UseCase) GetColumns() ([]dal.Column, error) {
	return c.repo.GetColumns()
}

func (c *UseCase) GetColumn(id int) (*dal.ExtendedColumn, error) {
	return c.repo.GetColumn(id)
}
func (c *UseCase) GetProjectColumn(projectID, columnID int) (*dal.ExtendedColumn, error) {

	column, err := c.repo.GetColumn(columnID)
	if err != nil {
		return nil, err

	}
	if column.ProjectID != projectID {
		return nil, errors.New("projectID is missmatched with column")
	}
	return column, nil
}

func (c *UseCase) CreateColumn(column *dal.Column) error {
	return c.repo.CreateColumn(column)
}

func (c *UseCase) DeleteColumn(projectID, columnID int) error {
	return c.repo.DeleteColumn(projectID, columnID)
}

func (c *UseCase) UpdateColumn(updatedColumn *dal.Column) error {
	_, err := c.repo.GetColumn(updatedColumn.ID)
	if err != nil {

		return err
	}
	err = c.repo.UpdateColumn(updatedColumn)
	return err
}

func (c *UseCase) GetAllByProjectID(projectID int) ([]dal.ExtendedColumn, error) {
	project, err := c.repo.GetProject(projectID)
	if err != nil {
		return nil, err
	}
	return project.Columns, nil
}

//=======================================================================================//
