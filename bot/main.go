package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram/bot"
)

// ---------- База данных (простая реализация) ----------------------------------

type PlayerData struct {
	ID       int64
	Username string
	Card     *PlayerCard
}

type PlayerCard struct {
	Name      string
	Strength  int
	Agility   int
	Intellect int
	Vitality  int
}

var (
	registered = make(map[int64]string)
	mu         sync.Mutex
	cards      []*PlayerCard
)

// Инициализация карточек
func init() {
	rand.Seed(time.Now().UnixNano())
	generateCards()
}

func generateCards() {
	cardNames := []string{
		"Воин-Паладин", "Следопыт", "Маг Огня",
		"Теневой Ninja", "Лекарь", "Крушитель",
	}
	for _, name := range cardNames {
		cards = append(cards, &PlayerCard{
			Name:      name,
			Strength:  rand.Intn(10) + 1,
			Agility:   rand.Intn(10) + 1,
			Intellect: rand.Intn(10) + 1,
			Vitality:  rand.Intn(10) + 1,
		})
	}
}

// ---------- Обработчики -------------------------------------------------------

func startHandler(ctx context.Context, b *tgbotapi.Bot, msg *tgbotapi.Message) {
	text := "Добро пожаловать в карточную игру LightsOut: Grid Wars"
	params := &tgbotapi.SendMessageParams{
		ChatID:      msg.Chat.ID,
		Text:        text,
		ParseMode:   "HTML",
		ReplyMarkup: registrationKeyboard(),
	}
	_, err := b.SendMessage(ctx, params)
	if err != nil {
		log.Printf("send start message error: %v", err)
	}
}

func callbackHandler(ctx context.Context, b *tgbotapi.Bot, cb tgbotapi.CallbackQuery) {
	data := cb.Data

	switch data {
	case "register":
		username := cb.From.Username
		mu.Lock()
		registered[cb.Message.Chat.ID] = username
		mu.Unlock()

		text := fmt.Sprintf("✅ Вы успешно зарегистрированы как @%s", username)
		editParams := &tgbotapi.EditMessageTextParams{
			ChatID:    cb.Message.Chat.ID,
			MessageID: cb.Message.MessageID,
			Text:      text,
		}
		b.EditMessageText(ctx, editParams)

	case "info":
		mu.Lock()
		username := registered[cb.Message.Chat.ID]
		mu.Unlock()

		if username == "" {
			text := "❌ Сначала выполните регистрацию"
			editParams := &tgbotapi.EditMessageTextParams{
				ChatID:      cb.Message.Chat.ID,
				MessageID:   cb.Message.MessageID,
				Text:        text,
				ReplyMarkup: registrationKeyboard(),
			}
			b.EditMessageText(ctx, editParams)
			return
		}

		info := "Информация о пользователе:\n" +
			"Username: @" + username + "\n" +
			"Статус: Игрок\n" +
			"Игра: LightsOut: Grid Wars"

		editParams := &tgbotapi.EditMessageTextParams{
			ChatID:      cb.Message.Chat.ID,
			MessageID:   cb.Message.MessageID,
			Text:        info,
			ReplyMarkup: registrationKeyboard(),
		}
		b.EditMessageText(ctx, editParams)

	case "get_card":
		mu.Lock()
		if _, exists := registered[cb.Message.Chat.ID]; !exists {
			mu.Unlock()
			text := "❌ Сначала выполните регистрацию"
			editParams := &tgbotapi.EditMessageTextParams{
				ChatID:      cb.Message.Chat.ID,
				MessageID:   cb.Message.MessageID,
				Text:        text,
				ReplyMarkup: registrationKeyboard(),
			}
			b.EditMessageText(ctx, editParams)
			return
		}
		mu.Unlock()

		card := cards[rand.Intn(len(cards))]

		text := fmt.Sprintf("🏆 Ваша карточка:\n\n%s", formatCard(card))
		editParams := &tgbotapi.EditMessageTextParams{
			ChatID:      cb.Message.Chat.ID,
			MessageID:   cb.Message.MessageID,
			Text:        text,
			ReplyMarkup: getCardKeyboard(),
		}
		b.EditMessageText(ctx, editParams)
	}
}

func formatCard(card *PlayerCard) string {
	return "Имя: " + card.Name +
		"\nСила: ⚔️" + intToString(card.Strength) +
		"\nЛовкость: 🏃" + intToString(card.Agility) +
		"\nИнтеллект: 🧠" + intToString(card.Intellect) +
		"\nВыносливость: ❤️" + intToString(card.Vitality)
}

func intToString(n int) string {
	if n < 10 {
		return "0" + strconv.Itoa(n)
	}
	return strconv.Itoa(n)
}

// ---------- Клавиатуры -------------------------------------------------------

func registrationKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
			{
				{Text: "Регистрация", CallbackData: "register"},
				{Text: "Инфо", CallbackData: "info"},
			},
		},
	}
}

func getCardKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.InlineKeyboardMarkup{
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

	bot, err := tgbotapi.NewBot(token)
	if err != nil {
		log.Fatalf("bot init error: %v", err)
	}

	ctx := context.Background()

	// Исправленный способ регистрации обработчиков
	bot.HandleMessage(func(ctx context.Context, bot *tgbotapi.Bot, msg *tgbotapi.Message) {
		startHandler(ctx, bot, msg)
	})

	bot.HandleFunc("callback_query", func(ctx context.Context, bot *tgbotapi.Bot, cb tgbotapi.CallbackQuery) {
		callbackHandler(ctx, bot, cb)
	})

	// Health check server (port 8080 для Render)
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})

		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}
		log.Printf("Health check server listening on :%s", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	log.Println("Bot started")
	bot.Start(ctx)
}
