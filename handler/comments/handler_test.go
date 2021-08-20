package comments

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/Boobuh/golang-school-project/handler/comments/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetAllComments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       []dal.Comment
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
					service.EXPECT().GetComments().Return([]dal.Comment{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/comments/",
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
					service.EXPECT().GetComments().Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/comments/",
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
			router.HandleFunc("/comments/", h.GetAllComments)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_GetComment(t *testing.T) {
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
					service.EXPECT().GetComment(1, 1, 1, 1).Return(&dal.Comment{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/1/comments/1",
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
					service.EXPECT().GetComment(0, 0, 0, 0).Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0/tasks/0/comments/0",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID}", h.GetComment)

			recorder := httptest.NewRecorder()
			//body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, nil)
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_CreateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       dal.Comment
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
					service.EXPECT().CreateComment(&dal.Comment{TaskID: 1}).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/1/comments/",
				body: dal.Comment{
					TaskID: 1,
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
					service.EXPECT().CreateComment(&dal.Comment{TaskID: 1}).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/1/tasks/1/comments/",
				body: dal.Comment{
					TaskID: 1,
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/", h.CreateComment)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_DeleteComment(t *testing.T) {
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
					service.EXPECT().DeleteComment(1, 1, 1, 1).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/1/comments/1",
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
					service.EXPECT().DeleteComment(0, 0, 0, 0).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0/tasks/0/comments/0",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID}", h.DeleteComment)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_UpdateComment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       dal.Comment
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
					service.EXPECT().UpdateComment(&dal.Comment{ID: 1, TaskID: 1, Description: "one_default"}).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/1/comments/1",
				body: dal.Comment{
					ID:          1,
					TaskID:      1,
					Description: "one_default",
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
					service.EXPECT().UpdateComment(&dal.Comment{ID: 0, TaskID: 0, Description: "one"}).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0/tasks/0/comments/0",
				body: dal.Comment{
					ID:          0,
					TaskID:      0,
					Description: "one",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID}", h.UpdateComment)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_GetAllByTaskID(t *testing.T) {
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
					service.EXPECT().GetAllByTaskID(1).Return([]dal.Comment{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1/tasks/1/comments/",
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
					service.EXPECT().GetAllByTaskID(0).Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0/tasks/0/comments/",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/", h.GetAllByTaskID)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}
