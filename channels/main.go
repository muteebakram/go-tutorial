package main

import (
	"log"
	"time"
)

func main() {
	// Creating channel
	ch := make(chan int)

	// Nothing in channel but reading it will throw error
	// fatal error: all goroutines are asleep - deadlock!
	// <-ch
	log.Println("Here")

	go func() {
		// Push number 3 into int channel
		ch <- 3
	}()

	// Get the value from channel
	val := <-ch
	log.Println("Value: ", val)

	log.Println("Push multiple values.")
	go func() {
		for i := 0; i < 3; i++ {
			log.Println("Sending: ", i)
			ch <- i
			time.Sleep(1 * time.Second)
		}
		// close channel telling sending done
		close(ch)
	}()
	log.Printf("Channel: %v\n", ch)
	log.Println("Poping multiple values.")
	for val := range ch {
		log.Println("Receiving: ", val)
		time.Sleep(1 * time.Second)
	}
}
