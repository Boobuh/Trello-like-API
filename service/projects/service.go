package projects

import (
	"fmt"
	"log"

	"github.com/Boobuh/golang-school-project/dal"
)

type Service interface {
	//--------------------------------------------------------------//
	GetProjects() ([]dal.Project, error)
	GetProject(id int) (*dal.ExtendedProjectEntities, error)
	CreateProject(project *dal.Project) error
	DeleteProject(id int) error
	UpdateProject(updatedProject *dal.Project) error
	//--------------------------------------------------------------//

}

type UseCase struct {
	repo   dal.Repository
	logger *log.Logger
}

func NewUseCase(repo dal.Repository, logger *log.Logger) *UseCase {
	return &UseCase{repo: repo, logger: logger}
}

//=======================================================================================//

func (c *UseCase) UpdateProject(updatedProject *dal.Project) error {
	_, err := c.repo.GetProject(updatedProject.ID)
	if err != nil {
		fmt.Println("project not found by id %s", err)
		return err
	}
	err = c.repo.UpdateProject(updatedProject)
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
	_, err = c.repo.GetProject(project.ID)
	if err != nil {
		return err
	}
	column := &dal.Column{ProjectID: project.ID, Name: project.Name + "_default"}
	return c.repo.CreateColumn(column)
}

func (c *UseCase) DeleteProject(id int) error {
	return c.repo.DeleteProject(id)
}
