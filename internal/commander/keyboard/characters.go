package keyboard

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// ForCharacters keyboard contains rows with max CountButtonsInRow buttons
func (m *Manager) ForCharacters(element string, chatId int64, gem string) tgbotapi.ReplyKeyboardMarkup {
	var inlineButtons []tgbotapi.KeyboardButton
	list := m.getCharactersListByPage(element, chatId)

	if gem != "" {
		inlineButtons = m.getInlineButtons([]string{gem})
	}

	buttons := m.getCharactersButtons(list, chatId, element)

	return m.getResizedKeyboard(inlineButtons, buttons, CountButtonsInRow)
}

func (m *Manager) getCharactersButtons(list []string, chatId int64, element string) []tgbotapi.KeyboardButton {
	buttons := make([]tgbotapi.KeyboardButton, 0)

	for _, character := range list {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(character))
	}

	// Previous Page button. Example: "⬅ 🔶 Гео"
	if !m.pager.IsFirstPage(chatId) {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(fmt.Sprintf(
			"%s %1s",
			m.renderer.PreviousPageEmoji,
			m.renderer.AddEmojiToElement(element),
		)))
	}

	buttons = append(buttons, tgbotapi.NewKeyboardButton(ReplyKeyboardTextToAllElements))

	// Next Page Button. Example: "❄ Крио ➡"
	if m.pager.IsFirstPage(chatId) && m.pager.HasToPaginate(element, chatId) {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(fmt.Sprintf(
			"%s %s",
			m.renderer.AddEmojiToElement(element),
			m.renderer.NextPageEmoji,
		)))
	}

	return buttons
}

func (m *Manager) getCharactersListByPage(element string, chatId int64) []string {
	first, last := m.pager.GetPositions(chatId, element)

	return m.service.Characters.GetElementCharacters(element, first, last)
}

func (m *Manager) getInlineButtons(list []string) []tgbotapi.KeyboardButton {
	buttons := make([]tgbotapi.KeyboardButton, 0)

	for _, ib := range list {
		buttons = append(buttons, tgbotapi.NewKeyboardButton(ib))
	}

	return buttons
}
