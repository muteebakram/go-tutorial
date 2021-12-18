package main

import (
	"log"
	"net/http"
)

var ch = make(chan string)

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

	for _, url := range urls {
		// Keyword 'go' is used to declare func as goroutine
		go func(url string) {
			ch <- contentType(url)
		}(url)
	}
	close(ch)

	for ctype := range ch {
		log.Println("Value: ", ctype)
	}
}

func main() {
	urls := []string{"https://google.com", "https://cisco.com", "https://microsoft.com", "https://xbox.com"}

	log.Print("Reading without goroutines")
	for _, url := range urls {
		contentType(url)
	}
	log.Print("Completed")

	log.Print("Reading with goroutines")
	goRoutines(urls)
	log.Print("Completed")
}
