package web_crawler

import (
	"sync"
)

// CrawledPagesStore is a structure for crawled pages links.
type CrawledPagesStore struct {
	m            sync.Mutex
	crawledPages map[string][]string
}

// NewCrawledPagesStore creates a new CrawledPagesStore.
func NewCrawledPagesStore() Store {
	return &CrawledPagesStore{
		m:            sync.Mutex{},
		crawledPages: map[string][]string{},
	}
}

// GetItems returns all the crawled pages.
func (s *CrawledPagesStore) GetItems() map[string][]string {
	return s.crawledPages
}

// AddItem adds a new crawled page to the store.
// It returns true, if the page is successfully added to the store.
// Otherwise, the function returns false, without doing anything else.
func (s *CrawledPagesStore) AddItem(pageUrl string, pageLinks []string) bool {
	s.m.Lock()
	defer s.m.Unlock()
	if s.ExistsItem(pageUrl) {
		return false
	}
	s.crawledPages[pageUrl] = pageLinks
	return true
}

// UpdateItem updates the links of a page in the store.
func (s *CrawledPagesStore) UpdateItem(pageUrl string, pageLinks []string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.crawledPages[pageUrl] = pageLinks
}

// RemoveItem removes a page from the store.
func (s *CrawledPagesStore) RemoveItem(pageUrl string) {
	s.m.Lock()
	defer s.m.Unlock()
	delete(s.crawledPages, pageUrl)
}

// ExistsItem returns true if the page URL exists in the store.
func (s *CrawledPagesStore) ExistsItem(pageUrl string) bool {
	_, ok := s.crawledPages[pageUrl]
	return ok
}
