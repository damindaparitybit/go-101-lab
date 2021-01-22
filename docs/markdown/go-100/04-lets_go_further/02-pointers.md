<!-- .slide: class="with-code" -->

# Let's Go further - 13

## Pointers

Go has pointers. A pointer contains the memory address of a value.

The type **`*T`** is a **pointer** to a value of type **T**.

`var p *int`

Here, p is a pointer to an integer.


##==##
<!-- .slide: class="with-code" -->

# Let's Go further - 13

## Pointers

Zero value of a pointer is **nil**.

The **&** operator generates a pointer to its operand.

```Go
var p *int  // p is nil here
i := 42
p = &i  // p is a pointer to i
```
<!-- .element: class="big-code" -->

Notes:
- ampersand


##==##
<!-- .slide: class="with-code" -->

# Let's Go further - 13

## Pointers

The **\*** operator denotes the underlying value of the pointer.

```Go
i := 3
p := &i
fmt.Println(*p)   // here i equals 3 through the pointer p
*p = 21           // set i through the pointer p
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Let's Go further - 13

## Pointers to structures

The struct fields can be accessed via a struct pointer.
Indirection via the pointer is transparent.

```Go
func main() {
	v := Vector{1, 2}
	p := &v
	p.X = 7
	fmt.Println(v)  // { X:7, Y:2 }
}
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide:-->

# Let's Go further - 13 - exercise

## Pointers and structures

Answer the questions by replacing REPLACEME tag with true or false.


##==##
<!-- .slide: class="first-slide" sfeir-level="1" sfeir-techno="Go" -->

If you enjoy the lab, send a tweet!

#golang

@Sfeirlille @sfeir

@sebastienfriess







