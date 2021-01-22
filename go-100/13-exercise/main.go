package main

import "fmt"

const REPLACEME = "replaceme"

// Structure
type Age int

type Person struct {
	name   string
	age    Age
	friend *Person
}

func main() {

	john := Person{"John", 25, nil}
	pJohn := &john

	// in Go, pointer indirection through the pointeur is transparent.
	// this means that this line of code is equivalent to (*pJohn).age++
	pJohn.age++

	// replace REPLACEME with true or false
	assert("Q11", john.age == 25, REPLACEME)
	assert("Q12", pJohn.age == 26, REPLACEME)

	pAgeJ := &(john.age)
	bob := Person{"Bob", 26, &john}
	pAgeB := &(bob.age)

	assert("Q2", pAgeJ == pAgeB, REPLACEME)

	assert("Q3", *pAgeJ == bob.age, REPLACEME)

	*pAgeJ = 10

	assert("Q4", *pAgeJ == john.age, REPLACEME)

	john.age = 12

	assert("Q5", *pAgeJ == john.age, REPLACEME)

	john.age = *pAgeB

	assert("Q6", *pAgeJ == john.age, REPLACEME)

	*pAgeB = 18

	assert("Q7", bob.age == 18, REPLACEME)

	bob.friend.age = *pAgeB + 1

	assert("Q8", john.age == 19, REPLACEME)

	bob.friend = &bob
	bob.friend.age = 20

	assert("Q9", bob.age == 20, REPLACEME)

	bob.friend = &bob
	bob.friend.friend = &john

	assert("Q10", bob.friend == pJohn, REPLACEME)
}

func assert(message string, value bool, expected interface{}) {
	ve, ok := expected.(bool)
	if !ok {
		fmt.Printf("%5s : %s\n", message, "Missing answer")
		return
	}

	if value != ve {
		fmt.Printf("%5s : %s\n", message, "Wrong answer")
	} else {
		fmt.Printf("%5s : %s\n", message, "Correct answer")
	}
}
