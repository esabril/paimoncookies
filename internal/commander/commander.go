package commander

import (
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/tools/renderer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	ReplyKeyboardTextAgenda = "🗓️ Расскажи, что я могу сделать сегодня?"
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

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ParseMode = "markdown"

	replyKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(ReplyKeyboardTextAgenda),
		),
	)

	msg.ReplyMarkup = replyKeyboard

	c.service.SetTodayWeekday()

	if update.Message.IsCommand() {
		switch update.Message.Command() {
		case CommandStart:
			msg.Text = c.GetStart()
			break
		case CommandAgenda:
			msg.Text = c.GetAgenda()
			break
		default:
			msg.Text = "Паймон перестает тебя понимать. Пойдем лучше поедим?" // todo: random Paimon phrases
		}

		c.bot.Send(msg)

		return
	}

	switch update.Message.Text {
	case ReplyKeyboardTextAgenda:
		msg.Text = c.GetAgenda()
		break
	default:
		msg.Text = "Это очень интересная мыс... О, смотри, бабочка!" // todo: random Paimon phrases
		break
	}

	c.bot.Send(msg)
}
