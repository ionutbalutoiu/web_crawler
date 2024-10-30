package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"web_crawler/pkg/html_parser"
	"web_crawler/pkg/web_crawler"

	log "github.com/sirupsen/logrus"
)

var (
	baseUrl      = flag.String("url", "https://monzo.com", "Base URL to crawl (default: https://monzo.com)")
	depth        = flag.Uint("depth", 3, "Crawl depth (default: 3)")
	logLevel     = flag.String("log-level", "info", "Log level")
	outputFormat = flag.String("output", "json", "Output format (default: json). Valid values json, plaintext")
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

	// Start crawling
	webCrawler.StartCrawling()

	// Get the crawled pages
	crawledPages := webCrawler.GetCrawledPages()

	switch *outputFormat {
	case "plaintext":
		fmt.Println("Pages crawled:")
		for page, links := range crawledPages {
			fmt.Printf("Page URL: %s\n", page)
			if len(links) > 0 {
				fmt.Printf("Page Links:\n")
				for _, link := range links {
					fmt.Printf(" - %s\n", link)
				}
			}
		}
	case "json":
		jsonOutput, err := json.MarshalIndent(crawledPages, "", "  ")
		if err != nil {
			log.Fatalf("error marshalling JSON: %v", err)
		}
		fmt.Println(string(jsonOutput))
	}
}
