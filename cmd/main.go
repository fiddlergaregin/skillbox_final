package main

import (
	"log"
	"skillbox_final/config"
	"skillbox_final/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
