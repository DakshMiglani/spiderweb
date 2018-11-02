package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func validateRequest(url string) (isValid bool, validUrl string) {
	/*
	 * Goroutine to make GET Request to see if URL is available
	 */

	// Normalise URL
	url = normalisePath(url)

	// Make GET request
	resp, err := http.Get(url)

	// Validate for user error
	if err != nil {
		fmt.Printf("User error: %s\n", err.Error())
		return
	}

	// Validate for non 2xx response
	statusGroup := strconv.Itoa(resp.StatusCode)[:1]
	if statusGroup != "2" {
		fmt.Printf("Server error: %sxx returned\n", statusGroup)
		return
	}

	return true, url
}