package commander

import (
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/tools/renderer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot      *tgbotapi.BotAPI
	service  *service.Service
	renderer *renderer.Renderer
}

func New(bot *tgbotapi.BotAPI, s *service.Service, templatePath string) *Commander {
	return &Commander{
		bot:      bot,
		service:  s,
		renderer: renderer.NewRenderer(templatePath),
	}
}

func (c *Commander) HandleCommands(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.IsCommand() {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ParseMode = "markdown"

		switch update.Message.Command() {
		case CommandAgenda:
			msg.Text = c.GetAgenda()
		default:
			msg.Text = "Паймон перестает тебя понимать. Пойдем лучше поедим?" // todo: random Paimon phrases
		}

		c.bot.Send(msg)
	}
}
