package comments

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
	_, err := c.repo.GetComment(comment.ID)
	if err != nil {

		return err
	}
	err = c.repo.UpdateComment(comment)
	return err
}

func (c *UseCase) GetAllByTaskID(taskID int) ([]dal.Comment, error) {
	task, err := c.repo.GetTask(taskID)
	if err != nil {
		return nil, err
	}
	return task.Comments, nil
}
