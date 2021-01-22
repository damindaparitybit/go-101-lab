<!-- .slide: class="with-code" -->

# Let's Go further - 11

## Type declaration

We can declare a new type using the **type** keyword.

For example, it is possible to declare a type whose underlying type is a number:

`type Distance int64`

This is exactly what is done in the Go time package :

`type Duration int64`


##==##
<!-- .slide: class="with-code" -->

# Let's Go further - 11

## Type declaration

It is possible to perform the same operations as on the underlying type, but only between variables of the same type.
This is very useful for avoiding homogeneity errors in a formula.

```Go
var d Distance = 300
var t Duration = 60
v := d / t // compilation error
v := float64(d) / float64(t)  // ok
```
<!-- .element: class="big-code" -->

Notes:
It is possible to switch back to the underlying type for example int64 (d)


##==##
<!-- .slide: class="with-code" -->

# Let's Go further - 12

## Structures

A structure is a collection of fields. To declare a structure is to declare a new type.
So it makes sense to declare a structure using the **type** keyword and the underlying type struct:

```Go
type Vector struct {
	X int
	Y int
}
```
<!-- .element: class="big-code" -->

##==##
<!-- .slide: class="with-code" -->

# Let's Go further - 12

## Structures

A structure can be literally instantiated, offering the possibility of initializing the value of the fields:

```Go
v1 := Vector{1, 2}  // X:1 and Y:2
v2 := Vector{X: 1}  // X:1, and Y:0 is implicite
v3 := Vector{}      // X:0 and Y:0
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Let's Go further - 12

## Structures

The fields of the structure are accessible using a dot :

```Go
func main() {
	v := Vector{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```
<!-- .element: class="big-code" -->
