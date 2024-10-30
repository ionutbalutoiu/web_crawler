package mocks

import (
	net_url "net/url"
	"os"
	"path/filepath"
	"web_crawler/pkg/utils"

	"golang.org/x/net/html"
)

// MockedHTTPClient is a mocked implementation of the HTTPClient interface from html_parser package.
type MockedHTTPClient struct {
	testDataDir string
}

// NewMockedHTTPClient creates a new instance of MockedHTTPClient.
func NewMockedHTTPClient(testDataDir string) MockedHTTPClient {
	return MockedHTTPClient{
		testDataDir: testDataDir,
	}
}

// LoadPage loads a page from the test data directory. It returns the parsed HTML node of the page.
func (m MockedHTTPClient) LoadPage(url string) (*html.Node, error) {
	parsedUrl, err := net_url.ParseRequestURI(url)
	if err != nil {
		return nil, err
	}

	domain, err := utils.GetHostDomain(parsedUrl.Host)
	if err != nil {
		return nil, err
	}

	testPagePath := filepath.Join(m.testDataDir, domain, parsedUrl.Path, "index.html")
	f, err := os.Open(testPagePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	node, err := html.Parse(f)
	if err != nil {
		return nil, err
	}

	return node, nil
}
