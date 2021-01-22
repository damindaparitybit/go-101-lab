package main

import "fmt"

var a bool = false
var b = false
var c bool

func uneFonction() {
	var w bool = false
	var x = false
	var y bool
	z := false

	fmt.Printf("w = %v and its type is %[1]T\n", w)
	fmt.Printf("x = %v and its type is %[1]T\n", x)
	fmt.Printf("y = %v and its type is %[1]T\n", y)
	fmt.Printf("z = %v and its type is %[1]T\n", z)
}


func fibonacci(iter int) int {
	n, prev := 1, 0
	for i := 0; i < iter; i++ {
		n, prev = n+prev, n
	}
	return n
}

func main() {
	fmt.Printf("a = %v and its type is %[1]T\n", a)
	fmt.Printf("b = %v and its type is %[1]T\n", b)
	fmt.Printf("c = %v and its type is %[1]T\n", c)
	uneFonction()

	age, name, alive := 30, "John", true
	fmt.Println(age, name, alive)

	fmt.Println(fibonacci(20))
}
