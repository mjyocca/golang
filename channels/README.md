# Channels

## Synopsis

Channels allow for sharing data between goroutines in a seamless fashion without passing data between functions directly.

## In Practice

Create a new channel with the make(chan interface{})

```go
ch := make(chan string)
```

Send value into the channel with 
```go
ch <- "value"
```

Receive calue from channel with
```go
// reads value from channel
fmt.Println(<-ch)
```

## Behaviors

Type types of channels:
- Buffered
- Unbuffered (default)


Default Behavior:
* Channel send (chan<-) & receive (<-chan) operations are Blocking until both sender & receiver are ready.
  - Code execution will stop until the send or receive is successfully completed.
  - The definition of completion depends on the setup of the channel.

