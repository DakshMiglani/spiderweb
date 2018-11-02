package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Formatter - this could be a bit nicer
func formatToString(links []string) string {
	str := ""
	for _, link := range links {
		str = str + link + "\n"
	}
	return str
}

// Output
func saveToFile(content string) {
	err := ioutil.WriteFile("links.txt", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

func saveAsXml(links []string) {
	str := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n"
	for _, link := range links {
		str = str + "<url><loc>" + link + "</loc></url>\n"
	}
	content := str + "</urlset>"
	err := ioutil.WriteFile("sitemap.xml", []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}

// Path normalisation
func normalisePath(url string) string {
	// Strip URL of protocol identifier
	if strings.Contains(url, "https://") {
		url = url[8:]
	} else if strings.Contains(url, "http://") {
		url = url[7:]
	}

	// Ensure there's no trailing slash
	if url[len(url)-1:] == "/" {
		url = url[:len(url)-1]
	}

	return "http://" + url
}

// Contains method for a string slice bc this isn't native to golang (??)
func stringInSlice(slice []string, str string) bool {
	for _, sliceItem := range slice {
		if sliceItem == str {
			return true
		}
	}
	return false
}

// Gets user input without CRLF chars which break everything bc Windows
func userInput(req string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(req)
	stdin, _ := reader.ReadString('\n') // Reads line from console
	stdin = strings.TrimSuffix(stdin, "\r\n") // Removes CRLF if present
	stdin = strings.TrimSuffix(stdin, "\r") // Removes CR if present
	stdin = strings.TrimSuffix(stdin, "\n") // Removes LF if present
	return stdin
}

// Gets user input as int
func intUserInput(req string) *int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(req)
	stdin, _ := reader.ReadString('\n') // Reads line from console
	stdin = strings.TrimSuffix(stdin, "\r\n") // Removes CRLF if present
	stdin = strings.TrimSuffix(stdin, "\r") // Removes CR if present
	stdin = strings.TrimSuffix(stdin, "\n") // Removes LF if present

	if 1 > len(stdin) {
		return nil
	}

	i, err := strconv.Atoi(stdin)
	if err != nil {
		fmt.Println("Invalid number. Please try again.")
		os.Exit(2)
	}
	return &i
}