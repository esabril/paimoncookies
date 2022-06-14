package pager

import "sync"

const MaxDataValuesOnMenu = 6

type Pager struct {
	// Chats characters map mutex
	dmu      sync.RWMutex
	chatData map[int64]int
	// Total pages map mutex
	pmu        sync.RWMutex
	totalPages map[string]int
}

func NewPager(data map[string][]string) *Pager {
	p := &Pager{
		chatData:   map[int64]int{},
		totalPages: map[string]int{},
	}

	p.countTotalPages(data)

	return p
}

func (p *Pager) IsFirstPage(chatId int64) bool {
	return p.CurrentPage(chatId) == 1
}

func (p *Pager) HasToPaginate(key string, chatId int64) bool {
	return (p.CurrentPage(chatId) * MaxDataValuesOnMenu) < p.getTotal(key)
}

func (p *Pager) SetCurrentPage(page int, chatId int64) {
	p.dmu.RLock()
	defer p.dmu.RUnlock()

	p.chatData[chatId] += page
}

func (p *Pager) CurrentPage(chatId int64) int {
	p.dmu.RLock()
	defer p.dmu.RUnlock()

	page, ok := p.chatData[chatId]
	if !ok {
		p.chatData[chatId] = 1
		return 1
	}

	return page
}

// GetPositions for slices with immutable data for current page
func (p *Pager) GetPositions(chatId int64, key string) (first int, last int) {
	curPage := p.CurrentPage(chatId)
	first = (curPage - 1) * MaxDataValuesOnMenu

	total := p.getTotal(key)
	last = first + MaxDataValuesOnMenu

	if last > total {
		last = total
	}

	return
}

func (p *Pager) Flush(chatId int64) {
	p.dmu.RLock()
	defer p.dmu.RUnlock()

	delete(p.chatData, chatId)
}

func (p *Pager) getTotal(key string) int {
	p.pmu.RLock()
	defer p.pmu.RUnlock()

	return p.totalPages[key]
}

func (p *Pager) countTotalPages(data map[string][]string) {
	for key, values := range data {
		p.totalPages[key] = len(values)
	}
}
