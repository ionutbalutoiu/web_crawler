package html_parser_test

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"syscall"
	"testing"
	"web_crawler/internal/mocks"
	"web_crawler/pkg/html_parser"

	"github.com/stretchr/testify/suite"
)

type HTMLParserTestSuite struct {
	suite.Suite

	testDataDir string
	htmlParser  html_parser.HTMLParser
}

func (s *HTMLParserTestSuite) SetupTest() {
	testDataDir, err := filepath.Abs(filepath.Join("..", "..", "internal", "mocks", "testdata"))
	if err != nil {
		s.FailNow("error getting test data path: %v", err)
	}
	mockedHttpClient := mocks.NewMockedHTTPClient(testDataDir)

	s.testDataDir = testDataDir
	s.htmlParser = html_parser.NewWebHTMLParser(mockedHttpClient)
}

func (s *HTMLParserTestSuite) TestGetPageLinks() {
	testCases := []struct {
		Name          string
		Url           string
		ExpectedLinks []string
		Err           error
	}{
		{
			Name: "HTML_With_Valid_Links",
			Url:  "http://localhost",
			ExpectedLinks: []string{
				"http://localhost",
				"http://localhost/path1_1",
				"http://localhost/path1_2",
				"http://localhost/path1_3",
			},
			Err: nil,
		},
		{
			Name: "HTML_With_Complex_Links_Formats",
			Url:  "http://example.com/complex_links_formats",
			ExpectedLinks: []string{
				"http://example.com",
			},
			Err: nil,
		},
		{
			Name:          "HTML_Without_Links",
			Url:           "http://example.com",
			ExpectedLinks: []string{},
			Err:           nil,
		},
		{
			Name:          "HTML_Page_Not_Found",
			Url:           "http://example.com/not_found",
			ExpectedLinks: nil,
			Err:           &fs.PathError{Op: "open", Path: fmt.Sprintf("%s/example.com/not_found/index.html", s.testDataDir), Err: syscall.Errno(0x2)},
		},
	}

	for _, tc := range testCases {
		s.Run(tc.Name, func() {
			links, err := s.htmlParser.GetPageLinks(tc.Url)
			s.Equal(tc.Err, err)
			if err == nil {
				s.Equal(tc.ExpectedLinks, links)
			}
		})
	}
}

func TestHTMLParserTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(HTMLParserTestSuite))
}
