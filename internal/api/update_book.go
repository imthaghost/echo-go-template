package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// UpdateBook update a book
func (s *Server) UpdateBook(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

