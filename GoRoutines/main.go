package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, pages chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	pages <- Page{URL: url, Size: len(body)}
}

func main() {
	pages := make(chan Page)
	urls := []string{"http://example.com", "http://golang.org"}

	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("Size of the page retrieved from %s: %d bytes\n", page.URL, page.Size)
	}
}
