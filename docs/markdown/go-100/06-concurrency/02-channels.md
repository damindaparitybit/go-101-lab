<!-- .slide: class="with-code" -->

# Concurrency - 28

## Channels

**channels** are typed pipes through which you can send and receive values with the channel operator, **<-**

```Go
ch <- v    // send v on channel ch.
v := <- ch // receive from ch and assign the value to v.
```
<!-- .element: class="big-code" -->

(The data flow follows the arrow.)

##==##

# Concurrency - 28

## Channels

**channels** have to be instantiated before use :

```Go
    ch := make(chan int)
```
<!-- .element: class="big-code" -->

By default, sending and receiving block until the other side is ready. This allows goroutines to be synchronized without explicit locks or condition variables.
