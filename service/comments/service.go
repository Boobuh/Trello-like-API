package comments

import (
	"log"

	"github.com/Boobuh/golang-school-project/dal"
)

type Service interface {
	GetComments() ([]dal.Comment, error)
	GetComment(projectID, columnID, taskID, commentID int) (*dal.Comment, error)
	CreateComment(task *dal.Comment) error
	DeleteComment(projectID, columnID, taskID, commentID int) error
	UpdateComment(task *dal.Comment) error
}

type UseCase struct {
	repo   dal.Repository
	logger *log.Logger
}

func NewUseCase(repo dal.Repository, logger *log.Logger) *UseCase {
	return &UseCase{repo: repo, logger: logger}
}

func (c *UseCase) GetComments() ([]dal.Comment, error) {
	return c.repo.GetComments()
}

func (c *UseCase) GetComment(projectID, columnID, taskID, commentID int) (*dal.Comment, error) {
	return c.repo.GetComment(commentID)
}

func (c *UseCase) CreateComment(comment *dal.Comment) error {
	comment = &dal.Comment{TaskID: comment.TaskID}
	return c.repo.CreateComment(comment)
}

func (c *UseCase) DeleteComment(projectID, columnID, taskID, commentID int) error {
	return c.repo.DeleteComment(projectID, columnID, taskID, commentID)
}

func (c *UseCase) UpdateComment(comment *dal.Comment) error {
	oldComment, err := c.repo.GetComment(comment.ID)
	if err != nil {

		return err
	}
	err = c.repo.CreateComment(oldComment)
	return err
}
