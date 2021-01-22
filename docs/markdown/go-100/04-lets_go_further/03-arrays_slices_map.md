<!-- .slide: class="with-code" -->

# Let's Go further - 14

## Arrays

The type **[n]T** is an array of n values of type **T**.

The expression :

`var a [10]int`

declares a variable **a** as an array of ten integers.

##==##

# Let's Go further - 14

## Arrays

The length of an array is part of its type.

⇒ tables cannot be resized.

##==##

# Let's Go further - 15

## Slices

A **slice** points to an array of values and has a length.

**[]T** is a slice of elements of type **T**.

**len(s)** returns the length of slice **s**.

<img class="float-right" src="./assets/go-100/images/slice-1-array.png">

**cap(s)** returns the capacity of slice **s**.

Zero value of a slice is **nil**.

##==##

<!-- .slide: class="with-code" -->

# Let's Go further - 15

## Slices

Slices can be created with the **make** function. This allocates an empty array and returns a slice that references that array :

`a := make([]int, 5) // len(a)=5`

##==##

# Let's Go further - 15

## Slices

The slices can be resized, creating a new slice that points to the same array.

The expression :

**s[low:high]**

gets the slice elements from **low** to **high-1** inclusive (or **high** excluded).

##==##

<!-- .slide: class="two-column-layout with-code" -->

# Let's Go further - 15

## Slices

##--##

<br><br>

<img src="./assets/go-100/images/slice-2-array.png">

<br>

<img src="./assets/go-100/images/slice-2-array-decoupe.png">

##--##

<!-- .slide: class="with-code" -->

<br><br>

```Go
var a = make([]byte, 5)
len(a) == 5
cap(a) == 5
```

<!-- .element: class="big-code" -->

<br>

```Go
var b = a[2:4]
len(b) == 2
cap(b) == 3
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Let's Go further - 15

## Slices

To specify a capacity, pass a third argument to make:

```Go
b := make([]int, 0, 5)      // len(b)=0, cap(b)=5
b = b[:cap(b)]              // len(b)=5, cap(b)=5
b = b[1:]                   // len(b)=4, cap(b)=4
```

<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Let's Go further - 15

## Slices

It's common to add new items to a slice, and so Go offers a built-in **append** function.

`func append(s []T, vs ...T) []T`

The first parameter s of append is a slice of type T, and the following are T values to add to the slice.

Notes:
The resulting value of append is a slice containing all the items from the original slice and the supplied values.

If the supporting array of s is too small to hold all of the given values a larger array will be allocated. The slice returned will point to the newly allocated array.

##==##

# Let's Go further - 16

## For … range

The **range** form of the **for** loop allows to iterate over all the elements of a slice, an array, or a map.

Two values are returned for each iteration:
- the index
- a **copy** of the element at this index.

##==##

<!-- .slide: class="with-code" -->

# Let's Go further - 16

## For … range

```Go
for i, v := range s { }
```
<!-- .element: class="big-code" -->

is equivalent to :

```Go
for i := 0; i < len(s); i++ {
    v := s[i] // s[i] is editable here
}
```

<!-- .element: class="big-code" -->


##==##

# Let's Go further - 16

## For … range

You can omit the index by replacing it with “\_” :

`for _, v := range s { }`

You can also omit the value.

`for i := range s { }`

##==##

# Let's Go further - 16 - exercise

## Slices and range

Write the **indexOf** function.

##==##

<!-- .slide: class="with-code" -->

# Let's Go further - 17

## Maps

A map is a list of key values. A map is declared :

`map[typeKey]typeValue`

Maps must be created with make. Their zero value is nil.

`m = make(map[string]Vector)`

##==##

<!-- .slide: class="with-code" -->

# Let's Go further - 17

## Maps

Like structs, a map can be instantiated literally. However, you must indicate the value of each key :

```Go
var m = map[string]Vector{
	"Bell Labs": Vector{40.68433, -74.39967},
	"Google": Vector{37.42202, -122.08408},
}
```
<!-- .element: class="big-code" -->

##==##

# Let's Go further - 17

## Maps

- Assign : `m[key] = elem`
- Delete : `delete(m, key)`
- Get :

  - `elem = m[key]`     // elem takes zero value of its type if key is missing.
  - `elem, ok = m[key]` // ok is true if key is existing, false if missing.

##==##

# Let's Go further - 18

## A bit further with functions

In Go, functions are also values. They can be used as function parameters like any other value.

Function values can be used as function arguments or return values.

We say that the functions are **first class citizen** in Go.
