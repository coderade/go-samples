package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://amazon.com",
		"https://golang.org",
	}

	for _, url := range urls {
		checkSite(url)
	}
}

func checkSite(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
	}
	fmt.Println(url, "is up!")
}
