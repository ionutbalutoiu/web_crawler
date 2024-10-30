package web_crawler

// Generic interface for a store of items
type Store interface {
	GetItems() []string
	AddItem(item string) bool
}
