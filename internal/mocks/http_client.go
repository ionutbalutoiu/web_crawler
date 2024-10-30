package mocks

import (
	net_url "net/url"
	"os"
	"path/filepath"

	"web_crawler/pkg/web_crawler"

	"golang.org/x/net/html"
)

type MockedHTTPClient struct {
	testDataDir string
}

func NewMockedHTTPClient(testDataDir string) MockedHTTPClient {
	return MockedHTTPClient{
		testDataDir: testDataDir,
	}
}

func (m MockedHTTPClient) LoadPage(url string) (*html.Node, error) {
	parsedUrl, err := net_url.ParseRequestURI(url)
	if err != nil {
		return nil, err
	}

	domain, err := web_crawler.GetHostDomain(parsedUrl.Host)
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
