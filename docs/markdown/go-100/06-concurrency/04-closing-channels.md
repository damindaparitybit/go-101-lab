<!-- .slide: class="with-code" -->

# Concurrency - 30

## Closing channels

A sender can close a channel to indicate that no more value will be sent :

```Go
close(ch)
```
<!-- .element: class="big-code" -->

Receivers can check if a channel has been closed :

```Go
v, ok := <-ch
```
<!-- .element: class="big-code" -->

**ok** is false if the channel is closed and there is no more value to receive.

##==##

# Concurrency - 30

## Closing channels

Warning :

- **Only the sender should close a channel**, never the receiver.
  Sending on a closed channel causes panic.
- Channels are not like files, you usually don't **need to close them**.
  Closing is only necessary when the receiver needs to be informed that there are no more values to come, such as ending a range loop.
