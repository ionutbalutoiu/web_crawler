package main

import (
	"flag"
	"fmt"
	"web_crawler/pkg/html_parser"
	"web_crawler/pkg/web_crawler"

	log "github.com/sirupsen/logrus"
)

var (
	logLevel = flag.String("log-level", "info", "Log level")
	baseUrl  = flag.String("url", "https://monzo.com", "Base URL to crawl (default: https://monzo.com)")
	depth    = flag.Uint("depth", 3, "Crawl depth (default: 3)")
)

func initLogging() error {
	lvl, err := log.ParseLevel(*logLevel)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}

func main() {
	flag.Parse()

	if err := initLogging(); err != nil {
		log.Fatalf("error setting log level: %v", err)
	}

	httpClient := html_parser.NewWebClient()
	webHtmlParser := html_parser.NewWebHTMLParser(httpClient)

	webCrawler, err := web_crawler.NewCrawler(*baseUrl, *depth, webHtmlParser)
	if err != nil {
		log.Fatalf("error creating web crawler: %v", err)
	}

	fmt.Printf("Crawling page %s with depth %d\n", *baseUrl, *depth)
	webCrawler.StartCrawling()

	fmt.Println("Links crawled:")
	for _, link := range webCrawler.GetCrawledPages() {
		fmt.Println(link)
	}
}
