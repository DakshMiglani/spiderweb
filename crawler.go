package main

import (
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

func crawlLink(rootUrl string, givenUrl string) []string {
	// Normalise Given URL
	givenUrl = normalisePath(givenUrl)

	// Make GET request
	resp, _ := http.Get(givenUrl)
	body := resp.Body

	// Extract links
	links := []string{}
	page := html.NewTokenizer(body)

	for {
		tokenType := page.Next()

		switch {
		case tokenType == html.StartTagToken:
			token := page.Token()
			if token.Data == "a" {
				link := getHref(token, givenUrl)
				if len(link)>1 {
					link = normalisePath(link)
					if strings.Contains(link, rootUrl) {
						links = append(links, link)
					}
				}
			}
		case tokenType == html.ErrorToken:
			return links
		}
	}

	body.Close()
	return links
}

func crawlSite(rootUrl string) []string {
	// Begin by crawling root
	links := crawlLink(rootUrl, rootUrl)

	for _, link := range links {
		// Crawl to second depth
		secondLinks := crawlLink(rootUrl, link)
		for _, secondLink := range secondLinks {
			// Crawl to third depth
			thirdLinks := crawlLink(rootUrl, link)
			for _, thirdLink := range thirdLinks {
				if !stringInSlice(thirdLink, links) {
					links = append(links, thirdLink)
				}
			}
			if !stringInSlice(secondLink, links) {
				links = append(links, secondLink)
			}
		}
	}

	// Return multiple links
	return links
}

// Gets href attribute from a token
func getHref(token html.Token, url string) string {
	for _, a := range token.Attr {
		if a.Key == "href" {
			value := a.Val

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
	}

	return ""
}

