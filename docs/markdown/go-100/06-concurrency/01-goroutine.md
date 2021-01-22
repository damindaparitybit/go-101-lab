<!-- .slide: class="with-code" -->

# Concurrency - 27

## Goroutine

A **goroutine** is a **lightweight** process managed by the "Go runtime".

```Go
go f(x, y, z)
```
<!-- .element: class="big-code" -->

starts a new **goroutine** running

```Go
f(x, y, z)
```
<!-- .element: class="big-code" -->

The evaluation of **f**, **x**, **y** and **z** is done in the current goroutine and the execution of **f** is done in the new goroutine.
