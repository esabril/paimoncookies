package main

import (
	"github.com/BurntSushi/toml"
	"github.com/esabril/paimoncookies/internal/commander"
	"github.com/esabril/paimoncookies/internal/server"
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/internal/telegram_bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Panic")
		}
	}()

	var c *service.Config

	_, err := toml.DecodeFile("configs/test.toml", &c)
	if err != nil {
		log.Fatalf("Unable to read config file: %s\n", err.Error())
	}

	log.Printf("\n\nWelcome to Paimon Cookies Application. Current version: %s\n\n", c.Version)

	s := service.NewService(c)

	apiCh := make(chan string, 1)
	// maybe should change to Cobra if CLI-format will more useful
	apiEndpoint := server.GetRouter(s)

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
	comm := commander.New(bot, s)

	for update := range updates {
		comm.HandleCommands(update)
	}
}
