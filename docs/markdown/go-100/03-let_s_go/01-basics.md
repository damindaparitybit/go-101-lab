<!-- .slide: class="with-code" -->

# Basics - 01

## Hello, 世界

```Go
package main

import "fmt" // Programs are organised into packages
             // It has at least a main package, containing a main function

func main() {
   fmt.Println("Hello, 世界")
}
```
<!-- .element: class="big-code" -->

##==##

# Basics - 01

## Hello, 世界

Note : as expected, 世界 means “world”, but these characters are not available in ANSI.

- By convention in Go, all source files are **UTF8** encoded.
- All string are **UTF8** encoded by default.
- It is possible to name a variable **Δt** for exemple.


##==##
<!-- .slide: class="with-code" -->

# Basics - 01

## Compilation, execution

Go to 01 folder, then :
- Compile the program : `$ go build main.go`
- Then run it : `$ ./main`
- Or more directly : `$ go run main.go`

Go is a compiled language. The compiler produces an executable self-contained binary.
To compile Go for another OS :
` $ GOOS=linux GOARCH=amd64 go build main.go`

##==##

# Basics - 02

## Imports

- With Go, it is **forbidden** to import a package and not using it !
- To import several packages, we well prefer the factored import statement
- By convention, the package name is the same than the last path element. Ex: “math/rand” ⇒ package “rand”


##==##

# Basics - 03

## Export an identifier

For an identifier (variable, function, type, etc.) to be visible outside its declared package it has to start with an uppercase.
There is no “public” nor “private” keyword in Go

**⇒ an identifier is exported (ie. public) if and only if it starts with an uppercase.**


##==##
<!-- .slide: class="with-code" -->

# Basics - 04

## Functions

Fonction is declared with **func** keyword followed by function name.

The function body is declared between **{** and **}**.

The **{** character has to be on the same line as the **func** keyword.

`
    func myFonction() { }
`

Notes:
- {} braces
- () brackets
- [] square brackets

##==##
<!-- .slide: class="with-code" -->

# Basics - 04

## Functions

Function has zero or several parameters.

`func myFonction(n int, s string) { }`

Note that that type comes **after** the parameter name.


##==##
<!-- .slide: class="with-code" -->

# Basics - 04

## Functions

A fonction may return a single value.

```Go
func myFonction(n int, s string) string {
    return s + strconv.Itoa(n)
}
```
<!-- .element: class="big-code" -->



##==##
<!-- .slide: class="with-code" -->

# Basics - 04

## Functions

With Go, a function may also return several **values** values.

```Go
func maFonction(n int, s string) (string, bool) {
	return s + strconv.Itoa(n), n == 0
}
```
<!-- .element: class="big-code" -->



##==##
<!-- .slide: class="with-code" -->

# Basics - 04

## Functions

When several parameters share the same type, you can declare it once at the end :

```Go
func myFonction(a bool, b int, c int, d int, e string, f string) { ... }
⇔
func myFonction(a bool, b, c, d int, e, f string) { ... }
```
<!-- .element: class="big-code" -->

##==##
<!-- .slide: class="with-code" -->

# Basics - 05

## Variables

One or several variables are declared using **var** instruction followed by its name, then its type and finally its initial value.

`var name type = expression`

Type may be ignored if an initial value is set.


##==##
<!-- .slide: class="with-code" -->

# Basics - 05

## Variables

Here are three ways to declare the same variable :

```Go
var name string = ""    // unnecessarily complete

var name = ""           // type is inferred by initial value

var name string         // zero value of string is empty
```
<!-- .element: class="big-code" -->

##==##
<!-- .slide: class="with-code" -->

# Basics - 05

## Variables

It is also possible to factorise variable declaration with brackets :

```Go
var (
    toto = ""
    titi float32
    tata = true
)
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Basics - 05

## Variables

Several variables can be declared on a single line :

```Go
var a, b, c = "Hello", 42, true
```
<!-- .element: class="big-code" -->

or if they share same type :
```Go
var x, y, z float64
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Basics - 05

## Variables

Inside a function, the **short** declaration can be used :

```Go
func plop() {
	a := "plop!" // same as : var a = "plop!"
	fmt.Println(a)
}
```
<!-- .element: class="big-code" -->


##==##

# Basics - 05

## Variables

Warning, keep in mind that <span style="color:blue">:=</span> is a <span style="color:blue">declaration+assignation</span>, whereas <span style="color:green">=</span> is an <span style="color:green">assignation</span>

This is why, it is <span style="color:red">impossible</span> to assign a vale to an <span style="color:red">already declared</span> variable using <span style="color:red">:=</span>


##==##

# Basics - 06

## Types

- bool
- string
- int, int8, int16, int32, int64
- uint, uint8, uint16, uint32, uint64

**int** is either an **int32** or an **int64**, depending on the platform and the compiler.

It is the same for **uint** that is either a **uint32** or **uint64**.


##==##

# Basics - 06

## Types

- **byte** is an alias for int8
- **rune** is an alias for int32 and represents a Unicode “code point”
- **uintptr** is an unsigned integer, large enough to contain the value of a pointer (generally 32 or 64 bits).


##==##

# Basics - 06

## Types

- **float32, float64**
- **complex64, complex128**

Type **complex64** is a complex number where the real part and the imaginary part are both **float32**. Type **complexe128** uses **float64**.


##==##

# Basics - 06

## Zero values

Each type has a **zero value**. A variable declared without its initiale value (ex: var name type) has the **zero value** of its type.
Zero values of basics types are the following :
- 0 pour numeric types,
- false for booleans,
- "" (empty string) for strings.

##==##
<!-- .slide: class="with-code" -->
# Basics - 06

## Type conversion

With Go, every type conversion must be **explicit**.

The expression T(v) converts value v in the type T.

````Go
var i int = 42
var f float64 = float64(i)
````
<!-- .element: class="big-code" -->
