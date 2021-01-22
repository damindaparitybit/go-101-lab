package main

import "fmt"


func inverse(x float64) (float64, bool) {
	if x == 0 {
		return 0, false
	}
	return 1 / x, true
}

func main() {

	if inv, ok := inverse(3); ok {
		fmt.Printf("inverseof %f is %f\n", 3.0, inv)
	} else {
		fmt.Printf("impossible to compute the inverse of %f", 3.0)
	}


	if inv, ok := inverse(0); ok {
		fmt.Printf("inverseof %f is %f\n", 0.0, inv)
	} else {
		fmt.Printf("impossible to compute the inverse of %f", 0.0)
	}
}
