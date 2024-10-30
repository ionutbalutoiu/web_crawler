package web_crawler

// Generic interface for a store of crawled pages, and their URLs.
type Store interface {
	GetItems() map[string][]string
	AddItem(pageUrl string, pageLinks []string)
	UpdateItem(pageUrl string, pageLinks []string)
	ExistsItem(pageUrl string) bool
}
