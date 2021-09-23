package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// GetBooks returns all of a users books
func (s *Server) GetBooks(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}