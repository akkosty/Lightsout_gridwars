package main

import (
	"context"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Handler manages Telegram bot updates and game state
type Handler struct {
	bot    *tgbotapi.BotAPI
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// NewHandler creates a new handler instance
func NewHandler(bot *tgbotapi.BotAPI) (*Handler, error) {
	return &Handler{
		bot: bot,
	}, nil
}

// Start begins processing updates
func (h *Handler) Start() error {
	ctx, cancel := context.WithCancel(context.Background())
	h.cancel = cancel

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := h.bot.GetUpdatesChan(ctx, u)
	if err != nil {
		return err
	}

	for update := range updates {
		h.wg.Add(1)
		go func(u tgbotapi.Update) {
			defer h.wg.Done()
			h.processUpdate(u)
		}(update)
	}

	return nil
}

// Stop gracefully stops the handler
func (h *Handler) Stop() {
	if h.cancel != nil {
		h.cancel()
	}
	h.wg.Wait()
}

func (h *Handler) processUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello!")
	h.bot.Send(msg)
}
