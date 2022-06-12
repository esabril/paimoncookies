package main

import (
	"github.com/esabril/paimoncookies/internal/commander"
	"github.com/esabril/paimoncookies/internal/server"
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/internal/telegram_bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic with error:", r)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	c := service.ParseConfigFromEnv()

	log.Printf("\n\nWelcome to Paimon Cookies Application. Current version: %s\n\n", os.Getenv("APP_VERSION"))

	s := service.NewService(c)

	apiCh := make(chan string, 1)
	// maybe should change to Cobra if CLI-format will more useful
	apiEndpoint := server.GetRouter(s, c)

	go func() {
		apiCh <- "Running API endpoint server"

		apiEndpoint.Run("127.0.0.1:8087")
	}()

	log.Println(<-apiCh)
	close(apiCh)

	bot := telegram_bot.GetBot(c)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = c.Bot.Timeout

	updates := bot.GetUpdatesChan(u)
	comm := commander.New(bot, s, "internal/commander/template/")

	for update := range updates {
		comm.HandleCommands(update)
	}
}
