package main

import "fmt"

type MyInterface interface {
	M()
}

type MyStructure struct {
	S string
}

// Type *MyStructure implements MyInterface
func (t *MyStructure) M() {
	fmt.Println(t.S)
}

func main() {
	t := MyStructure{"hello"}
	myFonction(t) // compilation error : type MyStructure does not implement MyInterface but the pointer does.
}

func myFonction(param MyInterface) {
	fmt.Printf("(%v, %T)\n", param, param)
}
