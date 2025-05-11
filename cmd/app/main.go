package main

import (
	"flag"
	"gpt-service/config"
	"gpt-service/internal/app"
	"log"
)

func main() {
	devmode := flag.Bool("dev", false, "Run server in development mode")
	flag.Parse()
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Run(cfg, *devmode)
}
