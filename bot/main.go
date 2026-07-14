package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	// Хранилище зарегистрированных пользователей:
	registered = make(map[int64]string) // key – chat ID, value – username
	mu         sync.RWMutex            // мьютекс для безопасного доступа к registered
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

func startHandler(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	msg := &tgbotapi.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Добро пожаловать в карточную игру LightsOut: Grid Wars",
		ParseMode: "HTML",
	}
	_, err := b.SendMessage(ctx, msg)
	if err != nil {
		log.Printf("send start message error: %v", err)
	}
}

func callbackHandler(ctx context.Context, b *tgbotapi.Bot, update *models.Update) {
	cb := update.CallbackQuery
	data := cb.Data

	switch data {
	case "info":
		msg := &tgbotapi.SendMessageParams{
			ChatID:    cb.Message.Chat.ID,
			Text:      "🎉 Танечка — наш главный персонаж! Она любит собирать карты и делиться ими с друзьями.",
			ParseMode: "HTML",
		}
		_, err := b.SendMessage(ctx, msg)
		if err != nil {
			log.Printf("send info message error: %v", err)
		}

	case "register":
		username := cb.From.UserName
		mu.Lock()
		registered[cb.Message.Chat.ID] = username
		mu.Unlock()
		msg := &tgbotapi.SendMessageParams{
			ChatID:    cb.Message.Chat.ID,
			Text:      "✅ Вы успешно зарегистрированы! Теперь можете получить карточку.",
			ParseMode: "HTML",
		}
		_, err := b.SendMessage(ctx, msg)
		if err != nil {
			log.Printf("send register message error: %v", err)
		}

	case "get_card":
		path, err := randomImage()
		if err != nil {
			log.Printf("random image error: %v", err)
			ans := &tgbotapi.AnswerCallbackQueryParams{
				CallbackQueryID: cb.ID,
				Text:            "Не удалось найти карточку.",
			}
			_, err = b.AnswerCallbackQuery(ctx, ans)
			if err != nil {
				log.Printf("answer callback query error: %v", err)
			}
			return
		}
		// Отправка файла по пути (библиотека поддерживает файловую систему напрямую)
		_, err = b.SendPhoto(ctx, &tgbotapi.SendPhotoParams{
			ChatID: cb.Message.Chat.ID,
			Photo:  path,
		})
		if err != nil {
			log.Printf("send photo error: %v", err)
		}

	default:
		// неизвестный callback – игнорируем
		log.Printf("unknown callback data: %s", data)
	}
}

// ---------- Клавиатуры -------------------------------------------------------

func registrationKeyboard() *tgbotapi.InlineKeyboardMarkup {
	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				{Text: "Регистрация", CallbackData: "register"},
				{Text: "Инфо", CallbackData: "info"},
			},
		},
	}
}

func getCardKeyboard() *tgbotapi.InlineKeyboardMarkup {
	return &tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
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

	bot, err := tgbotapi.New(token,
		tgbotapi.WithDefaultHandler(startHandler),
		tgbotapi.WithCallbackQueryDataHandler("^register$", tgbotapi.MatchTypeExact, callbackHandler),
		tgbotapi.WithCallbackQueryDataHandler("^info$", tgbotapi.MatchTypeExact, callbackHandler),
		tgbotapi.WithCallbackQueryDataHandler("^get_card$", tgbotapi.MatchTypeExact, callbackHandler),
	)
	if err != nil {
		log.Fatalf("bot init error: %v", err)
	}

	ctx := context.Background()
	log.Println("Bot started")
	if err = bot.Start(ctx); err != nil {
		log.Fatalf("bot start error: %v", err)
	}
}
