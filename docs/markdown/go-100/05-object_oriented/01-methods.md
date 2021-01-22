
# Object-oriented - 19

## Methods

Go does not have **classes**.

However, we can define methods on types.


##==##

<!-- .slide: class="with-code" -->


# Object-oriented - 19

## Methods

A **method** is a **function** with a **receiver**, indicated between the keyword **func** and the **name** of the method.

```Go
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
<!-- .element: class="big-code" -->

The **Abs** method has a receptor of type **Vertex** named **v**.


##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 19

## Methods

A **method** is therefore just a **function** with a **receiver**.

Here is Abs written as a regular function with no change in functionality :

```Go
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 19

## Methods

The difference is in the way the function / method is called.

Method :

```
v := Vertex{1, 3}
a := v.Abs()
```
<!-- .element: class="big-code" -->

Function :

```
v := Vertex{1, 3}
a := Abs(v)
```
<!-- .element: class="big-code" -->

##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 19

## Methods

You can also declare a method on non-struct types.

````Go
type MyFloat float64
func (f MyFloat) IsPositive() bool {
	return f >= 0
}

````
<!-- .element: class="big-code" -->


##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 19

## Methods

On the other hand, the **type of the receiver** must imperatively be declared in the same **package** as the method.

It is therefore impossible to write :

````Go
func (f float64) IsPositive() bool {
	return f >= 0
}
````
<!-- .element: class="big-code" -->

since float64 is a type defined in another package.



##==##
<!-- .slide: class="with-code" -->

# Object-oriented - 19

## Methods, pointer receiver

Methods can be declared with **pointer** receivers.

This allows the method to change the value the receiver points to :

````GO
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
````
<!-- .element: class="big-code" -->



##==##

# Object-oriented - 20

Choice of a value or pointer receiver

There are two reasons for using a pointer receiver :

- the case where the method must modify the value of the pointed receiver.

- avoid copying the value on each method call.

⇒ Usually, pointer receivers will be used.
