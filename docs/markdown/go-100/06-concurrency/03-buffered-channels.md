<!-- .slide: class="with-code" -->

# Concurrenc - 29

## Buffered channels

**channels** may have a **buffer**.
You must indicate the size of the buffer as the second argument of make to initialize a channel with buffer:

```Go
ch := make(chan int, 100)
```
<!-- .element: class="big-code" -->

Sending to a buffered channel **blocks** only **when the buffer is full**.
Reception **blocks when the buffer is empty**.
