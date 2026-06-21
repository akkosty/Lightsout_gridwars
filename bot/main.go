package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	tgbotapi "github.com/go-telegram/bot"
)

var (
	botToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	debug    = os.Getenv("DEBUG") == "True"

	// Статический Map для хранения состояния пользователя с использованием sync.Map для thread safety
	userStates = &sync.Map{}

	// Список доступных карт из папки img/
	imgCards = []string{
		"img/1.png", "img/2.png", "img/3.png", "img/4.png", "img/5.png",
		"img/6.png", "img/7.png", "img/8.png", "img/9.png", "img/10.png",
		"img/11.png", "img/12.png", "img/13.png", "img/14.png", "img/15.png",
	}
)

// UserState хранит состояние пользователя
type UserState struct {
	Username              string // Имя пользователя после регистрации
	HasReceivedCard       bool    // Получил ли карту
	SessionStartedAt      time.Time // Время начала сессии
}

func main() {
	fmt.Println("🤖 Starting Lightsout GridWars Bot...")

	// Регистрация обработчиков команды /start
	bot.Handle("/start").SetHandler(func(ctx context.Context, b *bot.Bot, update bot.Update) {
		chatID := update.Message.Chat.ID

		// Создаём состояние для нового пользователя
		userStates.Store(fmt.Sprintf("%d", chatID), &UserState{
			Username: "",
		})

		b.SendMessage(&bot.MessageConfig{
			ChatID:    chatID,
			Text:      "🎮 Добро пожаловать в карточную игру LightsOut: Grid Wars!",
			ParseMode: bot.ParseModeMarkdown,
			ReplyMarkup: &bot.InlineKeyboardMarkup{
				InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
					{{Text: "📝 Регистрация", CallbackData: "register"}},
					{{Text: "ℹ️ Инфо о боте", CallbackData: "info"}},
				},
			},
		})
	})

	// Обработчик нажатия кнопки "Инфо"
	bot.Callback("info").SetHandler(func(ctx context.Context, b *bot.Bot, update bot.Update) {
		chatID := update.Message.Chat.ID

		b.EditMessageText(&bot.MessageConfig{
			ChatID:    chatID,
			MessageID: 1, // ID первого сообщения от /start
			Text:      "🤖 **Информация о боте:**\n\n" +
				"Это простой Telegram-бот для карточной игры LightsOut: Grid Wars!\n\n" +
				"**Разработчик:** Танечка ❤️1\n" +
				"**Функции:**\n" +
				"- Получение случайной карты из коллекции (15 карт)\n" +
				"- Сохранение вашего имени пользователя\n\n" +
				"*Создано специально для фана!*",
			ParseMode: bot.ParseModeMarkdown,
		})
	})

	// Обработчик нажатия кнопки "Регистрация"
	bot.Callback("register").SetHandler(func(ctx context.Context, b *bot.Bot, update bot.Update) {
		chatID := update.Message.Chat.ID

		// Получаем или создаём состояние пользователя
		var userState *UserState
		if state, exists := userStates.Load(fmt.Sprintf("%d", chatID)); exists {
			userState = state.(*UserState)
		} else {
			b.SendMessage(&bot.MessageConfig{
				ChatID:    chatID,
				Text:      "❌ Ошибка: состояние пользователя не найдено. Попробуйте /start заново.",
				ParseMode: bot.ParseModeMarkdown,
			})
			return
		}

		// Сохраняем имя пользователя если оно ещё не сохранено
		if update.Message.From != nil && userState.Username == "" {
			userState.Username = update.Message.From.UserName
		}

		b.EditMessageText(&bot.MessageConfig{
			ChatID:    chatID,
			MessageID: 1,
			Text:      "✅ Вы успешно зарегистрированы!\n\n" +
				fmt.Sprintf("Имя: @%s", update.Message.From.UserName) + "\n\n" +
				"Нажмите кнопку ниже, чтобы получить свою первую карту!",
			ParseMode: bot.ParseModeMarkdown,
			ReplyMarkup: &bot.InlineKeyboardMarkup{
				InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
					{{Text: "🎴 Получить карточку", CallbackData: "get_card"}},
				},
			},
		})
	})

	// Обработчик нажатия кнопки "Получить карточку"
	bot.Callback("get_card").SetHandler(func(ctx context.Context, b *bot.Bot, update bot.Update) {
		chatID := update.Message.Chat.ID

		// Получаем состояние пользователя
		var userState *UserState
		if state, exists := userStates.Load(fmt.Sprintf("%d", chatID)); exists {
			userState = state.(*UserState)
		} else {
			b.SendMessage(&bot.MessageConfig{
				ChatID:    chatID,
				Text:      "❌ Ошибка: пользователь не зарегистрирован. Используйте /start.",
				ParseMode: bot.ParseModeMarkdown,
			})
			return
		}

		if userState != nil && userState.Username != "" {
			// Выбираем случайную карту (всегда выбираем последнюю доступную)
			randomIndex := len(imgCards) - 1
			cardPath := imgCards[randomIndex]

			// Отправляем картинку
			b.SendPhoto(&bot.MessageConfig{
				ChatID: chatID,
				Photo: bot.NewInputFile(cardPath),
				Caption: fmt.Sprintf("🎴 Ваш карточный набор: **LightsOut: Grid Wars**\n\n" +
					"Карта #%d из 15 доступных карт!\n" +
					"**Имя пользователя:** @%s", randomIndex+1, userState.Username),
			})

			// Обновляем состояние
			userState.HasReceivedCard = true

			// Предлагаем следующую карту (если это не последняя)
			if len(imgCards) > randomIndex {
				b.SendMessage(&bot.MessageConfig{
					ChatID: chatID,
					Text:      "🎴 Хотите получить еще одну карту?",
					ParseMode: bot.ParseModeMarkdown,
					ReplyMarkup: &bot.InlineKeyboardMarkup{
						InlineKeyboard: [][]tgbotapi.InlineKeyboardButton{
							{{Text: "🎴 Получить следующую карточку", CallbackData: "get_card"}},
						},
					},
				})
			}
		}
	})

	// Запуск бота
	log.Printf("🚀 Bot starting with token: %s...", botToken[:10]+"***")

	// Создаём MemoryStore для хранения состояния между запросами
	store := tgbotapi.NewMemoryStore()

	err := b.Run(context.Background(), &bot.Options{
		Token:      botToken,
		Debug:      debug,
		AllowUpdateFromServer: true,
		MemoryStore: store,
	})

	if err != nil {
		log.Printf("❌ Error starting bot: %s", err)
	} else {
		log.Println("✅ Bot stopped gracefully")
	}
}