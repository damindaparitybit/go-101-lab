package main

import (
	"fmt"
	"time"
)

func say(s string, ch chan string) {
	fmt.Println("ch <- s waiting for someone to read the channel ch.")
	ch <- s
	fmt.Println("s  has been sent to the channel.")
}

func main() {
	ch := make(chan string)
	go say("Hello", ch)

	time.Sleep(10 * time.Millisecond)
	fmt.Println("s := <-ch waiting for some data on ch.")

	s := <-ch

	time.Sleep(10 * time.Millisecond)
	fmt.Printf("String '%s' has been read on channel ch.\n", s)
}
