package api

import (
	"github.com/labstack/echo"
	"net/http"
)

// DeleteBook [hard](https://www.quora.com/What-is-the-difference-between-soft-delete-and-hard-delete-in-SQL-Informatica-power-center-and-Informatica-cloud) deletes a book
func (s *Server) DeleteBook(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
