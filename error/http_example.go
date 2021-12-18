package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.cisco.com/")
	if err != nil {
		fmt.Printf("ERR: %s", err)
	} else {
		fmt.Printf("Response: %s", resp.Status)
	}
}
