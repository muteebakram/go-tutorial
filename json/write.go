package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

type Book struct {
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
	Author Author    `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	a := Author{Name: "Muteeb", Age: 20}
	b := Book{Name: "How to live?", Date: time.Now(), Author: a}
	log.Printf("Go struct: %+v", b)

	data, err := json.MarshalIndent(b, "", "     ")
	if err != nil {
		log.Fatal("Marshal failed: ", err)
	}
	log.Printf("byte data: %+v", data)

	log.Printf("json data: %+v", string(data))
	_ = ioutil.WriteFile("books.json", data, 0666)
}
