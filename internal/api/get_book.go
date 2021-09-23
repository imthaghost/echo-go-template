package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// GetBook get a book by given ID
func (s *Server) GetBook(c echo.Context) error {
	store := s.db

	store.Get()

	return c.NoContent(http.StatusOK)
}
