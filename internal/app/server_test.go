package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francois-roseberry/go-todo-app/internal/task"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
	"github.com/yosssi/gohtml"
)

func TestBaseRoute(t *testing.T) {
	s := NewServer(task.NewApp(), 6789)
	req := httptest.NewRequest(http.MethodGet, BaseRoute, nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	snaps.MatchSnapshot(t, gohtml.Format(rec.Body.String()))
}

func TestPutStatus(t *testing.T) {
	s := NewServer(task.NewApp(), 6789)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s?locked=true", StatusRoute), nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	snaps.MatchSnapshot(t, gohtml.Format(rec.Body.String()))
}

func TestCreateTask(t *testing.T) {
	s := NewServer(task.NewApp(), 6789)
	req := httptest.NewRequest(http.MethodPost, TaskListRoute, nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	snaps.MatchSnapshot(t, gohtml.Format(rec.Body.String()))
}

func TestDeleteTask(t *testing.T) {
	s := NewServer(task.NewApp(), 6789)
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s?id=1", TaskListRoute), nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Result().StatusCode)
}
