package web_crawler

import (
	"fmt"
	"net/url"
	"sync"
	"web_crawler/pkg/html_parser"

	log "github.com/sirupsen/logrus"
)

type Crawler struct {
	wg sync.WaitGroup

	baseUrl       string
	baseUrlDomain string
	depth         uint

	crawledPages  Store
	webHTMLParser html_parser.HTMLParser
}

// NewCrawler creates a new web crawler instance.
func NewCrawler(baseUrl string, depth uint, webHTMLParser html_parser.HTMLParser) (*Crawler, error) {
	baseUrlParsed, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base url: %v", err)
	}

	domain, err := GetHostDomain(baseUrlParsed.Host)
	if err != nil {
		return nil, fmt.Errorf("failed to get host domain from url %s: %v", baseUrl, err)
	}

	crawler := &Crawler{
		wg: sync.WaitGroup{},

		baseUrl:       baseUrl,
		baseUrlDomain: domain,
		depth:         depth,

		crawledPages:  NewPagesLinksStore(),
		webHTMLParser: webHTMLParser,
	}
	return crawler, nil
}

// StartCrawling starts the web crawling process.
func (c *Crawler) StartCrawling() {
	c.wg.Add(1)
	go c.crawl(c.baseUrl, c.depth)
	c.wg.Wait()
}

// GetCrawledPages returns the list of crawled pages from the base url.
func (c *Crawler) GetCrawledPages() []string {
	return c.crawledPages.GetItems()
}

// crawl recursively crawls the web pages starting from an url, with a maximum depth.
func (c *Crawler) crawl(url string, depth uint) {
	defer c.wg.Done()

	if depth == 0 {
		// we reached the maximum depth.
		return
	}

	if !c.crawledPages.AddItem(url) {
		// the url was already crawled.
		log.Debugf("url %s was already crawled", url)
		return
	}

	pageLinks, err := c.webHTMLParser.GetPageLinks(url)
	if err != nil {
		log.Warnf("failed to get %s page links: %v", url, err)
		return
	}

	for _, link := range c.getSanitizedLinks(pageLinks) {
		c.wg.Add(1)
		go c.crawl(link, depth-1)
	}
}

// getSanitizedLinks returns the sanitized links from a list of links.
func (c *Crawler) getSanitizedLinks(links []string) []string {
	sanitizedLinks := []string{}
	for _, link := range links {
		link := c.sanitizeLink(link)
		if link != "" {
			sanitizedLinks = append(sanitizedLinks, link)
		}
	}
	return sanitizedLinks
}

// sanitizeLink returns the sanitized link.
// A link is considered sanitized if it's a relative link, or an absolute link
// from the same domain. If the link doesn't meet these conditions, it's
// considered invalid and an empty string is returned.
func (c *Crawler) sanitizeLink(link string) string {
	// NOTE: "url.Parse" provides a weaker URL parsing than "url.ParseRequestURI".
	// We will use this to determine if the link is relative or absolute.
	linkParsed, err := url.Parse(link)
	if err != nil {
		log.Warnf("link %s failed url.Parse: %v", link, err)
		return ""
	}

	// If the link is relative, we will join it with the base url.
	if !linkParsed.IsAbs() {
		link, err = url.JoinPath(c.baseUrl, link)
		if err != nil {
			log.Warnf("failed to join url %s with %s: %v", c.baseUrl, link, err)
			return ""
		}
	}

	// NOTE: This is a stronger URL parsing than "url.Parse". We will use this
	// to validate that we are only crawling links from the same domain.
	// It also provides a better way to validate that we have a valid link.
	linkParsed, err = url.ParseRequestURI(link)
	if err != nil {
		log.Warnf("link %s failed url.ParseRequestURI: %v", link, err)
		return ""
	}

	domain, err := GetHostDomain(linkParsed.Host)
	if err != nil {
		log.Warnf("failed to get host domain from absolut url %s: %v", link, err)
		return ""
	}

	if domain == c.baseUrlDomain {
		return link
	}

	// If we reach this point, the link is not valid.
	return ""
}
