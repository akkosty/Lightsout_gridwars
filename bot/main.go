package main

import (
    "context"
    "log"
    "os"
    "github.com/go-telegram/bot"
    "github.com/go-telegram/bot/models"
)

func main() {
    token := os.Getenv("TELEGRAM_BOT_TOKEN")
    if token == "" {
        log.Fatal("TELEGRAM_BOT_TOKEN not set")
    }
    // Create bot instance
    client, err := bot.New(token)
    if err != nil {
        log.Fatalf("Failed to create bot: %v", err)
    }
    defer client.Close()

    // Set up handlers
    client.HandleMessage(func(ctx context.Context, b *bot.Bot, msg *models.Message) error {
        if msg.Text == "/start" {
            keyboard := []bot.InlineKeyboardButton{{
                Text:         "👋 Привет!",
                CallbackData: "say_hello",
            }, {
                Text:         "ℹ️ Инфо о боте",
                CallbackData: "info",
            }}
            _, err := b.SendMessage(ctx, &bot.SendMessageParams{
                ChatID:    msg.Chat.ID,
                Text:      "Привет! Выбери действие:",
                ReplyMarkup: bot.InlineKeyboardMarkup{InlineKeyboard: [][]bot.InlineKeyboardButton{{keyboard[0]}, {keyboard[1]}}},
            })
            return err
        }
        // Echo for any other text (debug)
        if msg.Text != "" {
            _, err := b.SendMessage(ctx, &bot.SendMessageParams{ChatID: msg.Chat.ID, Text: "🗣 Вы сказали: " + msg.Text})
            return err
        }
        return nil
    })

    client.HandleCallbackQuery(func(ctx context.Context, b *bot.Bot, query *models.CallbackQuery) error {
        // Acknowledge callback
        _, _ = b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{CallbackQueryID: query.ID})
        switch query.Data {
        case "say_hello":
            _, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
                ChatID:    query.Message.Chat.ID,
                MessageID: query.Message.MessageID,
                Text:      "👋 Привет, Танечка! Ты лучшая ❤️1",
            })
            return err
        case "info":
            _, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
                ChatID:    query.Message.Chat.ID,
                MessageID: query.Message.MessageID,
                Text:      "🤖 Я простой бот‑скелет. Добавляйте свои команды и логику!",
            })
            return err
        default:
            _, err := b.EditMessageText(ctx, &bot.EditMessageTextParams{
                ChatID:    query.Message.Chat.ID,
                MessageID: query.Message.MessageID,
                Text:      "❓ Неизвестный callback: " + query.Data,
            })
            return err
        }
    })

    // Start polling
    log.Println("Bot started, polling…")
    if err := client.Start(context.Background()); err != nil {
        log.Fatalf("Polling error: %v", err)
    }
}
