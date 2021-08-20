package columns

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/Boobuh/golang-school-project/dal/mocks"
	"github.com/Boobuh/golang-school-project/handler/columns/mocks"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHandler_GetAllColumns(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       []dal.Column
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
					service.EXPECT().GetColumns().Return([]dal.Column{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/columns/",
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
					service.EXPECT().GetColumns().Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/columns/",
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
			router.HandleFunc("/columns/", h.GetAllColumns)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_GetColumn(t *testing.T) {
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
					service.EXPECT().GetProjectColumn(1, 1).Return(&dal.ExtendedColumn{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1",
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
					service.EXPECT().GetProjectColumn(0, 0).Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}", h.GetColumn)

			recorder := httptest.NewRecorder()
			//body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, nil)
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_CreateColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       dal.Column
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
					service.EXPECT().CreateColumn(&dal.Column{ProjectID: 1}).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/",
				body: dal.Column{
					ProjectID: 1,
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
					service.EXPECT().CreateColumn(&dal.Column{ProjectID: 1}).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/",
				body: dal.Column{
					ProjectID: 1,
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
			router.HandleFunc("/projects/{projectID}/columns/", h.CreateColumn)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_DeleteColumn(t *testing.T) {
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
					service.EXPECT().DeleteColumn(1, 1).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1",
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
					service.EXPECT().DeleteColumn(0, 0).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}", h.DeleteColumn)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_UpdateColumn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		logger  *log.Logger
		service Service
	}
	type args struct {
		urlRequest string
		body       dal.Column
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
					service.EXPECT().UpdateColumn(&dal.Column{ID: 1, ProjectID: 1, Name: "one_default"}).Return(nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/1",
				body: dal.Column{
					ID:        1,
					ProjectID: 1,
					Name:      "one_default",
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
					service.EXPECT().UpdateColumn(&dal.Column{ID: 0, ProjectID: 0, Name: "one"}).Return(errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/0",
				body: dal.Column{
					ID:        0,
					ProjectID: 0,
					Name:      "one",
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
			router.HandleFunc("/projects/{projectID}/columns/{columnID}", h.UpdateColumn)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}

func TestHandler_GetAllByProjectID(t *testing.T) {
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
					service.EXPECT().GetAllByProjectID(1).Return([]dal.ExtendedColumn{}, nil).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/1/columns/",
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
					service.EXPECT().GetAllByProjectID(0).Return(nil, errors.New("failed")).Times(1)
					return service
				}(),
			},
			args: args{
				urlRequest: "/projects/0/columns/",
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
			router.HandleFunc("/projects/{projectID}/columns/", h.GetAllByProjectID)

			recorder := httptest.NewRecorder()
			body, err := json.Marshal(tt.args.body)
			req, err := http.NewRequest(tt.args.method, tt.args.urlRequest, bytes.NewReader(body))
			assert.NoError(t, err)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tt.expected.code, recorder.Code)
		})
	}
}
