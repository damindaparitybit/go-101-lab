package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//Scale uses a pointer receiver as it has to modify the value of v.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Abs uses a value receiver. In a real case we would use a pointer receiver here also.
//it would avoid the copy of of v for each call.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	pv := &v
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	pv.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
