package main

import (
	"fmt"
	"os"
)

/*
	Todos for future development:
	- Figure out go concurrency for deeper crawl
	- Sitemap as XML with layout, rather than just txt links
	- Integration Test
	- Acceptance tests for each func
 */

func main() {
	fmt.Println("Welcome to spiderWeb 🕷")

	// Request URL
	url := userInput("Please enter a valid URL to scan: ")

	// Validate point of entry URL
	isValid, validUrl := validateRequest(url)

	if isValid {
		// Crawl
		fmt.Println("Crawling " + validUrl + " to a depth of three...\nThis operation take up to 10 minutes.\n")
		links := crawlSite(validUrl)

		if len(links) > 1 {
			fmt.Println("Saving to file...")
			saveToFile(formatToString(links))
			fmt.Println("Saved to links.txt. Goodbye!")
		} else {
			fmt.Println("Your site does not appear to link out to any other sitepages, so there are no links to see. Goodbye!")
		}

		os.Exit(1)
	} else {
		fmt.Printf("We could not reach %s, so will be unable to crawl it. Please try again later.\n", url)
		os.Exit(13) // 13 = sys code for invalid data
	}
}
