package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("I talk whenever I want !")
	say("Stop interrupting me !")
}
