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
	file, _ := ioutil.ReadFile("books.json")
	data := Book{}
	_ = json.Unmarshal([]byte(file), &data)
	log.Printf("json data: %+v", data)
}
