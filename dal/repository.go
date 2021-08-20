package dal

//go:generate   $GOPATH/bin/mockgen -package mocks -destination=mocks/mock_repository.go -package=mocks github.com/Boobuh/golang-school-project/dal Repository

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
	//-----------------------------------------//
	GetProjects() ([]Project, error)
	GetProject(id int) (*ExtendedProjectEntities, error)
	UpdateProject(project *Project) error
	CreateProject(project *Project) (*Project, error)
	DeleteProject(id int) error
	//-----------------------------------------//
	GetColumns() ([]Column, error)
	GetColumn(id int) (*ExtendedColumn, error)
	UpdateColumn(updatedColumn *Column) error
	CreateColumn(column *Column) error
	DeleteColumn(projectID, columnID int) error
	//-----------------------------------------//
	GetTasks() ([]Task, error)
	GetTask(id int) (*ExtendedTask, error)
	UpdateTask(updatedTask *Task) error
	CreateTask(task *Task) error
	DeleteTask(projectID, columnID, taskID int) error
	//-----------------------------------------//
	GetComments() ([]Comment, error)
	GetComment(id int) (*Comment, error)
	UpdateComment(updatedComment *Comment) error
	CreateComment(comment *Comment) error
	DeleteComment(projectID, columnID, taskID, commentID int) error
	//-----------------------------------------//
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository() Repository {
	db, err := gorm.Open(sqlite.Open("projects.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&Column{})
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&Comment{})
	return &RepositoryImpl{db: db}
}

//----------------------------------------------------------------------------------------//

type ExtendedProjectEntities struct {
	Project
	Columns []ExtendedColumn
}

type ExtendedColumn struct {
	Column
	Tasks []ExtendedTask
}

type ExtendedTask struct {
	Task
	Comments []Comment
}

func (r *RepositoryImpl) GetProjects() ([]Project, error) {
	var projects []Project
	err := r.db.Find(&projects).Error
	return projects, err

}

func (r *RepositoryImpl) GetProject(id int) (*ExtendedProjectEntities, error) {
	//TODO: create projectEntities wich has no project but all columns tasks and comments that belongs to this project
	var extendedProject ExtendedProjectEntities
	var project *Project
	err := r.db.First(&project, id).Error
	if err != nil {
		fmt.Printf("error retreiving project by id:%s\n", err.Error())
		return nil, err
	}
	extendedProject.Project = *project
	var columns []Column
	err = r.db.Find(&columns, "project_id = ?", project.ID).Error
	if err != nil {
		fmt.Printf("error finding columns by project_id:%s\n", err.Error())
		return nil, err
	}

	var extendedColumns []ExtendedColumn
	for _, column := range columns {
		var extColumn ExtendedColumn
		extColumn.Column = column
		var tasks []Task
		err := r.db.Find(&tasks, "column_id = ?", column.ID).Error
		if err != nil {
			fmt.Printf("error finding tasks by column_id:%s\n", err.Error())
			return nil, err
		}
		var extTasks []ExtendedTask
		for _, task := range tasks {
			var extTask ExtendedTask
			extTask.Task = task
			var comments []Comment
			err := r.db.Find(&comments, "task_id = ?", task.ID).Error
			if err != nil {
				fmt.Printf("error finding comments by task_id:%s\n", err.Error())
				return nil, err
			}
			extTask.Comments = comments

			extTasks = append(extTasks, extTask)

		}
		extColumn.Tasks = extTasks
		extendedColumns = append(extendedColumns, extColumn)
	}
	extendedProject.Columns = extendedColumns
	return &extendedProject, nil

}

func (r *RepositoryImpl) UpdateProject(updatedProject *Project) error {
	return r.db.Model(&updatedProject).Updates(updatedProject).Error
}

func (r *RepositoryImpl) CreateProject(project *Project) (*Project, error) {
	err := r.db.Create(&project).Error
	if err != nil {
		return nil, err
	}

	return project, err
}

func (r *RepositoryImpl) DeleteProject(id int) error {
	project := &Project{ID: id}
	return r.db.Delete(&project).Error
}

//----------------------------------------------------------------------------------------//

func (r *RepositoryImpl) GetColumns() ([]Column, error) {
	var columns []Column
	err := r.db.Find(&columns).Error
	return columns, err

}

func (r *RepositoryImpl) GetColumn(id int) (*ExtendedColumn, error) {

	var column Column
	err := r.db.First(&column, id).Error
	if err != nil {
		return nil, err
	}
	var extColumn ExtendedColumn
	extColumn.Column = column
	var tasks []Task
	err = r.db.Find(&tasks, "column_id = ?", column.ID).Error
	if err != nil {
		fmt.Printf("error finding tasks by column_id:%s\n", err.Error())
		return nil, err
	}
	var extTasks []ExtendedTask
	for _, task := range tasks {
		var extTask ExtendedTask
		extTask.Task = task
		var comments []Comment
		err := r.db.Find(&comments, "task_id = ?", task.ID).Error
		if err != nil {
			fmt.Printf("error finding comments by task_id:%s\n", err.Error())
			return nil, err
		}
		extTask.Comments = comments

		extTasks = append(extTasks, extTask)

	}
	extColumn.Tasks = extTasks

	return &extColumn, nil

}

func (r *RepositoryImpl) UpdateColumn(updatedColumn *Column) error {
	return r.db.Save(updatedColumn).Error
}

func (r *RepositoryImpl) CreateColumn(column *Column) error {
	return r.db.Create(column).Error
}

func (r *RepositoryImpl) DeleteColumn(projectID, columnID int) error {
	column := &Column{ID: columnID, ProjectID: projectID}
	return r.db.Delete(&column).Error
}

//----------------------------------------------------------------------------------------//

func (r *RepositoryImpl) GetTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err

}

func (r *RepositoryImpl) GetTask(id int) (*ExtendedTask, error) {
	var task Task
	err := r.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	var extTask ExtendedTask
	extTask.Task = task
	var comments []Comment
	err = r.db.Find(&comments, "task_id = ?", task.ID).Error
	if err != nil {
		fmt.Printf("error finding comments by task_id:%s\n", err.Error())
		return nil, err
	}

	extTask.Comments = comments

	return &extTask, nil
}

func (r *RepositoryImpl) UpdateTask(updatedTask *Task) error {
	return r.db.Save(updatedTask).Error
}

func (r *RepositoryImpl) CreateTask(task *Task) error {
	return r.db.Create(task).Error
}

func (r *RepositoryImpl) DeleteTask(projectID, columnID, taskID int) error {
	task := &Task{ID: taskID, ColumnID: columnID}
	return r.db.Delete(&task).Error
}

//----------------------------------------------------------------------------------------//

func (r *RepositoryImpl) GetComments() ([]Comment, error) {
	var comments []Comment
	err := r.db.Find(&comments).Error
	return comments, err
}

func (r *RepositoryImpl) GetComment(id int) (*Comment, error) {
	var comment *Comment
	err := r.db.Find(&comment, "id = ?", id).Error
	return comment, err
}

func (r *RepositoryImpl) UpdateComment(updatedComment *Comment) error {
	err := r.db.Save(updatedComment).Error
	return err
}

func (r *RepositoryImpl) CreateComment(comment *Comment) error {
	return r.db.Create(comment).Error
}

func (r *RepositoryImpl) DeleteComment(projectID, columnID, taskID, commentID int) error {
	comment := &Comment{ID: commentID, TaskID: taskID}
	return r.db.Delete(&comment).Error
}

//----------------------------------------------------------------------------------------//
