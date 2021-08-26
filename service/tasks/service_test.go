package tasks

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/Boobuh/golang-school-project/dal/mocks"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/golang/mock/gomock"
)

func TestUseCase_GetTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		want    []dal.Task
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetTasks().Return([]dal.Task{}, nil).Times(1)
					return repo
				}(),
			},
			want:    []dal.Task{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetTasks().Return(nil, errors.New("failed")).Times(1)
					return repo
				}(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			got, err := c.GetTasks()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTasks() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		projectID int
		columnID  int
		taskID    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dal.ExtendedTask
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetTask(1).Return(&dal.ExtendedTask{}, nil).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 1,
				columnID:  1,
				taskID:    1,
			},
			want:    &dal.ExtendedTask{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetTask(0).Return(nil, errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 0,
				columnID:  0,
				taskID:    0,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			got, err := c.GetTask(tt.args.projectID, tt.args.columnID, tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		task *dal.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().CreateTask(&dal.Task{0, "", false, "", 1}).Return(nil).Times(1)
					return repo
				}(),
			},
			args: args{
				task: &dal.Task{
					ID:          1,
					Name:        "",
					Status:      true,
					Description: "success",
					ColumnID:    1,
				},
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().CreateTask(&dal.Task{0, "", false, "", 0}).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				task: &dal.Task{
					ID:          0,
					Name:        "",
					Status:      false,
					Description: "",
					ColumnID:    0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			if err := c.CreateTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("CreateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_DeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		projectID int
		columnID  int
		taskID    int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().DeleteTask(1, 1, 1).Return(nil).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 1,
				columnID:  1,
				taskID:    1,
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().DeleteTask(0, 0, 0).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 0,
				columnID:  0,
				taskID:    0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			if err := c.DeleteTask(tt.args.projectID, tt.args.columnID, tt.args.taskID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_UpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		task *dal.Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetTask(1).Return(nil, nil).Times(1)
					repo.EXPECT().CreateTask(&dal.Task{ID: 1, Name: "success", ColumnID: 1}).Return(nil).Times(1)
					return repo
				}(),
			},
			wantErr: false,
			args: args{
				task: &dal.Task{
					ID:       1,
					Name:     "success",
					ColumnID: 1,
				},
			},
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetTask(0).Return(nil, nil).Times(1)
					repo.EXPECT().CreateTask(&dal.Task{}).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			wantErr: true,
			args: args{
				task: &dal.Task{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			if err := c.UpdateTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetAllByColumnID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		columnID int
	}

	var secondCall []dal.ExtendedTask
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dal.ExtendedTask
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetColumn(1).Return(&dal.ExtendedColumn{}, nil).Times(1)
					secondCall = dal.ExtendedColumn{}.Tasks
					return repo
				}(),
			},
			want:    secondCall,
			wantErr: false,
			args: args{
				columnID: 1,
			},
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetColumn(0).Return(&dal.ExtendedColumn{}, errors.New("failed")).Times(1)
					secondCall = dal.ExtendedColumn{}.Tasks
					return repo
				}(),
			},
			want:    secondCall,
			wantErr: true,
			args: args{
				columnID: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			got, err := c.GetAllByColumnID(tt.args.columnID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByColumnID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllByColumnID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
