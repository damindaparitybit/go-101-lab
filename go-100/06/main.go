package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	x      float64
	str    string
)

func main() {
	const f = "%T(%[1]v)\n"
	fmt.Printf(f, ToBe)
	fmt.Printf(f, MaxInt)
	fmt.Printf(f, z)
	fmt.Printf(f, x)
	fmt.Printf(f, str)

	// Type conversion
	i := 42
	f64 := float64(i)
	fmt.Printf(f, i)
	fmt.Printf(f, f64)
}
