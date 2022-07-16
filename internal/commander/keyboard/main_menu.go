package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	CountMainMenuButtonsInRow = 1

	ReplyKeyboardTextAgenda     = "🗓️ Расскажи, что я могу сделать сегодня?"
	ReplyKeyboardTextCharacters = "📋️ Покажи мне информацию о персонаже"
	ReplyKeyboardTextToMainMenu = "🍵 В главное меню"
)

func (m *Manager) GetMainMenu() tgbotapi.ReplyKeyboardMarkup {
	buttons := []tgbotapi.KeyboardButton{
		tgbotapi.NewKeyboardButton(ReplyKeyboardTextAgenda),
		tgbotapi.NewKeyboardButton(ReplyKeyboardTextCharacters),
	}

	return m.getKeyboard(nil, buttons, CountMainMenuButtonsInRow)
}
