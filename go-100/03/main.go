package main

import (
	"fmt"
	"math"
)

var iAmInvisibleOutsideOfMyPackage int

var IAmVisibleOutsideOfMyPackage int

func main() {
	fmt.Println(math.pi) //cannot refer to unexported name math.pi
}
