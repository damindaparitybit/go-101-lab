package main

import "fmt"

var days = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

func main() {
	for i, v := range days {
		fmt.Printf("days[%d] = %s\n", i, v)
	}

	powerOf2 := make([]int, 10)
	for i := range powerOf2 {
		powerOf2[i] = 1 << uint(i)
	}
	for _, value := range powerOf2 {
		fmt.Printf("%d\n", value)
	}
}
