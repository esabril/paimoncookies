package commander

import (
	"github.com/esabril/paimoncookies/internal/commander/commands"
	"github.com/esabril/paimoncookies/internal/commander/pager"
	"github.com/esabril/paimoncookies/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const DefaultMaxDataValuesOnMenu = 6

type Commander struct {
	bot       *tgbotapi.BotAPI
	commander *commands.Commander
	service   *service.Service
}

func New(bot *tgbotapi.BotAPI, s *service.Service, templatePath string) *Commander {
	return &Commander{
		bot:     bot,
		service: s,
		commander: commands.NewCommander(s, templatePath, pager.NewPager(
			s.Characters.GetElements(),
			DefaultMaxDataValuesOnMenu,
		)),
	}
}

func (c *Commander) HandleCommands(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	msg.ParseMode = "markdown"
	msg.ReplyMarkup = c.commander.KeyboardManager.GetMainMenu()

	c.service.SetTodayWeekday()

	if update.Message.IsCommand() {
		c.commander.HandleCommand(&msg, update.Message.Command())
	} else {
		c.commander.HandleMessage(&msg, update.Message.Text)
	}

	c.bot.Send(msg)
}
