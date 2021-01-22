package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("The double of %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q has a length of %v octets\n", v, len(v))
	default:
		fmt.Printf("I don't know type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
