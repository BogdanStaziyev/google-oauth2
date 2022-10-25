package main

import (
	application "exemple_oauth"
	"exemple_oauth/config"
)

func main() {
	cfg := config.NewConfig()
	application.Start(cfg)
}
