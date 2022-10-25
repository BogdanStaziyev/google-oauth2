package routes

import (
	"exemple_oauth/server"
	"exemple_oauth/server/handlers"
)

func ConfigureRoutes(server *server.Server) {
	oauthHandler := handlers.NewOauthHandler(server)

	r := server.Echo.Group("")

	r.GET("/auth/google/login", oauthHandler.GetInfo)
	r.GET("/auth/google/callback", oauthHandler.CallBack)
}
