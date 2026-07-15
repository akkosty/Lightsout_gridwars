package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	updateConfig := tgbotapi.UpdateConfig{
		Offset:  0,
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			keyboard := tgbotapi.NewInlineKeyboardMarkup(
				tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Инфо", "info"),
					tgbotapi.NewInlineKeyboardButtonData("Зарегистрироваться", "register"),
				),
			)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
			msg.ReplyMarkup = keyboard

			if _, err := bot.Send(msg); err != nil {
				log.Println("Error sending message:", err)
			}
		}
	}
}
