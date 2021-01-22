<!-- .slide: class="with-code" -->

# Concurrency - 31

## For range on channels

The lop for `i := range c` receives values from the channel until it is closed.

```Go
for i := range c {
 fmt.Println(i)
}
```
<!-- .element: class="big-code" -->
