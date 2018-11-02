package main

import (
	"fmt"
	"os"
)

/*
	Todos for future development:
	- Figure out go concurrency for faster crawl
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
		depth := intUserInput("What depth would you like to run to? (2): ")
		fmt.Println("Crawling " + validUrl + "...\nThis operation take up to 5 minutes.\n")
		links := crawlSite(validUrl, depth)

		if len(links) > 1 {
			fmt.Println("Saving to files...")
			saveToFile(formatToString(links))
			fmt.Println("1 - Saved to links.txt")
			saveAsXml(links)
			fmt.Println("2 - Saved to webmap.xml")
			fmt.Println("Goodbye!")
		} else {
			fmt.Println("Your site does not appear to link out to any other sitepages, so there are no links to see. Goodbye!")
		}

		os.Exit(1)
	} else {
		fmt.Printf("We could not reach %s, so will be unable to crawl it. Please try again later.\n", url)
		os.Exit(13) // 13 = sys code for invalid data
	}
}
