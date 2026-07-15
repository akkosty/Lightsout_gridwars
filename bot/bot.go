package bot

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Run starts the Telegram bot
func Run() error {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN not set")
	}

	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	config := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	updates := botAPI.GetUpdatesChan(config)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello!")
		if _, err := botAPI.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
		}
	}

	return nil
}
