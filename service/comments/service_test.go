package comments

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/Boobuh/golang-school-project/dal/mocks"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/golang/mock/gomock"
)

func TestUseCase_GetComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		want    []dal.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetComments().Return([]dal.Comment{}, nil).Times(1)
					return repo
				}(),
			},
			want:    []dal.Comment{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetComments().Return(nil, errors.New("failed")).Times(1)
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
			got, err := c.GetComments()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetComments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetComments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetComment(t *testing.T) {
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
		commentID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dal.Comment
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetComment(1).Return(&dal.Comment{}, nil).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 1,
				columnID:  1,
				taskID:    1,
				commentID: 1,
			},
			want:    &dal.Comment{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetComment(0).Return(nil, errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 0,
				columnID:  0,
				taskID:    0,
				commentID: 0,
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
			got, err := c.GetComment(tt.args.projectID, tt.args.columnID, tt.args.taskID, tt.args.commentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetComment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_CreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		comment *dal.Comment
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
					repo.EXPECT().CreateComment(&dal.Comment{Description: "", TaskID: 1, ID: 0}).Return(nil).Times(1)
					return repo
				}(),
			},
			args: args{
				comment: &dal.Comment{
					Description: "success",
					TaskID:      1,
					ID:          0,
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
					repo.EXPECT().CreateComment(&dal.Comment{Description: "", TaskID: 0, ID: 0}).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				comment: &dal.Comment{
					Description: "fail",
					TaskID:      0,
					ID:          0,
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
			if err := c.CreateComment(tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("CreateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_DeleteComment(t *testing.T) {
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
		commentID int
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
					repo.EXPECT().DeleteComment(1, 1, 1, 1).Return(nil).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 1,
				columnID:  1,
				taskID:    1,
				commentID: 1,
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().DeleteComment(0, 0, 0, 0).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 0,
				columnID:  0,
				taskID:    0,
				commentID: 0,
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
			if err := c.DeleteComment(tt.args.projectID, tt.args.columnID, tt.args.taskID, tt.args.commentID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_UpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		comment *dal.Comment
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
					repo.EXPECT().GetComment(1).Return(nil, nil).Times(1)
					repo.EXPECT().UpdateComment(&dal.Comment{ID: 1}).Return(nil).Times(1)
					return repo
				}(),
			},
			wantErr: false,
			args: args{
				comment: &dal.Comment{
					ID: 1,
				},
			},
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetComment(0).Return(nil, nil).Times(1)
					repo.EXPECT().UpdateComment(&dal.Comment{}).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			wantErr: true,
			args: args{
				comment: &dal.Comment{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			if err := c.UpdateComment(tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("UpdateComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetAllByTaskID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		taskID int
	}
	var secondCall []dal.Comment
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dal.Comment
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
					secondCall = dal.ExtendedTask{}.Comments
					return repo
				}(),
			},
			want:    secondCall,
			wantErr: false,
			args: args{
				taskID: 1,
			},
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetTask(0).Return(&dal.ExtendedTask{}, errors.New("failed")).Times(1)
					secondCall = dal.ExtendedTask{}.Comments
					return repo
				}(),
			},
			want:    secondCall,
			wantErr: true,
			args: args{
				taskID: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			got, err := c.GetAllByTaskID(tt.args.taskID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByTaskID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllByTaskID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
