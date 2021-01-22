package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Is it weekend soon ?")
	today := time.Now().Weekday()
	switch today {
	case time.Saturday:
		fallthrough
	case time.Sunday:
		fmt.Println("It's weekend 😁 !")
	case time.Friday:
		fmt.Println("Soon but not yet 😊")
	default:
		fmt.Println("Too long before weekend 😢")
	}
}
