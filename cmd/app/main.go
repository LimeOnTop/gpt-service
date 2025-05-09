package main

import (
	"flag"
	"gpt-service/config"
	"gpt-service/internal/app"
	"log"
	"os"
)

func main() {
	devmode := flag.Bool("dev", false, "Run server in development mode")
	flag.Parse()
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	// Устанавливаем переменную окружения
	err = os.Setenv("GPT_AUTHORIZATION_KEY", cfg.Token.AuthorizationKey)
	if err != nil {
		log.Fatalf("Failed to set env variable: %v", err)
	}

	// Проверяем, что переменная установлена
	_, ok := os.LookupEnv("GPT_AUTHORIZATION_KEY")
	if !ok {
		log.Fatal("GPT_AUTHORIZATION_KEY not found in environment")
	}

	app.Run(cfg, *devmode)
}
