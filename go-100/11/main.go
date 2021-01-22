package main

import (
	"errors"
	"fmt"
)

//Distance in meters
type Distance float64

//Time in seconds
type Time float64

//Speed in m/s
type Speed float64

func computeSpeed(d Distance, t Time) (Speed, error) {
	if t == 0 {
		return 0, errors.New("impossible to compute a speed with a null time")
	}
	return Speed(float64(d) / float64(t)), nil
}

func main() {
	// _ allows to ignore the second value, here the error
	speed, _ := computeSpeed(130000, 3600)

	fmt.Printf("You are going at %.2f m/s\n", speed)
}
