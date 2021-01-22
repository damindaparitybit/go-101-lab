
# Object-oriented - 21

## Interfaces

In Go, there is no notion of inheritance between **struct**, but the notion of **interface** is present.


##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 21

## Interfaces

As expected, an interface type is defined as a set of method signatures.

````Go
type Bird interface {
	Fly(direction Vector)
	Eat() (int, bool)
}
````
<!-- .element: class="big-code" -->


##==##

# Object-oriented - 22

## Interfaces

Specificity Go: Interfaces are implemented **implicitly**.

A type “automatically” implements an **interface** by implementing its methods.

⇒ There is no explicit “implements” keyword.



##==##
<!-- .slide: class="with-code" -->
# Object-oriented - 22

## Interfaces

If we have :

````Go
type I interface {
	M()
}
type T struct { }
````
<!-- .element: class="big-code" -->


Then you just need to delcare the method:

````Go
func (t T) M() {
	/*...*/
}
````
<!-- .element: class="big-code" -->

for type T to  implement I.


##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 22

## Interfaces

If we have :
```Go
type Eagle struct {}
func (a Eagle) Fly() { /*...*/ }
func (a Eagle) Eat() { /*...*/ }
```
<!-- .element: class="big-code" -->

and :
```Go
type Bird interface {
	Fly()
	Eat()
}
type Flyer interface { Fly() }
```
<!-- .element: class="big-code" -->


⇒ What interface (s) does Eagle implement?


##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 22

## Interfaces

If we have :

```Go
type Eagle struct {}
func (a Eagle) Eat() { /*...*/ }
func (a Eagle) Fly() { /*...*/ }
```
<!-- .element: class="big-code" -->

and :

```Go
type Bird interface {
	Eat()
	Fly()
}
type Flyer interface { Fly() }
```
<!-- .element: class="big-code" -->

⇒ **Eagle** implements both **Bird** and **Flyer** interfaces.


##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 22

## Interfaces

However, if we have:

```Go
type A380 struct {}
func (a A380) Fly() { /*...*/ }
```
<!-- .element: class="big-code" -->


and :

```Go
type Bird interface {
	Eat()
	Fly()
}
type Flyer interface { Fly() }
```
<!-- .element: class="big-code" -->


⇒ What interface (s) does A380 implement ?



##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 22

## Interfaces

On the other hand, if we have:

```Go
type A380 struct {}
func (a A380) Fly() { /*...*/ }
```
<!-- .element: class="big-code" -->

and :

```Go
type Bird interface {
	Eat()
	Fly()
}
type Flyer interface { Fly() }
```
<!-- .element: class="big-code" -->

⇒ **A380** only implements the **Flyer** interface.



##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 23

## Interfaces

Warning,

You just need to declare the method:

```Go
func (t T) M() {
	fmt.Println(t.S)
}
```
<!-- .element: class="big-code" -->

so that the **T** type implements **I**.


but if you declare :

```Go
func (t *T) M() {
	fmt.Println(t.S)
}
```
<!-- .element: class="big-code" -->

it is the ***T** type that implements **I**.


##==##

# Object-oriented - 23

## Interfaces

⇒ In general, all methods on a given type **T** must have either a value receiver (T) or a pointer receiver (* T), but not a mix of the two.

Otherwise, **T** would implement some interfaces while ***T** would implement others, which would be a nightmare to use.



##==##

# Object-oriented - 24

## Empty interface

The empty interface does not define, as its name suggests, any method:

**interface{}**

All types implement the empty interface since a type always has at least zero methods. We therefore use the empty interface to handle values of unknown type.

Ex: *fmt.Print* takes parameters of type **interface{}**



##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 25

## Type assertion

A type assertion gives access to the concrete **(underlying)** value of an interface value.

`t := i.(T)`

This statement asserts that the interface value **i** contains the concrete type **T** and assigns the underlying value to the variable **t**.



##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 25

## Type assertion

If **i** does not hold **T**, the declaration will trigger a **panic**.

But the type assertion can also be done like this :

`t, ok := i.(T)`

If i does not hold **T**, no panic will be triggered, but **ok** will be **false**.



##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 26

## Type switch

A type switch is used to manage the different concrete types that an interface could have :

```Go
switch v := i.(type) {
case T:    // here v is of type T
case S:    // here v is of type S
default:   // here v is of same type as i
}
```
<!-- .element: class="big-code" -->
