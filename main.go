package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akkosty/Lightsout_gridwars/bot"
)

func main() {
	if err := bot.Run(); err != nil {
		log.Printf("Fatal error: %v", err)
		os.Exit(1)
	}

	fmt.Println("Bot stopped gracefully")
}
