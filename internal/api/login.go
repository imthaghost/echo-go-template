package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// Login ...
func (s *Server) Login(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
