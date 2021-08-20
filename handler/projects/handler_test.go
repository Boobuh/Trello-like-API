package projects

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/Boobuh/golang-school-project/dal"

	"github.com/golang/mock/gomock"

	"github.com/Boobuh/golang-school-project/handler/projects/mocks"
)

func TestHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       int
		method     string
	}

	type expected struct {
		code int
		body string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().GetProject(1).Return(&dal.ExtendedProjectEntities{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1",
				body:       1,
				method:     http.MethodGet,
			},
			expected: expected{code: http.StatusOK},
		},
		{
			name: "failed",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().GetProject(1).Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1",
				body:       0,
				method:     http.MethodGet,
			},
			expected: expected{code: http.StatusBadRequest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				logger:  tt.fields.logger,
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/projects/{id}", h.Get)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       dal.Project
		method     string
	}

	type expected struct {
		code int
		body string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().CreateProject(&dal.Project{ID: 1, Name: "one", Description: "success"}).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/",
				body: dal.Project{
					ID:          1,
					Name:        "one",
					Description: "success",
				},
				method: http.MethodPost,
			},
			expected: expected{code: http.StatusCreated},
		},
		{
			name: "failed",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().CreateProject(&dal.Project{}).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/",
				body:       dal.Project{},
				method:     http.MethodPost,
			},
			expected: expected{code: http.StatusBadRequest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				logger:  tt.fields.logger,
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/projects/", h.Create)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       int
		method     string
	}

	type expected struct {
		code int
		body string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().DeleteProject(1).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1",
				body:       1,
				method:     http.MethodDelete,
			},
			expected: expected{code: http.StatusNoContent},
		},
		{
			name: "failed",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().DeleteProject(1).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1",
				body:       0,
				method:     http.MethodDelete,
			},
			expected: expected{code: http.StatusBadRequest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				logger:  tt.fields.logger,
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/projects/{id}", h.Delete)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       dal.Project
		method     string
	}

	type expected struct {
		code int
		body string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().UpdateProject(&dal.Project{ID: 1, Name: "one", Description: "success"}).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1",
				body: dal.Project{
					ID:          1,
					Name:        "one",
					Description: "success",
				},
				method: http.MethodPut,
			},
			expected: expected{code: http.StatusOK},
		},
		{
			name: "failed",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().UpdateProject(&dal.Project{ID: 1, Name: "one", Description: "failed"}).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1",
				body:       dal.Project{ID: 1, Name: "one", Description: "failed"},
				method:     http.MethodPut,
			},
			expected: expected{code: http.StatusBadRequest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				logger:  tt.fields.logger,
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/projects/{id}", h.Update)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       io.Reader
		method     string
	}

	type expected struct {
		code int
		body string
	}

	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().GetProjects().Return([]dal.Project{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/",
				body:       nil,
				method:     http.MethodGet,
			},
			expected: expected{code: http.StatusOK},
		},
		{
			name: "failed",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().GetProjects().Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/",
				body:       nil,
				method:     http.MethodGet,
			},
			expected: expected{code: http.StatusBadRequest},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				logger:  tt.fields.logger,
				service: tt.fields.service,
			}
			router := mux.NewRouter()
			router.HandleFunc("/projects/", h.GetAll)

			recorder := httptest.NewRecorder()
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, tt.args.body)
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}
