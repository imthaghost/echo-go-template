package api

import (
	"github.com/imthaghost/echo-template/internal/datastore"
	"github.com/labstack/echo"
)

// Server ...
type Server struct {
	// Echo server
	Echo *echo.Echo
	// inject other services here
	// data store
	db datastore.Service

}

// NewRESTServer creates a new HTTP server with injected services
func NewRESTServer(db *datastore.Service) *Server {
	return &Server{
		Echo:          echo.New(),
	}
}

// Start initializes the HTTP server
func (s *Server) Start() {
	server := s.Echo
	// register routes
	s.UnprotectedRoutes()
	s.ProtectedRoutes()
	// err  := server.Start()
	// TODO dont use Logger Fatal in Prod - don't panic
	s.Echo.Logger.Fatal(server.Start(":"+"8080"))
}

// Close gracefully shuts down server and closes all connections
func (s *Server) Close() (err error) {
	return s.Echo.Close()
}


