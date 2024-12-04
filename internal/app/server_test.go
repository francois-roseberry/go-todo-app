package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/francois-roseberry/go-todo-app/internal/task"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/yosssi/gohtml"
)

func TestBaseRoute(t *testing.T) {
	s := NewServer(task.NewApp(), 6789)
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	snaps.MatchSnapshot(t, gohtml.Format(rec.Body.String()))
}

func TestPutStatusRoute(t *testing.T) {
	s := NewServer(task.NewApp(), 6789)
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("%s?locked=true", StatusRoute), nil)
	rec := httptest.NewRecorder()

	s.e.ServeHTTP(rec, req)
	snaps.MatchSnapshot(t, gohtml.Format(rec.Body.String()))
}
