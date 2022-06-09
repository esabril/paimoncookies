package commander

import (
	"github.com/esabril/paimoncookies/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot     *tgbotapi.BotAPI
	service *service.Service
}

func New(bot *tgbotapi.BotAPI, s *service.Service) *Commander {
	return &Commander{
		bot:     bot,
		service: s,
	}
}

func (c *Commander) HandleCommands(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case CommandAgenda:
			msg.Text = "Здесь будет расписание сегодняшнего дня!"
		default:
			msg.Text = "Паймон перестает тебя понимать. Пойдем лучше поедим?" // todo: random Paimon phrases
		}

		c.bot.Send(msg)
	}
}
