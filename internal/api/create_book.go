package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// CreateBook create a book
func (s *Server) CreateBook(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
