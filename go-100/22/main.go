package main

import "fmt"

type Eagle struct{}

func (e Eagle) Eat() {
	fmt.Println("Eagle eats.")
}
func (e Eagle) Fly() {
	fmt.Println("Eagle flies.")
}

// A380 is a ✈️
type A380 struct{}

func (a A380) Fly() {
	fmt.Println("A380 flies.")
}

// Bird is an interface defining a bird's behaviour.
type Bird interface {
	Eat()
	Fly()
}

// Flyer is an interface defining the behaviour of something that is flying.
type Flyer interface {
	Fly()
}

func main() {
	// Eagle implements Bird interface
	var Bird Bird = Eagle{}

	// A380 implements  Flyer interface
	var plane Flyer = A380{}

	Bird.Fly()
	Bird.Eat()
	plane.Fly()

	// Eagle implements also Flyer interface
	plane = Eagle{}
	plane.Fly()
}
