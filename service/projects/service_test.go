package projects

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/Boobuh/golang-school-project/dal/mocks"
)

//func TestUseCase_UpdateProject(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	type fields struct {
//		repo   dal.Repository
//		logger *log.Logger
//	}
//
//	type args struct {
//		updatedProject *dal.Project
//	}
//
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//		args    args
//	}{
//		// TODO: Add test cases.
//		{
//			name: "success",
//			fields: fields{
//				logger: log.Default(),
//				repo: func() dal.Repository {
//					repo := mocks.NewMockRepository(ctrl)
//					repo.EXPECT().UpdateProject(dal.Project{}).Return(nil).Times(1)
//					return repo
//				}(),
//			},
//			wantErr: false,
//			args: args{
//				updatedProject: nil,
//			},
//		},
//		{
//			name: "fail",
//			fields: fields{
//				logger: log.Default(),
//				repo: func() dal.Repository {
//					repo := mocks.NewMockRepository(ctrl)
//					repo.EXPECT().UpdateProject(nil).Return(errors.New("failed")).Times(1)
//					return repo
//				}(),
//			},
//			wantErr: true,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &UseCase{
//				repo:   tt.fields.repo,
//				logger: tt.fields.logger,
//			}
//			if err := c.UpdateProject(tt.args.updatedProject); (err != nil) != tt.wantErr {
//				t.Errorf("UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func TestUseCase_GetProjects(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		want    []dal.Project
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetProjects().Return([]dal.Project{}, nil).Times(1)
					return repo
				}(),
			},
			want:    []dal.Project{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetProjects().Return(nil, errors.New("failed")).Times(1)
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
			got, err := c.GetProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProjects() got = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestUseCase_GetProject(t *testing.T) {
//	type fields struct {
//		repo   dal.Repository
//		logger *log.Logger
//	}
//	type args struct {
//		id int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *dal.ExtendedProjectEntities
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &UseCase{
//				repo:   tt.fields.repo,
//				logger: tt.fields.logger,
//			}
//			got, err := c.GetProject(tt.args.id)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetProject() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetProject() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestUseCase_CreateProject(t *testing.T) {
//	type fields struct {
//		repo   dal.Repository
//		logger *log.Logger
//	}
//	type args struct {
//		project *dal.Project
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &UseCase{
//				repo:   tt.fields.repo,
//				logger: tt.fields.logger,
//			}
//			if err := c.CreateProject(tt.args.project); (err != nil) != tt.wantErr {
//				t.Errorf("CreateProject() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestUseCase_DeleteProject(t *testing.T) {
//	type fields struct {
//		repo   dal.Repository
//		logger *log.Logger
//	}
//	type args struct {
//		id int
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			c := &UseCase{
//				repo:   tt.fields.repo,
//				logger: tt.fields.logger,
//			}
//			if err := c.DeleteProject(tt.args.id); (err != nil) != tt.wantErr {
//				t.Errorf("DeleteProject() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func TestUseCase_UpdateProject1(t *testing.T) {
	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		updatedProject *dal.Project
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			if err := c.UpdateProject(tt.args.updatedProject); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProject() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
