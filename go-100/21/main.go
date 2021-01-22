package main

type Car interface {
	MoveTo(x, y, z float64)
	Stop()
	IsMoving() bool
}
