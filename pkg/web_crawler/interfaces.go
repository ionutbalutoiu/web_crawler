package web_crawler

// Generic interface for a store of crawled pages.
type Store interface {
	GetItems() map[string][]string
	AddItem(pageUrl string, pageLinks []string) bool
	UpdateItem(pageUrl string, pageLinks []string)
	ExistsItem(pageUrl string) bool
}
