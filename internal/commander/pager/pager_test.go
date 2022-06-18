package pager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPager(t *testing.T) {
	data := map[string][]string{
		"One":   {"1", "2", "3"},
		"Two":   {"1"},
		"Three": {"5", "7"},
	}
	expectedTotalPages := map[string]int{
		"One":   3,
		"Two":   1,
		"Three": 2,
	}

	p := NewPager(data, 2)

	assert.Equal(t, 2, p.maxValuesOnMenu)
	assert.Equal(t, expectedTotalPages, p.totalElementsByKey)
}

func TestPager_IsFirstPage(t *testing.T) {
	t.Parallel()

	p := Pager{
		chatCurrentPage: map[int64]int{
			123: 1,
			321: 2,
		},
	}

	testCases := []struct {
		ChatId   int64
		Expected bool
	}{
		{123, true},
		{321, false},
		{456, true},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("ChatId: %d", tt.ChatId), func(t *testing.T) {
			assert.Equal(t, tt.Expected, p.IsFirstPage(tt.ChatId))
		})
	}
}

func TestPager_HasToPaginate(t *testing.T) {
	p := Pager{
		totalElementsByKey: map[string]int{
			"One":   3,
			"Two":   2,
			"Three": 5,
		},
		chatCurrentPage: map[int64]int{
			123: 1,
			321: 2,
		},
	}

	testCases := []struct {
		MaxValues int
		ChatId    int64
		Key       string
		Expected  bool
	}{
		{1, 123, "One", true},
		{3, 123, "One", false},
		{2, 456, "Two", false},
		{3, 123, "Two", false},
		{2, 321, "Three", true},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("ChatId: %d, Key: %s", tt.ChatId, tt.Key), func(t *testing.T) {
			p.maxValuesOnMenu = tt.MaxValues
			assert.Equal(t, tt.Expected, p.HasToPaginate(tt.Key, tt.ChatId))
		})
	}
}

func TestPager_SetCurrentPage(t *testing.T) {
	t.Parallel()

	p := Pager{
		chatCurrentPage: map[int64]int{
			123: 1,
		},
	}

	testCases := []struct {
		ChatId   int64
		Page     int
		Expected int
	}{
		{123, 1, 2},
		{123, -1, 1},
		{123, -1, 1},
		{123, 1, 2},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Page %d", tt.Page), func(t *testing.T) {
			p.SetCurrentPage(tt.Page, tt.ChatId)
			assert.Equal(t, tt.Expected, p.CurrentPage(tt.ChatId))
		})
	}
}

func TestPager_GetPositions(t *testing.T) {
	t.Parallel()

	var chatId int64 = 123
	key := "One"

	p := Pager{
		totalElementsByKey: map[string]int{
			"One": 7,
		},
		chatCurrentPage: map[int64]int{
			chatId: 1,
		},
	}

	testCases := []struct {
		MaxValues     int
		CurrentPage   int
		ExpectedFirst int
		ExpectedLast  int
	}{
		{2, 1, 0, 2},
		{3, 2, 3, 6},
		{3, 3, 6, 7},
		{4, 2, 4, 7},
		{1, 2, 1, 2},
		{5, 2, 5, 7},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("Max values: %d, Current Page: %d", tt.MaxValues, tt.CurrentPage), func(t *testing.T) {
			p.maxValuesOnMenu = tt.MaxValues
			p.chatCurrentPage[chatId] = tt.CurrentPage

			f, l := p.GetPositions(chatId, key)

			assert.Equal(t, tt.ExpectedFirst, f)
			assert.Equal(t, tt.ExpectedLast, l)
		})
	}
}

func TestPager_Flush(t *testing.T) {
	t.Parallel()

	p := Pager{
		chatCurrentPage: map[int64]int{
			123: 1,
		},
	}

	testCases := []struct {
		ChatId int64
		Exists bool
	}{
		{123, true},
		{321, false},
	}

	for _, tt := range testCases {
		t.Run(fmt.Sprintf("%d", tt.ChatId), func(t *testing.T) {
			_, ok := p.chatCurrentPage[tt.ChatId]
			assert.Equal(t, tt.Exists, ok)

			p.Flush(tt.ChatId)

			_, ok = p.chatCurrentPage[tt.ChatId]
			assert.False(t, ok)
		})
	}
}
