package application

import (
	"exemple_oauth/config"
	"exemple_oauth/server"
	"exemple_oauth/server/routes"
	"log"
)

func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("PortAlready used")
	}

}
