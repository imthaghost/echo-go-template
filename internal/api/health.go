package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// HealthCheck - 200 OK
func (s*Server) HealthCheck(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
