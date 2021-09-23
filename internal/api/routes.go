package api

// UnprotectedRoutes will set up all the unprotected routes
func (s *Server) UnprotectedRoutes() {
	// HealthCheck
	s.Echo.GET("/", s.HealthCheck)
}

// ProtectedRoutes will set up all the protected routes
func (s *Server) ProtectedRoutes() {
	// Group Routes
	v1 := s.Echo.Group("v1/")

	// Book
	v1.PUT("/book", s.UpdateBook)
	v1.POST("/book", s.CreateBook)
	v1.DELETE("/book", s.DeleteBook)
	v1.GET("/book", s.GetBook)

	// Books
	v1.GET("/books", s.GetBooks)

	// Token
	v1.POST("/token/renew", s.RenewToken)

	// User
	v1.POST("/user/login", s.Login)
	v1.POST("/user/logout", s.LogOut)
	v1.POST("/user/signup", s.SignUp)
}
