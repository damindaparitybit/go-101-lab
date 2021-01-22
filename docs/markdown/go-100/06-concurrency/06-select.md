<!-- .slide: class="with-code" -->

# Concurrency - 32

## Select

The **select** statement allows a goroutine to wait on multiple channels simultaneously.

```Go
select {
    case c <- x:       // waits to send x on 'c'
        x, y = y, x+y
    case <- quit:      // wait to receive on 'quit'
        fmt.Println("quit")
}
```
<!-- .element: class="big-code" -->

##==##

<!-- .slide: class="with-code" -->

# Concurrency - 32

## Select

A **select** blocks until one of its cases can run, then it runs. He chooses a random one if more than one is ready.

##==##

<!-- .slide: class="with-code" -->

# Concurrency - 33

## Select

The **default** case in a **select** is executed if no other case is ready and therefore avoids blocking the select.

```Go
select {
    case i := <-c:
 // use i
    default:
 // executed if nothing can be read from 'c'
}
```
<!-- .element: class="big-code" -->
