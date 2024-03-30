package main

import (
	"context"

	"github.com/didasy/winmetrics/app"
	"github.com/didasy/winmetrics/config"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	cfg := config.New()
	err := cfg.Validate()
	if err != nil {
		log.Fatalf("Failed to validate configuration file: %v", err)
	}

	log.Infof("Starting server at %s", cfg.App.ListenAddress)

	err = app.Run(context.Background(), cfg, log)
	if err != nil {
		log.Fatalf("Server shutting down because of unexpected error: %v", err)
	}

	log.Infoln("Server shut down")
}
