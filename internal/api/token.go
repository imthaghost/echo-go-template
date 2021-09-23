package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// RenewToken renews JWT token
func (s *Server) RenewToken(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
