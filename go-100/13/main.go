package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // points to i
	fmt.Println(*p) // reads i through the pointer
	*p = 21         // modifies i through the pointer
	fmt.Println(i)  // displays the new value of i

	p = &j         // points to j
	*p = *p / 37   // divides j through the pointer
	fmt.Println(j) // displays the news value of j
}
