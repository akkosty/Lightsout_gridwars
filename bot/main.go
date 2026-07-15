package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-telegram/bot"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		slog.Error("TELEGRAM_BOT_TOKEN is not set")
	}

	b, err := bot.New(botToken)
	if err != nil {
		slog.Error("failed to create bot", "error", err)
		return
	}

	slog.Info("bot started")
	b.Start(ctx)
}
