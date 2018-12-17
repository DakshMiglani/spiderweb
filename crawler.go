package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func crawlLink(rootURL string, givenURL string) []string {
	// Normalise Given URL
	givenURL = normalisePath(givenURL)

	// Make GET request
	resp, err := http.Get(givenURL)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	body := resp.Body
	defer body.Close()

	dom, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	// Extract links
	links := []string{}

	// Find a tags and parse the links.
	dom.Find("a").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")

		if ok {
			link = getHref(link, givenURL)
			if len(link) > 1 {
				link = normalisePath(link)
				if strings.Contains(link, rootURL) {
					links = append(links, link)
				}
			}
		}
	})

	return links
}

func crawlSite(rootUrl string, depth *int) []string {
	// Default to 2 depth - 1 call
	var maxDepth int
	if depth == nil {
		maxDepth = 1
	} else {
		maxDepth = *depth - 1
	}

	// Begin by crawling root
	links := crawlSiteForLinks(rootUrl, rootUrl, nil, maxDepth, 0)

	// Return multiple links
	return links
}

func crawlSiteForLinks(rootUrl string, link string, givenLinks *[]string, maxDepth int, currentDepth int) []string {
	var links []string
	if givenLinks == nil {
		links = []string{}
	} else {
		links = *givenLinks
	}

	crawledLinks := crawlLink(rootUrl, link)
	for _, crawledLink := range crawledLinks {
		if !stringInSlice(links, crawledLink) {
			links = append(links, crawledLink)
		}
		if maxDepth > currentDepth {
			links = crawlSiteForLinks(rootUrl, crawledLink, &links, maxDepth, currentDepth+1)
		}
	}
	return links
}

// Gets href attribute from a token
func getHref(value string, url string) string {
	if len(value) > 0 {
		if value[len(value)-1:] == "/" {
			value = value[:len(value)-1]
		}

		if len(value) == 1 && (string(value[0]) == "/" || string(value[0]) == "#") {
			return url
		}

		if len(value) > 1 && string(value[0]) == "/" {
			value = url + value
		}

		if strings.Contains(value, "#") {
			return value[strings.Index(value, "#"):]
		}

		return value
	}

	return ""
}
