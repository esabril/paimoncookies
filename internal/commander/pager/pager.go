package pager

type Pager struct {
	chatCurrentPage    map[int64]int
	totalElementsByKey map[string]int
	maxValuesOnMenu    int
}

func NewPager(data map[string][]string, maxValues int) *Pager {
	p := &Pager{
		chatCurrentPage:    map[int64]int{},
		totalElementsByKey: map[string]int{},
		maxValuesOnMenu:    maxValues,
	}

	p.countTotalPages(data)

	return p
}

func (p *Pager) IsFirstPage(chatId int64) bool {
	return p.CurrentPage(chatId) == 1
}

func (p *Pager) HasToPaginate(key string, chatId int64) bool {
	return (p.CurrentPage(chatId) * p.maxValuesOnMenu) < p.getTotal(key)
}

func (p *Pager) SetCurrentPage(page int, chatId int64) {
	p.chatCurrentPage[chatId] += page

	if p.chatCurrentPage[chatId] < 1 {
		p.chatCurrentPage[chatId] = 1
	}
}

func (p *Pager) CurrentPage(chatId int64) int {
	page, ok := p.chatCurrentPage[chatId]
	if !ok {
		p.chatCurrentPage[chatId] = 1
		return 1
	}

	return page
}

// GetPositions for slices with immutable data for current page
func (p *Pager) GetPositions(chatId int64, key string) (first int, last int) {
	curPage := p.CurrentPage(chatId)
	first = (curPage - 1) * p.maxValuesOnMenu

	total := p.getTotal(key)
	last = first + p.maxValuesOnMenu

	if last > total {
		last = total
	}

	return
}

func (p *Pager) Flush(chatId int64) {
	delete(p.chatCurrentPage, chatId)
}

func (p *Pager) getTotal(key string) int {
	return p.totalElementsByKey[key]
}

func (p *Pager) countTotalPages(data map[string][]string) {
	for key, values := range data {
		p.totalElementsByKey[key] = len(values)
	}
}
