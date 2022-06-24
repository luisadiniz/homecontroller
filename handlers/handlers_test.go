package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type repositoryMock struct {
	lightbulbs map[string]bool
	err        error
}

func (l *repositoryMock) Get(ctx context.Context) (map[string]bool, error) {
	return l.lightbulbs, l.err
}

func (l *repositoryMock) GetById(ctx context.Context, name string) (bool, error) {
	return l.lightbulbs[name], l.err
}

func (l *repositoryMock) Create(ctx context.Context, name string, value bool) error {
	l.lightbulbs[name] = value
	return l.err
}

func (l *repositoryMock) Update(ctx context.Context, name string, value bool) error {
	l.lightbulbs[name] = value
	return l.err
}

func (l *repositoryMock) Delete(ctx context.Context, name string) error {
	delete(l.lightbulbs, name)
	return l.err
}

func TestHandleLightbulbs(t *testing.T) {
	type args struct {
		r    func() *http.Request
		repo func() Repository
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
		wantBody map[string]bool
	}{
		{
			name: "WhenGetFailReturnFailedDependency",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodGet, "/lightbulbs", nil)
				},
				repo: func() Repository {
					return &repositoryMock{
						err:        errors.New("Get Lightbulb -> failed to connect to database"),
						lightbulbs: map[string]bool{},
					}
				},
			},
			wantCode: http.StatusFailedDependency,
			wantBody: map[string]bool{},
		}, {
			name: "WhenGetWorksReturnOKWithDatabaseEntries",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodGet, "/lightbulbs", nil)
				},
				repo: func() Repository {
					return &repositoryMock{
						lightbulbs: map[string]bool{
							"sala":   true,
							"quarto": true,
						},
					}
				},
			},
			wantCode: http.StatusOK,
			wantBody: map[string]bool{
				"sala":   true,
				"quarto": true,
			},
		},
		{
			name: "WhenPostFailReturnBadRequest",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodPost, "/lightbulbs", nil)
				},
				repo: func() Repository {
					return &repositoryMock{
						err:        errors.New("Post Lightbulb -> failed to connect to database"),
						lightbulbs: map[string]bool{},
					}
				},
			},
			wantCode: http.StatusBadRequest,
			wantBody: map[string]bool{},
		},
		{
			name: "WhenPostFailReturnFailedDependency",
			args: args{
				r: func() *http.Request {
					body := Lightbulb{
						Name: "sala",
						On:   false,
					}
					var buf bytes.Buffer
					err := json.NewEncoder(&buf).Encode(&body)
					if err != nil {
						fmt.Println(err.Error())
					}
					request := httptest.NewRequest(http.MethodPost, "/lightbulbs", &buf)
					return request
				},
				repo: func() Repository {
					return &repositoryMock{
						err:        errors.New("Create Lightbulb -> failed to connect to database"),
						lightbulbs: map[string]bool{},
					}
				},
			},
			wantCode: http.StatusFailedDependency,
			wantBody: map[string]bool{},
		},
		{
			name: "WhenPostWorksReturnOKWithDatabaseEntries",
			args: args{
				r: func() *http.Request {
					body := Lightbulb{
						Name: "sala",
						On:   false,
					}
					var buf bytes.Buffer
					err := json.NewEncoder(&buf).Encode(&body)
					if err != nil {
						fmt.Println(err.Error())
					}
					request := httptest.NewRequest(http.MethodPost, "/lightbulbs", &buf)
					return request
				},
				repo: func() Repository {
					return &repositoryMock{
						lightbulbs: map[string]bool{},
					}
				},
			},
			wantCode: http.StatusOK,
			wantBody: map[string]bool{
				"sala": false,
			},
		},
		{
			name: "WhenPutFailReturnFailedDependency",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodPut, "/lightbulbs?name=sala&on=false", nil)
				},
				repo: func() Repository {
					return &repositoryMock{
						err:        errors.New("Switch Lightbulb -> failed to connect to database"),
						lightbulbs: map[string]bool{},
					}
				},
			},
			wantCode: http.StatusFailedDependency,
			wantBody: map[string]bool{},
		},
		{
			name: "WhenPutWorksReturnOKWithDatabaseEntries",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodPut, "/lightbulbs?name=sala&on=false", nil)
				},
				repo: func() Repository {
					return &repositoryMock{
						lightbulbs: map[string]bool{
							"sala": true,
						},
					}
				},
			},
			wantCode: http.StatusOK,
			wantBody: map[string]bool{
				"sala": false,
			},
		}, {
			name: "WhenDeleteFailReturnFailedDependency",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodDelete, "/lightbulbs?name=sala", nil)
				},
				repo: func() Repository {
					return &repositoryMock{
						err:        errors.New("Delete Lightbulb -> failed to connect to database"),
						lightbulbs: map[string]bool{},
					}
				},
			},
			wantCode: http.StatusFailedDependency,
			wantBody: map[string]bool{},
		},
		{
			name: "WhenDeleteWorksReturnOKWithDatabaseEntries",
			args: args{
				r: func() *http.Request {
					return httptest.NewRequest(http.MethodDelete, "/lightbulbs?name=sala", nil)
				},
				repo: func() Repository {
					return &repositoryMock{
						lightbulbs: map[string]bool{
							"sala":   true,
							"quarto": true,
						},
					}
				},
			},
			wantCode: http.StatusOK,
			wantBody: map[string]bool{
				"quarto": true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handler := HandleLightbulbs(tt.args.repo())

			handler(recorder, tt.args.r())

			if recorder.Code != tt.wantCode {
				t.Errorf("%v = %v, want code %v", tt.name, recorder.Code, tt.wantCode)
			}

			resp := map[string]bool{}

			json.NewDecoder(recorder.Body).Decode(&resp)

			if !reflect.DeepEqual(resp, tt.wantBody) {
				t.Errorf("%v = %v, want body %v", tt.name, resp, tt.wantBody)
			}
		})
	}
}
