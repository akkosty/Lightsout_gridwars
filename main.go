package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Конфигурация бота
type Config struct {
	Username       string
	Pass           string
	LogicalGrid    bool
	ShowLog        bool
	MaxDepth       int
	Delay          int
	Timeout        int
	CooldownPeriod int
	NewLineMode    bool
}

// Парсинг параметров из окружения или дефолтные значения
func loadConfig() Config {
	return Config{
		Username:       getEnv("USERNAME", "akkosty"),
		Pass:           getEnv("PASS", "password"), // Для безопасности лучше использовать .env файл
		LogicalGrid:    parseEnvBool("LOGICAL_GRID", true),
		ShowLog:        parseEnvBool("SHOW_LOG", true),
		MaxDepth:       parseEnvInt("MAX_DEPTH", 3),
		Delay:          parseEnvInt("DELAY", 200), // Задержка между ходами (мс)
		Timeout:        parseEnvInt("TIMEOUT", 60),
		CooldownPeriod: parseEnvInt("COOLDOWN_PERIOD", 15),
		NewLineMode:    parseEnvBool("NEW_LINE_MODE", false),
	}
}

func getEnv(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}

func parseEnvBool(key string, defaultValue bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		log.Printf("ParseBool(%q) error: %v, using default", key, err)
		return defaultValue
	}
	return b
}

func parseEnvInt(key string, defaultValue int) int {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("Atoi(%q) error: %v, using default", key, err)
		return defaultValue
	}
	return i
}

// Главная функция бота
func main() {
	config := loadConfig()

	fmt.Println("GridWars Bot started...")
	fmt.Printf("Username: %s\n", config.Username)
	fmt.Printf("Logical Grid: %v\n", config.LogicalGrid)
	fmt.Printf("Max Depth: %d\n", config.MaxDepth)
	fmt.Printf("Delay: %dms\n", config.Delay)

	// Здесь должна быть основная логика бота:
	// 1. Подключение к Telegram Webhook API
	// 2. Обработка входящих сообщений от игрового сервера
	// 3. Принятие решений (AI/логика игры)
	// 4. Отправка ходов на шахматную доску
	
	// Для примера - запуск веб-сервера для проверки статуса
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintln(w, "GridWars Bot is running!")
			fmt.Fprintf(w, "Config: %+v\n", config)
		}
	})

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO: Реализовать WebSocket соединение с игровым сервером
// TODO: Реализовать логику принятия решений
// TODO: Добавить обработку ошибок и логирование