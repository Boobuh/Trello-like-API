package projects

import (
	"log"
	"net/http"

	"github.com/Boobuh/golang-school-project/dal"
)

type Service interface {
	GetProjects() ([]dal.Project, error)
	GetProject(id int) (*dal.ExtendedProjectEntities, error)
	CreateProject(project *dal.Project) error
	DeleteProject(id int) error
	UpdateProject(r *http.Request, updatedProject dal.Project) error
}

type UseCase struct {
	repo   dal.Repository
	logger *log.Logger
}

func (c *UseCase) UpdateProject(r *http.Request, updatedProject dal.Project) error {
	err := c.repo.UpdateProject(r, updatedProject)
	return err
}

func (c *UseCase) GetProjects() ([]dal.Project, error) {
	return c.repo.GetProjects()
}

func (c *UseCase) GetProject(id int) (*dal.ExtendedProjectEntities, error) {
	return c.repo.GetProject(id)
}

func (c *UseCase) CreateProject(project *dal.Project) error {
	project, err := c.repo.CreateProject(project)
	if err != nil {
		return err
	}
	column := &dal.Column{ProjectID: project.ID, Name: "default"}
	return c.repo.CreateColumn(column)
}

func (c *UseCase) DeleteProject(id int) error {
	return c.repo.DeleteProject(id)
}

func NewUseCase(repo dal.Repository, logger *log.Logger) *UseCase {
	return &UseCase{repo: repo, logger: logger}
}
