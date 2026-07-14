package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	tgb "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	// Хранилище зарегистрированных пользователей:
	registered = make(map[int64]string) // key – chat ID, value – username
	// Путь к папке с изображениями (должна существовать в репозитории)
	imgDir = "img"
)

// ---------- Утилиты ---------------------------------------------------------

func randomImage() (string, error) {
	files, err := os.ReadDir(imgDir)
	if err != nil {
		return "", err
	}
	var images []string
	for _, f := range files {
		if !f.IsDir() {
			ext := filepath.Ext(f.Name())
			switch ext {
			case ".jpg", ".jpeg", ".png", ".gif":
				images = append(images, filepath.Join(imgDir, f.Name()))
			}
		}
	}
	if len(images) == 0 {
		return "", os.ErrNotExist
	}
	randIdx := rand.Intn(len(images))
	return images[randIdx], nil
}

// ---------- Обработчики -----------------------------------------------------

func startHandler(ctx context.Context, b *tgb.Bot, update *models.Update) {
	msg := &tgb.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Добро пожаловать в карточную игру LightsOut: Grid Wars",
		ReplyMarkup: registrationKeyboard(),
	}
	_, err := b.SendMessage(ctx, msg)
	if err != nil {
		log.Printf("send start message error: %v", err)
	}
}

func callbackHandler(ctx context.Context, b *tgb.Bot, update *models.Update) {
	cb := update.CallbackQuery
	data := cb.Data

	switch data {
	case "info":
		msg := &tgb.SendMessageParams{
			ChatID: cb.Message.Chat.ID,
			Text:   "🎉 Танечка — наш главный персонаж! Она любит собирать карты и делиться ими с друзьями.",
		}
		b.SendMessage(ctx, msg)

	case "register":
		user := cb.From
		registered[cb.Message.Chat.ID] = user.UserName
		msg := &tgb.SendMessageParams{
			ChatID:      cb.Message.Chat.ID,
			Text:        "✅ Вы успешно зарегистрированы! Теперь можете получить карточку.",
			ReplyMarkup: getCardKeyboard(),
		}
		b.SendMessage(ctx, msg)

	case "get_card":
		path, err := randomImage()
		if err != nil {
			log.Printf("random image error: %v", err)
			ans := &tgb.AnswerCallbackQueryParams{
				CallbackQueryID: cb.ID,
				Text:            "Не удалось найти карточку.",
			}
			b.AnswerCallbackQuery(ctx, ans)
			return
		}
		photo := &tgb.SendPhotoParams{
			ChatID: cb.Message.Chat.ID,
			Photo:   tgb.InputFile{File: path},
		}
		_, err = b.SendPhoto(ctx, photo)
		if err != nil {
			log.Printf("send photo error: %v", err)
		}

	default:
		// неизвестный callback – игнорируем
	}
}

// ---------- Клавиатуры -------------------------------------------------------

func registrationKeyboard() *tgb.InlineKeyboardMarkup {
	return &tgb.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgb.InlineKeyboardButton{
			{
				{Text: "Регистрация", CallbackData: "register"},
				{Text: "Инфо", CallbackData: "info"},
			},
		},
	}
}

func getCardKeyboard() *tgb.InlineKeyboardMarkup {
	return &tgb.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgb.InlineKeyboardButton{
			{
				{Text: "Получить карточку", CallbackData: "get_card"},
			},
		},
	}
}

// ---------- main ------------------------------------------------------------

func main() {
	rand.Seed(time.Now().UnixNano())

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN env var is required")
	}

	bot, err := tgb.New(token,
		tgb.WithDefaultHandler(startHandler), // будет вызвано для любого сообщения без специф. хендлера
	)
	if err != nil {
		log.Fatalf("bot init error: %v", err)
	}

	// Обрабатываем только callback‑query (нажатия на inline‑кнопки)
	bot.RegisterHandler(tgb.HandlerTypeCallbackQuery, callbackHandler)

	ctx := context.Background()
	log.Println("Bot started")
	if err = bot.Start(ctx); err != nil {
		log.Fatalf("bot start error: %v", err)
	}
}
