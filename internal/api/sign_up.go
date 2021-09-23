package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// SignUp ...
func (s *Server) SignUp(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

