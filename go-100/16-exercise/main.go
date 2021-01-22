package main

import "fmt"

//indexOf takes a search word and an array of strings where it may be contained,
//and returns the index of the word in the slice or -1 if it is not found
func indexOf(search string, s []string) int {
	return -1
}

func main() {
	var days = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	var index int

	index = indexOf("monday", days)
	fmt.Printf("Monday is the %dst day of week\n", index+1)

	index = indexOf("thursday", days)
	fmt.Printf("Thursday is the %drth day of week\n", index+1)

	index = indexOf("plop", days)
	fmt.Printf("'Plop' is a day of week ? : %v\n", index > -1)
}
