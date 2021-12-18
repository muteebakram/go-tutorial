package main

import (
	"log"
	"net/http"
	"sync"
)

func contentType(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Err:", err)
	}

	defer res.Body.Close()
	ctype := res.Header.Get("content-type")
	log.Printf("URL: %s, Content-Type: %s", url, ctype)
	return ctype
}

func goRoutines(urls []string) {
	// sync is used to wait unill goroutine completed
	var wg sync.WaitGroup
	for _, url := range urls {
		// Adding goroutine to wait group
		wg.Add(1)

		// Keyword 'go' is used to declare func as goroutine
		go func(url string) {
			contentType(url)

			// Mark goroutine as completed
			wg.Done()
		}(url)
	}

	// Wait until all goroutines are completed
	wg.Wait()
}

func main() {
	log.Print("Reading without goroutines")
	urls := []string{"https://google.com", "https://cisco.com", "https://microsoft.com", "https://xbox.com"}

	for _, url := range urls {
		contentType(url)
	}
	log.Print("Completed")

	log.Print("Reading with goroutines")
	goRoutines(urls)
	log.Print("Completed")
}
