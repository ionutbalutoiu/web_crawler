package html_parser

import (
	"slices"
	"strings"

	"golang.org/x/net/html"
)

// WebHTMLParser is an implementation of the HTMLParser interface.
type WebHTMLParser struct {
	httpClient HTTPClient
}

// NewWebHTMLParser creates a new WebHTMLParser instance (which implements
// HTMLParser interface).
func NewWebHTMLParser(httpClient HTTPClient) HTMLParser {
	return &WebHTMLParser{httpClient: httpClient}
}

// GetPageLinks returns all the links in the HTML page after parsing it.
func (p *WebHTMLParser) GetPageLinks(url string) ([]string, error) {
	htmlDoc, err := p.httpClient.LoadPage(url)
	if err != nil {
		return nil, err
	}
	return p.getHTMLDocLinks(htmlDoc), nil
}

// getHTMLDocLinks returns all the links in the HTML document.
// It recursively traverses the HTML document structure and returns all the
// links as given in the "href" attribute of the "<a ...> HTML tag.
func (p *WebHTMLParser) getHTMLDocLinks(htmlNode *html.Node) []string {
	links := []string{}
	if htmlNode.Data == HTMLTagLink {
		hrefValue := p.getHrefAttrValue(htmlNode)
		if hrefValue != "" && !slices.Contains(links, hrefValue) {
			links = append(links, hrefValue)
		}
	}
	for child := htmlNode.FirstChild; child != nil; child = child.NextSibling {
		for _, link := range p.getHTMLDocLinks(child) {
			if !slices.Contains(links, link) {
				links = append(links, link)
			}
		}
	}
	return links
}

// getHrefAttrValue returns the value of the "href" attribute of the HTML node.
// If the "href" attribute is not present, it returns an empty string.
func (p *WebHTMLParser) getHrefAttrValue(htmlNode *html.Node) string {
	for _, attr := range htmlNode.Attr {
		if attr.Key == HTMLTagLinkHrefAttr {
			// Trim the leading and trailing spaces from the attribute value.
			link := strings.TrimSpace(attr.Val)
			// Trim the following trailing characters from the link:
			// - "/" (slashes)
			// - "#"
			// - "%23" (which is "#" encoded).
			// - "?" (empty query parameters)
			for strings.HasSuffix(link, "/") || strings.HasSuffix(link, "#") || strings.HasSuffix(link, "%23") || strings.HasSuffix(link, "?") {
				link = strings.TrimSuffix(link, "/")
				link = strings.TrimSuffix(link, "#")
				link = strings.TrimSuffix(link, "%23")
				link = strings.TrimSuffix(link, "?")
			}
			return link
		}
	}
	return ""
}
