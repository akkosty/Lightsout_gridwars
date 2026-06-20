package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	tg "github.com/go-telegram-bot-api/v5"
)

// Config хранит конфигурационные данные бота
type Config struct {
	BotToken string `json:"bot_token"`
	// TODO: Добавить остальные конфигурации из оригинального python-бота
}

func main() {
	// Чтение config.json
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Ошибка при открытии config.json: %v", err)
	}
	defer file.Close()

	config := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Ошибка при чтении config.json: %v", err)
	}

	fmt.Println("Бот инициализирован!")
	fmt.Printf("Token: %s\n", config.BotToken)
	fmt.Println("-------------------")

	b, err := tg.NewBotAPI(config.BotToken)
	if err != nil {
		log.Fatalf("Ошибка инициализации бота: %v", err)
	}

	// Установка ID процесса для корректной работы в Docker
	b.Debug = true // отладочный режим

	upd := b.GetUpdatesConfig()
	
	// Поддержка запуска бота с указаным портом через команду PORT=8080 или PORT=19352
	portEnv, found := os.LookupEnv("PORT")
	if !found {
		portEnv = "8080" // значение по умолчанию
	}

	// Парсим порт - это для демонстрации возможности установки
	portNum, _ := strconv.Atoi(portEnv)
	fmt.Printf("Бот запущен на порту (информативно): %d\n", portNum)

	u := tg.NewUpdate(0)
	s := b.Listen(upd)
	
	log.Println("Начало работы Telegram бота...")
	for update = range s {
		switch update.Message.Text {
		case "/start":
			b.SendMessage(&tg.MessageConfig{ChatID: int64(update.Message.Chat.ID), Text: "Привет! Я готов к работе!"})
		case "/help", "/helpme":
			b.SendMessage(&tg.MessageConfig{ChatID: int64(update.Message.Chat.ID), Text: "Команда /start для начала"})
		default:
			if update.Message != nil {
				fmt.Printf("Получено сообщение от пользователя %d: %s\n", 
					update.Message.Chat.ID, update.Message.Text)
				
				// Здесь можно добавить логику бота на основе оригинальных python-функций
			}
		}
	}
}