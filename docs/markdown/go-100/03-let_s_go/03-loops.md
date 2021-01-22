
# Basics - 09

## Loops

- Go has only one loop : the **for** loop.
- It has three components separated by semicolons:
  - The **initialization** declaration: executed before the first iteration
  - The **condition** expression: evaluated before each iteration
  - The **post** declaration: executed at the end of each iteration


##==##
<!-- .slide: class="with-code"-->

# Basics - 09

## Loops

As for **if** instruction, there are no brackets surrounding the three components of the **for** declaration and the braces **{}** are also mandatory.

```Go
for i := 0; i < 10; i++ {
	sum += i
}
```
<!-- .element: class="big-code" -->


##==##

# Basics - 09

## Loops

Often, the initialization declaration will be a short variable declaration and variables so declared are only visible as part of the **for** declaration.

The iteration of the loop will be terminated once the boolean condition evaluates to false.


##==##
<!-- .slide: class="with-code" -->

# Basics - 09

## Loops

The initialization and completion declaration are optional.

`
for ; n < 1000; {
	n += n
}
`

can be written without ";"

`
for  n < 1000; {
	n += n
}
`

Notes:
Equivalent to while loop



##==##
<!-- .slide: class="with-code" -->

# Basics - 09

## Loops

The condition can also be omitted.
We thus obtain an infinite loop :

`for { }`



##==##
<!-- .slide: class="with-code" -->

# Basics - 10

## Switch

The **switch** does exactly what is expected of it. As with the "if" or "for", an expression can be added before the condition.

```Go
switch initialisation; expression {
    case a:
    case b:
    default:
}
```
<!-- .element: class="big-code" -->


##==##

# Basics - 10

## Switch

In Go, a **case** ends automatically (the **break** instruction is not necessary).
If, on the contrary, we want the following **case** to be executed, you must add the keyword **fallthrough**.


##==##
<!-- .slide: class="with-code" -->

# Basics - 10

## Switch

If the condition is omitted, the **switch** can elegantly replace long if-then-else strings.

```Go
switch { // equivalent to switch true
    case i < 0:
    case i == 0:
    default:
}
```
<!-- .element: class="big-code" -->
