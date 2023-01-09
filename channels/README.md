# Channels

## Synopsis

Channels allow for sharing data between goroutines in a seamless fashion without passing data between functions directly.

## Use

Create a new channel with the make(chan interface{})

```go
// unbuffered channel
ch := make(chan string)
// buffered channel
ch := make(chan string, 2)
```

Send value into the channel 
```go
ch <- "value"
```

Receive calue from channel
```go
// reads value from channel
fmt.Println(<-ch)
```

## Behaviors

Type types of channels:
- Buffered
- Unbuffered (default)


Unbuffered Channels (Default Behavior):
* Channel send (chan<-) & receive (<-chan) operations are blocking until both sender & receiver are ready (Synchronous communication).
  - Code execution will stop until the send or receive is successfully completed.
  - The definition of completion depends on the setup of the channel.
  - Unbuffered channels will block the goroutine when its empty and waiting to be filled.

Buffered Channels:
* Channel send (chan<-) & receive (<-chan) operations are non-blocking until the specified capacity is reached (Asynchronous Communication)
  - Up until capacity, the sender can put data onto the channel without blocking.
  - If attempt to send more than the channel capacity, the send (chan<-) operation will block until values are received (<-chan) from the channel.
  - Buffered channels will block the goroutine when:
    - Empty and waiting to be filled with data
    - At full-capacity and attempt to send (chan<-) data prior to receiver picks message off of the channel


## Range and Close

A sender can close a channel, to indicate no more messages will be sent. (Only the sender should close the channel)

Receivers can check whether a channel has been closed and that there are no more values to be received by

```go
v, ok := <-ch
if !ok {
  fmt.Println("channel is closed")
}
```

Closing channels is not a hard requirement. Only if the receiver needs to be informed that no more values are being sent.

Example, if the receiver is using the `:= range` loop


```go
go (ch chan string) {
  for i = 0; i <= 10; i++ {
    ch <- fmt.Sprintf("msg: %v", i)
  }
}(ch)

for i := range ch {
  fmt.Println("Received: ", i)
}
```

## Channel Directions

Whan passing channels to goroutines & functions, can specify to the compiler if a channel is Bi-directional or uni-directional.

Best practice to specify if a channel is send or receive only to increase type-safety

```go
func method(ch chan string) {} // Can send or receive values
func sendOnly(ch <-chan string){} // Send-Only channel
func receiveOnly(ch chan<- string){} // Receive-Only channel
```

## Select

Select lets you wait on multiple channel operations at the same time. In additon is helpful for timeouts / context timeout limits.

A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

```go
func main() {
	ch1, ch2 := make(chan string), make(chan string)

	timeout := time.After(5 * time.Second)

	go func(ch chan string) {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		ch <- "message 1"
	}(ch1)

	go func(ch chan string) {
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		ch <- "message 2"
	}(ch2)

  // infinite loop
	for {
		select {
		case val := <-ch1:
			fmt.Println("Channel 1: ", val)
		case val := <-ch2:
			fmt.Println("Channel 2: ", val)
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}
```

