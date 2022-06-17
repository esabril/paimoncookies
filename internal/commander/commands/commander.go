package commands

import (
	"github.com/esabril/paimoncookies/internal/commander/keyboard"
	"github.com/esabril/paimoncookies/internal/commander/pager"
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/tools/renderer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	service         *service.Service
	renderer        *renderer.Renderer
	KeyboardManager *keyboard.Manager
}

func NewCommander(s *service.Service, templatePath string, pager *pager.Pager) *Commander {
	r := renderer.NewRenderer(templatePath)

	return &Commander{
		service:         s,
		renderer:        r,
		KeyboardManager: keyboard.NewManager(s, pager, r),
	}
}

// HandleCommand example: /start
func (c *Commander) HandleCommand(msg *tgbotapi.MessageConfig, command string) {
	switch command {
	case CommandStart:
		msg.Text = c.GetStart()
		break
	case CommandAgenda:
		msg.Text = c.GetAgenda()
		break
	default:
		msg.Text = "Паймон перестает тебя понимать. Пойдем лучше поедим?"
	}
}

// HandleMessage from reply keyboard. It's a default bot behaviour
func (c *Commander) HandleMessage(msg *tgbotapi.MessageConfig, text string) {
	if c.isElement(text) {
		c.KeyboardManager.SetPageFromReply(text, msg.ChatID)

		element := c.getElementFromReply(text)
		msg.Text = c.GetCharacterMenuRules(element)
		msg.ReplyMarkup = c.KeyboardManager.ForCharacters(element, msg.ChatID)

		return
	}

	if c.isFlushPagerCase(text) {
		c.KeyboardManager.FlushPager(msg.ChatID)
	}

	if c.isCharacter(text) {
		info, element := c.GetCharacterInfo(text)
		msg.Text = info
		msg.ReplyMarkup = c.KeyboardManager.ForCharacters(element, msg.ChatID)

		return
	}

	switch text {
	case keyboard.ReplyKeyboardTextToMainMenu:
		msg.Text = "🌸 О чем Паймон может тебе рассказать?"
		msg.ReplyMarkup = c.KeyboardManager.GetMainMenu()
		break
	case keyboard.ReplyKeyboardTextToAllElements:
		msg.Text = "🌸 Давай поищем кого-нибудь еще..."
		msg.ReplyMarkup = c.KeyboardManager.ForElements()
		break
	case keyboard.ReplyKeyboardTextAgenda:
		msg.Text = c.GetAgenda()
		break
	case keyboard.ReplyKeyboardTextCharacters:
		msg.Text = c.GetCharacterMenuRules("")
		msg.ReplyMarkup = c.KeyboardManager.ForElements()
		break
	default:
		msg.Text = "Это очень интересная мыс... О, смотри, бабочка!" // todo: random Paimon phrases
		msg.ReplyMarkup = c.KeyboardManager.GetMainMenu()
		break
	}
}

// Cases for which we do not reset the paginator:
// - Characters menu
func (c *Commander) isFlushPagerCase(text string) bool {
	return text != keyboard.ReplyKeyboardTextCharacters && !c.isCharacter(text)
}
