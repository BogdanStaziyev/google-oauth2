package server

import (
	"exemple_oauth/config"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
	Conf *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo: echo.New(),
		Conf: cfg,
	}
}

func (s Server) Start(addr string) error {
	return s.Echo.Start(":" + addr)
}
