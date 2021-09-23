package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// LogOut ...
func (s *Server) LogOut(c echo.Context) error {



	return c.NoContent(http.StatusOK)
}
