package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s?%s=true", StatusRoute, ParamLocked), nil)
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
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("%s?%s=1", TaskListRoute, ParamId), nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Result().StatusCode)
}

func TestReorderTaskList(t *testing.T) {
	s := NewServer(task.NewApp(), 6789)

	form := make(url.Values)
	form.Set("task-id", "2")
	form.Add("task-id", "1")
	form.Add("task-id", "3")
	form.Add("task-id", "4")
	form.Add("task-id", "5")
	form.Add("task-id", "6")
	form.Add("task-id", "7")
	form.Add("task-id", "8")
	form.Add("task-id", "9")
	req := httptest.NewRequest(http.MethodPut, TaskListRoute, strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Result().StatusCode)
}
