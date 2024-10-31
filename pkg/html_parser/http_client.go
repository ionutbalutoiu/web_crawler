package html_parser

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

// WebClient is an implementation of the HTTPClient interface.
type WebClient struct{}

// NewWebClient creates a new WebClient instance.
func NewWebClient() HTTPClient {
	return &WebClient{}
}

// LoadPage loads the HTML page from the given URL (by doing HTTP GET), and returns the parsed HTML node.
func (w WebClient) LoadPage(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check for non-200 status code
	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("failed to get %s url: got status code %d)", url, resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
