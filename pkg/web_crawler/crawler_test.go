package web_crawler_test

import (
	"fmt"
	"path/filepath"
	"testing"
	"web_crawler/internal/mocks"
	"web_crawler/pkg/html_parser"
	"web_crawler/pkg/web_crawler"

	"github.com/stretchr/testify/suite"
)

type CrawlerTestSuite struct {
	suite.Suite

	mockedHtmlParser *html_parser.HTMLParser
}

func (s *CrawlerTestSuite) SetupTest() {
	testDataDir, err := filepath.Abs(filepath.Join("..", "..", "testdata", "mocks", "http_client"))
	if err != nil {
		s.FailNow("error getting test data path: %v", err)
	}
	mockedHttpClient := mocks.NewMockedHTTPClient(testDataDir)
	mockedHtmlParser := html_parser.NewWebHTMLParser(mockedHttpClient)

	s.mockedHtmlParser = &mockedHtmlParser
}

func (s *CrawlerTestSuite) TestCrawledPages() {
	testCases := []struct {
		Name                 string
		BaseUrl              string
		Depth                uint
		ExpectedCrawledPages map[string][]string
		err                  error
	}{
		{
			Name:    "Depth_1",
			BaseUrl: "http://localhost",
			Depth:   1,
			ExpectedCrawledPages: map[string][]string{
				"http://localhost": {
					"http://localhost",
					"http://localhost/path1_1",
					"http://localhost/path1_2",
					"http://localhost/path1_3",
				},
			},
			err: nil,
		},
		{
			Name:    "Depth_2",
			BaseUrl: "http://localhost",
			Depth:   2,
			ExpectedCrawledPages: map[string][]string{
				"http://localhost": {
					"http://localhost",
					"http://localhost/path1_1",
					"http://localhost/path1_2",
					"http://localhost/path1_3",
				},
				"http://localhost/path1_1": {
					"http://localhost",
					"http://localhost/path1_1/path2_1",
					"http://localhost/path1_1/path2_2",
				},
				"http://localhost/path1_2": {
					"http://localhost",
				},
				"http://localhost/path1_3": {
					"http://localhost",
				},
			},
			err: nil,
		},
		{
			Name:                 "Depth_0",
			BaseUrl:              "http://localhost",
			Depth:                0,
			ExpectedCrawledPages: map[string][]string{},
			err:                  nil,
		},
		{
			Name:    "Depth_100",
			BaseUrl: "http://localhost",
			Depth:   100,
			ExpectedCrawledPages: map[string][]string{
				"http://localhost": {
					"http://localhost",
					"http://localhost/path1_1",
					"http://localhost/path1_2",
					"http://localhost/path1_3",
				},
				"http://localhost/path1_1": {
					"http://localhost",
					"http://localhost/path1_1/path2_1",
					"http://localhost/path1_1/path2_2",
				},
				"http://localhost/path1_2": {
					"http://localhost",
				},
				"http://localhost/path1_3": {
					"http://localhost",
				},
				"http://localhost/path1_1/path2_1": {
					"http://localhost",
					"http://localhost/path1_1/path2_1/path3_1",
				},
				"http://localhost/path1_1/path2_1/path3_1": {
					"http://localhost",
				},
				"http://localhost/path1_1/path2_2": {
					"http://localhost",
				},
			},
			err: nil,
		},
		{
			Name:    "Domain_With_Port",
			BaseUrl: "http://localhost:80",
			Depth:   1,
			ExpectedCrawledPages: map[string][]string{
				"http://localhost:80": {
					"http://localhost",
					"http://localhost/path1_1",
					"http://localhost/path1_2",
					"http://localhost/path1_3",
				},
			},
			err: nil,
		},
		{
			Name:    "Domain_With_Port_Invalid_Format",
			BaseUrl: "http://localhost::80",
			Depth:   1,
			err:     fmt.Errorf("failed to get host domain from url http://localhost::80: address localhost::80: too many colons in address"),
		},
		{
			Name:    "Different_BaseUrl",
			BaseUrl: "http://example.com",
			Depth:   100,
			ExpectedCrawledPages: map[string][]string{
				"http://example.com": {},
			},
			err: nil,
		},
		{
			Name:    "Failed_To_Parse_BaseUrl",
			BaseUrl: "dummy-url",
			Depth:   100,
			err:     fmt.Errorf("failed to parse base url: parse \"dummy-url\": invalid URI for request"),
		},
	}
	for _, tc := range testCases {
		s.Run(tc.Name, func() {
			crawler, err := web_crawler.NewCrawler(tc.BaseUrl, tc.Depth, *s.mockedHtmlParser)
			s.Equal(tc.err, err)
			if err == nil {
				crawler.StartCrawling()
				pages := crawler.GetCrawledPages()
				s.Equal(tc.ExpectedCrawledPages, pages)
			}
		})

	}
}

func TestCrawlerTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CrawlerTestSuite))
}
