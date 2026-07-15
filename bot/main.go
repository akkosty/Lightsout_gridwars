package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

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

		switch update.Message.Text {
		case "/start":
			sendStartMenu(botAPI, update.Message.Chat.ID)
		case "Инфо":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Танечка ты супер!!!")
			sendStartMenu(botAPI, update.Message.Chat.ID)
			if _, err := botAPI.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
			}
		case "Зарегистрироваться":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Регистрация пока не доступна!")
			sendStartMenu(botAPI, update.Message.Chat.ID)
			if _, err := botAPI.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
			}
		case "Назад":
			sendStartMenu(botAPI, update.Message.Chat.ID)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
			sendStartMenu(botAPI, update.Message.Chat.ID)
			if _, err := botAPI.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	}

	return nil
}

func sendStartMenu(botAPI *tgbotapi.BotAPI, chatID int64) {
	keyboard := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Инфо"),
		tgbotapi.NewKeyboardButton("Зарегистрироваться"),
	)

	replyMarkup := tgbotapi.ReplyKeyboardMarkup{
		Keyboard:        [][]tgbotapi.KeyboardButton{keyboard},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}

	msg := tgbotapi.NewMessage(chatID, "Выберите действие:")
	msg.ReplyMarkup = &replyMarkup

	if _, err := botAPI.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func main() {
	if err := Run(); err != nil {
		log.Printf("Bot error: %v", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
