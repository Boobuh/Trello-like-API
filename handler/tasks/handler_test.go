package tasks

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/Boobuh/golang-school-project/handler/tasks/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetAllTasks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       []dal.Task
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
					service.EXPECT().GetTasks().Return([]dal.Task{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/tasks/",
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
					service.EXPECT().GetTasks().Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/tasks/",
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
			router.HandleFunc("/tasks/", h.GetAllTasks)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_GetTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		//body       int
		method string
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
					service.EXPECT().GetTask(1, 1, 1).Return(&dal.ExtendedTask{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/1",
				//body:       0,
				method: http.MethodGet,
			},
			expected: expected{code: http.StatusOK},
		},
		{
			name: "failed",
			fields: fields{
				logger: log.Default(),
				service: func() Service {
					service := mocks.NewMockService(ctrl)
					service.EXPECT().GetTask(0, 0, 0).Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0/tasks/0",
				//body:       0,
				method: http.MethodGet,
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}", h.GetTask)

			recorder := httptest.NewRecorder()
			//body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, nil)
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_CreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       dal.Task
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
					service.EXPECT().CreateTask(&dal.Task{ColumnID: 1}).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/",
				body: dal.Task{
					ColumnID: 1,
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
					service.EXPECT().CreateTask(&dal.Task{ColumnID: 1}).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/1/tasks/",
				body: dal.Task{
					ColumnID: 1,
				},
				method: http.MethodPost,
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/", h.CreateTask)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_DeleteTask(t *testing.T) {
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
					service.EXPECT().DeleteTask(1, 1, 1).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/1",
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
					service.EXPECT().DeleteTask(0, 0, 0).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0/tasks/0",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}", h.DeleteTask)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_UpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       dal.Task
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
					service.EXPECT().UpdateTask(&dal.Task{ID: 1, ColumnID: 1, Name: "one_default"}).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/1",
				body: dal.Task{
					ID:       1,
					ColumnID: 1,
					Name:     "one_default",
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
					service.EXPECT().UpdateTask(&dal.Task{ID: 0, ColumnID: 0, Name: "one"}).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0/tasks/0",
				body: dal.Task{
					ID:       0,
					ColumnID: 0,
					Name:     "one",
				},
				method: http.MethodPut,
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}", h.UpdateTask)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_GetAllByColumnID(t *testing.T) {
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
					service.EXPECT().GetAllByColumnID(1).Return([]dal.ExtendedTask{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/",
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
					service.EXPECT().GetAllByColumnID(0).Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0/tasks/",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/", h.GetAllByColumnID)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}
