package dal

import (
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
	//-----------------------------------------//
	GetProjects() ([]Project, error)
	GetProject(id int) (*ExtendedProjectEntities, error)
	UpdateProject(body *http.Request, updatedProject *Project) error
	CreateProject(project *Project) (*Project, error)
	DeleteProject(id int) error
	//-----------------------------------------//
	GetColumns() ([]Column, error)
	GetColumn(id int) (*Column, error)
	UpdateColumn(id int) error
	CreateColumn(column *Column) error
	DeleteColumn(id int) error
	//-----------------------------------------//
	GetTasks() ([]Task, error)
	GetTask(id int) (*Task, error)
	UpdateTask(id int) error
	CreateTask(task *Task) error
	DeleteTask(id int) error
	//-----------------------------------------//
	GetComments() ([]Comment, error)
	GetComment(id int) (*Comment, error)
	UpdateComment(id int) error
	CreateComment(comment *Comment) error
	DeleteComment(id int) error
	//-----------------------------------------//
	RunMigration()
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

func (r *RepositoryImpl) GetProjects() ([]Project, error) {
	var projects []Project
	err := r.db.Find(&projects).Error
	return projects, err

}

type ExtendedProjectEntities struct {
	Project
	Columns []extendedColumn
}

type extendedColumn struct {
	Column
	Tasks []extendedTask
}

type extendedTask struct {
	Task
	Comments []Comment
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

	var extendedColumns []extendedColumn
	for _, column := range columns {
		var extColumn extendedColumn
		extColumn.Column = column
		var tasks []Task
		err := r.db.Find(&tasks, "column_id = ?", column.ID).Error
		if err != nil {
			fmt.Printf("error finding tasks by column_id:%s\n", err.Error())
			return nil, err
		}
		var extTasks []extendedTask
		for _, task := range tasks {
			var extTask extendedTask
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

func (r *RepositoryImpl) UpdateProject(body *http.Request, id int) error {
	if id != 0 {
		return r.db.Save(&id).Error
	}
	return r.db.Save(&body).Error

}

func (r *RepositoryImpl) CreateProject(project *Project) (*Project, error) {
	err := r.db.Create(&project).Error

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

func (r *RepositoryImpl) GetColumn(id int) (*Column, error) {
	var column *Column
	err := r.db.First(&column, id).Error
	return column, err
}

func (r *RepositoryImpl) UpdateColumn(id int) error {
	err := r.db.Save(&id).Error
	return err
}

func (r *RepositoryImpl) CreateColumn(column *Column) error {
	return r.db.Create(&column).Error
}

func (r *RepositoryImpl) DeleteColumn(id int) error {
	column := &Column{ID: id}
	return r.db.Delete(&column).Error
}

//----------------------------------------------------------------------------------------//

func (r *RepositoryImpl) GetTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err

}

func (r *RepositoryImpl) GetTask(id int) (*Task, error) {
	var task *Task
	err := r.db.First(&task, id).Error
	return task, err
}

func (r *RepositoryImpl) UpdateTask(id int) error {
	err := r.db.Save(&id).Error
	return err
}

func (r *RepositoryImpl) CreateTask(task *Task) error {
	return r.db.Create(&task).Error
}

func (r *RepositoryImpl) DeleteTask(id int) error {
	task := &Task{ID: id}
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
	err := r.db.First(&comment, id).Error
	return comment, err
}

func (r *RepositoryImpl) UpdateComment(id int) error {
	err := r.db.Save(&id).Error
	return err
}

func (r *RepositoryImpl) CreateComment(comment *Comment) error {
	return r.db.Create(&comment).Error
}

func (r *RepositoryImpl) DeleteComment(id int) error {
	comment := &Comment{ID: id}
	return r.db.Delete(&comment).Error
}

//----------------------------------------------------------------------------------------//

func (r *RepositoryImpl) RunMigration() {
	r.db.AutoMigrate(&Project{})
	r.db.AutoMigrate(&Column{})
	r.db.AutoMigrate(&Task{})
	r.db.AutoMigrate(&Comment{})

}
