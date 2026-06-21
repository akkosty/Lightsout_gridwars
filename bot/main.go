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
	msg := tgb.NewMessage(update.Message.Chat.ID,
		"Добро пожаловать в карточную игру LightsOut: Grid Wars")
	msg.ReplyMarkup = registrationKeyboard()
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
		answer := tgb.NewMessage(cb.Message.Chat.ID,
			"🎉 Танечка — наш главный персонаж! Она любит собирать карты и делиться ими с друзьями.")
		b.SendMessage(ctx, answer)

	case "register":
		user := cb.From
		registered[cb.Message.Chat.ID] = user.UserName
		answer := tgb.NewMessage(cb.Message.Chat.ID,
			"✅ Вы успешно зарегистрированы! Теперь можете получить карточку.")
		answer.ReplyMarkup = getCardKeyboard()
		b.SendMessage(ctx, answer)

	case "get_card":
		path, err := randomImage()
		if err != nil {
			log.Printf("random image error: %v", err)
			b.AnswerCallbackQuery(ctx,
				tgb.NewAnswerCallbackQuery(cb.ID).WithText("Не удалось найти карточку."))
			return
		}
		photo := tgb.NewPhoto(cb.Message.Chat.ID, path)
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
	return tgb.NewInlineKeyboardMarkup(
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("Регистрация", "register"),
			tgb.NewInlineKeyboardButtonData("Инфо", "info"),
		),
	)
}

func getCardKeyboard() *tgb.InlineKeyboardMarkup {
	return tgb.NewInlineKeyboardMarkup(
		tgb.NewInlineKeyboardRow(
			tgb.NewInlineKeyboardButtonData("Получить карточку", "get_card"),
		),
	)
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
