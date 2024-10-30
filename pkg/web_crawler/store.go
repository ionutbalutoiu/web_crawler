package web_crawler

import (
	"slices"
	"sync"
)

// PageLinksStore is a store for crawled pages links.
type PagesLinksStore struct {
	m          sync.Mutex
	PagesLinks []string
}

// NewPagesLinksStore creates a new PagesLinksStore.
func NewPagesLinksStore() Store {
	return &PagesLinksStore{}
}

// AddItem adds a link to the store if it is not already present.
// It returns true if the link was added, false otherwise.
func (s *PagesLinksStore) AddItem(item string) bool {
	s.m.Lock()
	defer s.m.Unlock()

	if slices.Contains(s.PagesLinks, item) {
		return false
	}

	s.PagesLinks = append(s.PagesLinks, item)
	return true
}

// GetItems returns all the links in the store.
func (s *PagesLinksStore) GetItems() []string {
	return s.PagesLinks
}
