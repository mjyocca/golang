// source => https://gobyexample.com/worker-pools
// example pattern showing worker pools with unbuffered channels
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	// range on channels receives values from the channel repeatedly until it is *closed*
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	// jobs & results channel are *unbuffered* channel
	jobs := make(chan int)
	results := make(chan int)

	// creating 3 workers to process work
	// initially blocked because there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// send work over the jobs channel
	// this will unblock the workers to start processing the jobs, since both sender and receiver are ready
	/*
		since we're using unbuffered channels, we will need to send without blocking, so we wrap in an anonymous go routine function.
		otherwise will deadlock
	*/
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		// done sending jobs to channel
		// since we're sending, we can close
		// closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop
		close(jobs)
	}()

	for a := 1; a <= numJobs; a++ {
		res := <-results
		fmt.Println("main received <- ", res)
	}
}
