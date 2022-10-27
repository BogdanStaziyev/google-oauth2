package routes

import (
	"exemple_oauth/internal/server"
	"exemple_oauth/internal/server/handlers"
	"exemple_oauth/internal/server/validators"
)

func ConfigureRoutes(server *server.Server) {
	server.Echo.Validator = validators.NewValidator()
	oauthHandler := handlers.NewOauthHandler(server)
	r := server.Echo.Group("")
	r.GET("/auth/google/login", oauthHandler.GetInfo)
	r.GET("/auth/google/callback", oauthHandler.CallBack)
}
