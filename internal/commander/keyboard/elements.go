package keyboard

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const ReplyKeyboardTextToAllElements = "🌏 К стихиям"

func (m *Manager) ForElements() tgbotapi.ReplyKeyboardMarkup {
	elements := m.getElementsWithEmojis()
	buttons := m.getElementButtons(elements)

	return m.getResizedKeyboard(buttons, CountButtonsInRow)
}

func (m *Manager) getElementButtons(elements []string) []tgbotapi.KeyboardButton {
	buttons := make([]tgbotapi.KeyboardButton, 0)

	for _, el := range elements {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(el))
	}

	buttons = append(buttons, tgbotapi.NewKeyboardButton(ReplyKeyboardTextToMainMenu))

	return buttons
}

// Returns elements like "💧 Гидро"
func (m *Manager) getElementsWithEmojis() []string {
	elements := m.service.Characters.GetElements()
	result := make([]string, 0, len(elements))

	for el := range elements {
		el = m.renderer.AddEmojiToElement(el)
		result = append(result, el)
	}

	return result
}
