package keyboard

import (
	"github.com/esabril/paimoncookies/internal/commander/pager"
	"github.com/esabril/paimoncookies/internal/service"
	"github.com/esabril/paimoncookies/tools/renderer"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
)

const (
	CountButtonsInRow = 3
)

type Manager struct {
	service  *service.Service
	pager    *pager.Pager
	renderer *renderer.Renderer
}

func NewManager(s *service.Service, p *pager.Pager, r *renderer.Renderer) *Manager {
	return &Manager{
		service:  s,
		pager:    p,
		renderer: r,
	}
}

// SetPageFromReply Handle messages like "💧 Гидро ➡" and "⬅ 🔶 Гео"
func (m *Manager) SetPageFromReply(reply string, chatId int64) int {
	data := strings.Split(reply, " ")

	// Not a pagination message: we on the first page
	if len(data) < 3 {
		return 1
	}

	// return to previous page
	if data[0] == m.renderer.PreviousPageEmoji {
		m.pager.SetCurrentPage(-1, chatId)
	}

	// go to next page
	if data[2] == m.renderer.NextPageEmoji {
		m.pager.SetCurrentPage(1, chatId)
	}

	return m.pager.CurrentPage(chatId)
}

// FlushPager Remove chat data from pager
func (m *Manager) FlushPager(chatId int64) {
	m.pager.Flush(chatId)
}

func (m *Manager) getKeyboard(inlineButtons, buttons []tgbotapi.KeyboardButton, countButtonsInRow int) tgbotapi.ReplyKeyboardMarkup {
	rows := m.getRows(inlineButtons, buttons, countButtonsInRow)

	return tgbotapi.NewReplyKeyboard(rows...)
}

func (m *Manager) getResizedKeyboard(
	inlineButtons,
	buttons []tgbotapi.KeyboardButton,
	countButtonsInRow int,
) tgbotapi.ReplyKeyboardMarkup {
	kb := m.getKeyboard(inlineButtons, buttons, countButtonsInRow)
	kb.ResizeKeyboard = true

	return kb
}

func (m *Manager) getRows(inlineButtons, buttons []tgbotapi.KeyboardButton, countButtonsInRow int) [][]tgbotapi.KeyboardButton {
	rows := make([][]tgbotapi.KeyboardButton, 0)

	if len(inlineButtons) > 0 {
		for _, ib := range inlineButtons {
			row := []tgbotapi.KeyboardButton{ib}
			rows = append(rows, row)
		}
	}

	count := len(buttons)
	first, last := 0, countButtonsInRow

	if last >= count {
		last = count
	}

	for last <= count {
		row := buttons[first:last]
		rows = append(rows, row)
		first = last
		last += countButtonsInRow
	}

	rows = append(rows, buttons[first:])

	return rows
}
