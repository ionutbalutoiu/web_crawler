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
		crawledPages: make(map[string][]string),
	}
}

// GetItems returns all the links in the store.
func (s *CrawledPagesStore) GetItems() map[string][]string {
	s.m.Lock()
	defer s.m.Unlock()
	return s.crawledPages
}

// AddItem add the links of a page in the store.
func (s *CrawledPagesStore) AddItem(pageUrl string, pageLinks []string) {
	s.UpdateItem(pageUrl, pageLinks)
}

// UpdateItem updates the links of a page in the store.
func (s *CrawledPagesStore) UpdateItem(pageUrl string, pageLinks []string) {
	s.m.Lock()
	defer s.m.Unlock()
	s.crawledPages[pageUrl] = pageLinks
}

// ExistsItem returns true if the page URL exists in the store.
func (s *CrawledPagesStore) ExistsItem(pageUrl string) bool {
	s.m.Lock()
	defer s.m.Unlock()
	_, ok := s.crawledPages[pageUrl]
	return ok
}
