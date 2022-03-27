package main

import (
	"fmt"
	"net/http"
	"time"
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

	}

	for u := range c {
		go func(u string) {
			time.Sleep(3 * time.Second)
			checkSite(u, c)
		}(u)
	}

}

func checkSite(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- url
		return
	}
	fmt.Println(url, "is up!")
	c <- url
	return
}
