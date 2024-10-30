# Go Web Crawler

This is a simple web crawler, written in Go, that navigates a website URL and prints the URLs it discovers on the page. It recursively crawls page URLs up to a specified depth.

## Build

To build the project, run the following command:

```shell
make build
```

A binary will be created in the `build` directory.

## Usage

```shell
$ ./build/web_crawler -help
Usage of ./build/web_crawler:
  -depth uint
        Crawl depth (default: 3) (default 3)
  -log-level string
        Log level (default "info")
  -output string
        Output format (default: shell). Valid values shell, json (default "shell")
  -url string
        Base URL to crawl (default: https://monzo.com) (default "https://monzo.com")
```

### Examples

Show the crawling results in JSON format:

```shell
$ ./build/web_crawler -url https://monzo.com -depth 3 -output=json
{
  "https://monzo.com": [
    "https://monzo.com/%23mainContent",
    "https://monzo.com",
    "https://monzo.com/current-account"
    ...
    "https://monzo.com/information-about-current-account-services",
    "https://monzo.com/service-information"
  ],
  "https://monzo.com/%23mainContent": [
    "https://monzo.com/%23mainContent",
    "https://monzo.com",
    "https://monzo.com/current-account",
    ...
    "https://monzo.com/information-about-current-account-services",
    "https://monzo.com/service-information"
  ],
  ...
  "https://monzo.com/-deeplinks/cashback?": [
    "https://monzo.com/help/app-help/hiddenappmagiclink"
  ],
  "https://monzo.com/-deeplinks/connect-mortgage": [
    "https://monzo.com/help/app-help/hiddenappmagiclink"
  ...
  ...
  ...
}
```

The JSON output format is:

```json
{
  "CRAWLED_PAGE_1_URL": [
    "PAGE_1_URL_1",
    "PAGE_1_URL_2",
    ...
  ],
  "CRAWLED_PAGE_2_URL": [
    "PAGE_2_URL_1",
    "PAGE_2_URL_2",
    ...
  ],
  ...
}
```

To show only the crawled pages' URLs, pipe the JSON output command into `jq`:

```shell
$ ./build/web_crawler -url https://monzo.com -depth 3 -output=json | jq 'keys'
[
  "https://monzo.com",
  "https://monzo.com/%23mainContent",
  "https://monzo.com/-deeplinks/cashback",
  "https://monzo.com/-deeplinks/connect-mortgage",
  "https://monzo.com/-deeplinks/create-instant-access-pot",
  "https://monzo.com/-deeplinks/create_pot",
  ...
  ...
  ...
  "https://monzo.com/us/money/see-it-all",
  "https://monzo.com/us/money/spend-confidently",
  "https://monzo.com/us/personal-account",
  "https://monzo.com/us/savings"
]
```

## Tests

To run the tests, execute the following command:

```shell
make test
```

This will use a mocked HTTP client to test the web crawler.

## Limitations / Trade-offs

These are some of the current limitations of the web crawler implementations:

1. The web crawler spawns a goroutine for each URL it discovers (up to the specified depth). This can result in a high number of goroutines being created, leading to significant resource consumption.
   1. A more elegant solution would be to use a worker pool to limit the number of concurrent goroutines, though this would require additional implementation time.
1. HTTP calls are not retried if they fail, which can result in incomplete results if a page fails to load.
   1. Since HTTP calls can fail due to transient network issues, it would be beneficial to retry the calls a few times before giving up. This is crucial for a production-ready web crawler.
1. The web crawler needs a throttling mechanism to prevent it from overwhelming the target website.
   1. A throttling mechanism would prevent the crawler from making too many requests in a short period, which could lead to the target website blocking or rate-limiting the crawler.
