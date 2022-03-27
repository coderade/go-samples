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
	c := make(chan string)

	for _, url := range urls {
		go checkSite(url, c)
		fmt.Println(<-c)
	}

}

func checkSite(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		c <- url + " might be down"

	}
	c <- url
}
