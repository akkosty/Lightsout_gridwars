package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/obj"
)

var botToken = os.Getenv("TELEGRAM_BOT_TOKEN")

func init() {
	if botToken == "" {
		log.Fatal("Ошибка: переменная окружения TELEGRAM_BOT_TOKEN не задана!")
	}
}

// Обработчик команды /start
func HandleStart(update obj.Update) bool {
	msg := update.Message.(*obj.Message)
	
	// Создаем клавиатуру с кнопками
	keyboard := [][]string{
		{"👋 Привет!!", "say_hello"},
		{"ℹ️ Инфо о боте", "info"},
	}
	
	inlineKeyboard, err := obj.NewInlineKeyboardMarkup(keyboard)
	if err != nil {
		log.Printf("Ошибка создания клавиатуры: %v", err)
		return false
	}
	
	msg.Reply(&obj.ReplyOptions{
		Text: "Привет! Я базовый скелет Telegram-бота.\nВыбери действие:",
	})
	msg.InlineKeyboard = inlineKeyboard
	
	return true
}

// Обработчик нажатий на кнопки
func HandleCallback(update obj.Update) bool {
	callback := update.CallbackQuery.(*obj.CallbackQuery)
	
	callback.Message.Edit(&obj.EditedMessageOptions{
		Text: callback.Data,
	})
	
	switch callback.Data {
	case "say_hello":
		callback.Message.EditText(&obj.EditedMessageOptions{
			Text: "👋 Привет, Танечка! Ты лучшая ❤️1",
		})
	case "info":
		callback.Message.EditText(&obj.EditedMessageOptions{
			Text: "🤖 Я простой бот-скелет на Go.\n" +
				"Подключен через библиотеку @go-telegram/bot\n" +
				"Добавляйте свои команды и логику!",
		})
	default:
		callback.Message.EditText(&obj.EditedMessageOptions{
			Text: fmt.Sprintf("❓ Неизвестный callback: %s", callback.Data),
		})
	}
	
	return true
}

// Обработчик эхо-команды
func HandleEcho(update obj.Update) bool {
	msg := update.Message.(*obj.Message)
	if msg == nil {
		return false
	}
	
	msg.Reply(&obj.ReplyOptions{
		Text: fmt.Sprintf("🗣 Вы сказали: %s", msg.Text),
	})
	
	return true
}

func main() {
	// Создаем бота
	handler := bot.NewBotHandler(botToken, HandleStart, HandleCallback, HandleEcho)
	
	// Запуск в режиме polling (для локального теста и CI/CD)
	log.Println("Бот запущен, начинаем polling…")
	
	if err := handler.Run(); err != nil {
		log.Fatalf("Ошибка при запуске бота: %v", err)
	}
}
