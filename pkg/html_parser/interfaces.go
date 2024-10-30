package html_parser

import "golang.org/x/net/html"

// Generic interface for an HTTP client
type HTTPClient interface {
	LoadPage(url string) (*html.Node, error)
}

// Generic interface for an HTML parser
type HTMLParser interface {
	GetPageLinks(url string) ([]string, error)
}
