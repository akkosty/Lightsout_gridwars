package main

import (
	"fmt"
	"image/png"
	"io/fs"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() error {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		return fmt.Errorf("TELEGRAM_BOT_TOKEN not set")
	}

	botAPI, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}

	log.Printf("Authorized on account %s", botAPI.Self.UserName)

	config := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	updates := botAPI.GetUpdatesChan(config)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "/start":
			sendStartMenu(botAPI, update.Message.Chat.ID)
		case "Инфо":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Танечка ты супер!!!")
			sendStartMenu(botAPI, update.Message.Chat.ID)
			if _, err := botAPI.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
			}
		case "Зарегистрироваться":
			imgPath, err := getRandomImage("./img")
			if err != nil {
				log.Printf("Error getting random image: %v", err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка загрузки картинки")
				sendStartMenu(botAPI, update.Message.Chat.ID)
				if _, err := botAPI.Send(msg); err != nil {
					log.Printf("Error sending message: %v", err)
				}
			} else {
				sendImageWithStartMenu(botAPI, update.Message.Chat.ID, imgPath)
			}
		case "Назад":
			sendStartMenu(botAPI, update.Message.Chat.ID)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда")
			sendStartMenu(botAPI, update.Message.Chat.ID)
			if _, err := botAPI.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
			}
		}
	}

	return nil
}

func getRandomImage(imgDir string) (string, error) {
	var pngFiles []string
	err := filepath.WalkDir(imgDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".png" {
			pngFiles = append(pngFiles, path)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if len(pngFiles) == 0 {
		return "", fmt.Errorf("no PNG files found in %s", imgDir)
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return pngFiles[rng.Intn(len(pngFiles))], nil
}

func sendImageWithStartMenu(botAPI *tgbotapi.BotAPI, chatID int64, imgPath string) {
	file, err := os.Open(imgPath)
	if err != nil {
		log.Printf("Error opening image file: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Ошибка открытия картинки")
		sendStartMenu(botAPI, chatID)
		if _, err := botAPI.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
		}
		return
	}
	defer file.Close()

	// Verify it's a valid PNG
	_, err = png.Decode(file)
	if err != nil {
		log.Printf("Invalid PNG image: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Неверный формат картинки")
		sendStartMenu(botAPI, chatID)
		if _, err := botAPI.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
		}
		return
	}

	// Seek back to start after decoding
	if _, err := file.Seek(0, 0); err != nil {
		log.Printf("Error seeking file: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Ошибка при подготовке картинки")
		if _, err := botAPI.Send(msg); err != nil {
			log.Printf("Error sending message: %v", err)
		}
		return
	}

	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileReader{
		Name:   "grid.png",
		Reader: file,
	})

	if _, err := botAPI.Send(photo); err != nil {
		log.Printf("Error sending photo: %v", err)
	}

	sendStartMenu(botAPI, chatID)
}

func sendStartMenu(botAPI *tgbotapi.BotAPI, chatID int64) {
	keyboard := tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Инфо"),
		tgbotapi.NewKeyboardButton("Зарегистрироваться"),
	)

	replyMarkup := tgbotapi.ReplyKeyboardMarkup{
		Keyboard:        [][]tgbotapi.KeyboardButton{keyboard},
		ResizeKeyboard:  true,
		OneTimeKeyboard: false,
	}

	msg := tgbotapi.NewMessage(chatID, "")
	msg.ReplyMarkup = &replyMarkup

	if _, err := botAPI.Send(msg); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func main() {
	if err := Run(); err != nil {
		log.Printf("Bot error: %v", err)
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
