package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

type Config struct {
	name string
	dob  string
	city string
}

func NewConfig(name, dob, city string) (*Config, error) {
	if len(name) == 0 || len(dob) == 0 || len(city) == 0 {
		return nil, errors.New(fmt.Sprint("Invalid inputs: ", name, dob, city))
	}
	return &Config{name: name, dob: dob, city: city}, nil

}

func setUpLogger() {
	out, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func readConfig() error {
	_, err := os.OpenFile("config.json", os.O_RDONLY, 0664)
	if err != nil {
		return errors.Wrap(err, "Cannot open config file")
	}
	return nil
}

func main() {
	setUpLogger()
	log.Println("Starting app...")
	conf, err := NewConfig("a", "b", "c")
	if err != nil {
		log.Fatalln("Failed to initailize config: ", err)
	}
	log.Println("Config declared: ", conf)
	err = readConfig()
	if err != nil {
		log.Fatalln(err)
	}

}
