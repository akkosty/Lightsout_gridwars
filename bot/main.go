package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Run starts the bot with proper error handling and cleanup
func Run() error {
	botToken := getEnv("BOT_TOKEN", "")
	if botToken == "" {
		return fmt.Errorf("BOT_TOKEN environment variable is required")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return fmt.Errorf("failed to create bot: %w", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	handler, err := NewHandler(bot)
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}

	defer handler.Stop()

	go func() {
		port := getEnv("PORT", "8080")
		log.Printf("Health check server starting on port %s", port)
		if err := RunServer(port); err != nil {
			log.Printf("Health check server error: %v", err)
		}
	}()

	return handler.Start()
}

// getEnv returns environment variable or default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
