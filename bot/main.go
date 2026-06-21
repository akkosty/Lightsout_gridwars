package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-telegram/bot"
)

var (
	botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	debug    = os.Getenv("DEBUG") == "True"
)

func main() {
	fmt.Println("🤖 Starting Lightsout GridWars Bot...")

	// Обработчик команды /start
	bot.Handle("/start").SetHandler(func(ctx context.Context, b *bot.Bot, update bot.Update) {
		b.SendMessage(&bot.MessageConfig{
			ChatID:    update.Message.Chat.ID,
			Text:      "👋 Привет! Я Lightsout GridWars бот!\n\nЯ могу:\n• Эхо-ответить на сообщения /echo",
			ParseMode: bot.ParseModeMarkdown,
		})
	})

	// Обработчик команды /echo - отправка текста в левую руку
	bot.Handle("/echo").SetHandler(func(ctx context.Context, b *bot.Bot, update bot.Update) {
		messageText := ""
		
		// Извлекаем текст из кнопки (если это инлайн-кнопка), иначе берём из сообщения
		if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.InlineKeyboardMarkup != nil {
			for _, row := range update.Message.ReplyToMessage.InlineKeyboardMarkup.Keyboard {
				for _, btn := range row {
					if inlineBtn, ok := btn.(bot.InlineKeyboardButton); ok {
						messageText = inlineBtn.Text
					}
				}
			}
		}

		// Если текст не найден, используем имя пользователя
		if messageText == "" {
			username := update.Message.From.Username
			userName := fmt.Sprintf("@%s", username)
			messageText = "💬 " + username
		}

		b.SendMessage(&bot.MessageConfig{
			ChatID:    update.Message.Chat.ID,
			Text:      messageText,
			ParseMode: bot.ParseModeMarkdown,
		})
	})

	// Запуск бота
	log.Printf("🚀 Bot starting with token: %s...", botToken[:10]+"***")
	err := b.Run(context.Background(), &bot.Options{
		Token:      botToken,
		Debug:      debug,
		AllowUpdateFromServer: true,
	})

	if err != nil {
		log.Printf("❌ Error starting bot: %s", err)
	} else {
		log.Println("✅ Bot stopped gracefully")
	}
}
