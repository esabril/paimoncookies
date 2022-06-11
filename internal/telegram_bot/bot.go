package telegram_bot

import (
	"github.com/esabril/paimoncookies/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func GetBot(c *service.Config) *tgbotapi.BotAPI {
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = c.Bot.Debug

	log.Printf("Running Telegram Bot Handler. Authorized on account %s", bot.Self.UserName)

	return bot
}
