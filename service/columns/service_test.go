package columns

import (
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/Boobuh/golang-school-project/dal/mocks"

	"github.com/golang/mock/gomock"

	"github.com/Boobuh/golang-school-project/dal"
)

func TestUseCase_GetColumns(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		want    []dal.Column
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetColumns().Return([]dal.Column{}, nil).Times(1)
					return repo
				}(),
			},
			want:    []dal.Column{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetColumns().Return(nil, errors.New("failed")).Times(1)
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
			got, err := c.GetColumns()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetColumns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetColumns() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dal.ExtendedColumn
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
					return repo
				}(),
			},
			args: args{
				id: 1,
			},
			want:    &dal.ExtendedColumn{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetColumn(1).Return(nil, errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				id: 1,
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
			got, err := c.GetColumn(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetColumn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetColumn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_GetProjectColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		projectID int
		columnID  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dal.ExtendedColumn
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
					return repo
				}(),
			},
			args: args{
				projectID: 0,
				columnID:  1,
			},
			want:    &dal.ExtendedColumn{},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetColumn(1).Return(nil, errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 1,
				columnID:  1,
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
			got, err := c.GetProjectColumn(tt.args.projectID, tt.args.columnID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectColumn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProjectColumn() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_CreateColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		column *dal.Column
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
					repo.EXPECT().CreateColumn(&dal.Column{1, "", 1, 1, ""}).Return(nil).Times(1)
					return repo
				}(),
			},
			args: args{
				column: &dal.Column{
					ID:        1,
					Name:      "",
					ProjectID: 1,
					OrderNum:  1,
					Status:    "",
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
					repo.EXPECT().CreateColumn(&dal.Column{0, "", 0, 0, ""}).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				column: &dal.Column{
					ID:        0,
					Name:      "",
					ProjectID: 0,
					OrderNum:  0,
					Status:    "",
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
			if err := c.CreateColumn(tt.args.column); (err != nil) != tt.wantErr {
				t.Errorf("CreateColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_DeleteColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		projectID int
		columnID  int
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
					repo.EXPECT().DeleteColumn(1, 1).Return(nil).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 1,
				columnID:  1,
			},
			wantErr: false,
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().DeleteColumn(0, 0).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			args: args{
				projectID: 0,
				columnID:  0,
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
			if err := c.DeleteColumn(tt.args.projectID, tt.args.columnID); (err != nil) != tt.wantErr {
				t.Errorf("DeleteColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_UpdateColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		updatedColumn *dal.Column
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
					repo.EXPECT().GetColumn(1).Return(nil, nil).Times(1)
					repo.EXPECT().UpdateColumn(&dal.Column{ID: 1, Name: "success", ProjectID: 1}).Return(nil).Times(1)
					return repo
				}(),
			},
			wantErr: false,
			args: args{
				updatedColumn: &dal.Column{
					ID:        1,
					Name:      "success",
					ProjectID: 1,
				},
			},
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetColumn(0).Return(nil, nil).Times(1)
					repo.EXPECT().UpdateColumn(&dal.Column{}).Return(errors.New("failed")).Times(1)
					return repo
				}(),
			},
			wantErr: true,
			args: args{
				updatedColumn: &dal.Column{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			if err := c.UpdateColumn(tt.args.updatedColumn); (err != nil) != tt.wantErr {
				t.Errorf("UpdateColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUseCase_GetAllByProjectID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		repo   dal.Repository
		logger *log.Logger
	}
	type args struct {
		projectID int
	}
	var secondCall []dal.ExtendedColumn

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []dal.ExtendedColumn
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetProject(1).Return(&dal.ExtendedProjectEntities{}, nil).Times(1)
					secondCall = dal.ExtendedProjectEntities{}.Columns
					return repo
				}(),
			},
			want:    secondCall,
			wantErr: false,
			args: args{
				projectID: 1,
			},
		},
		{
			name: "fail",
			fields: fields{
				logger: log.Default(),
				repo: func() dal.Repository {
					repo := mocks.NewMockRepository(ctrl)
					repo.EXPECT().GetProject(0).Return(&dal.ExtendedProjectEntities{}, errors.New("failed")).Times(1)
					secondCall = dal.ExtendedProjectEntities{}.Columns
					return repo
				}(),
			},
			want:    secondCall,
			wantErr: true,
			args: args{
				projectID: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UseCase{
				repo:   tt.fields.repo,
				logger: tt.fields.logger,
			}
			got, err := c.GetAllByProjectID(tt.args.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllByProjectID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllByProjectID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
