package main

import (
	"fmt"
	"time"
)

func main() {
	problem1()
	fmt.Println("A line of dashes will appear in 2 seconds")
	Sleep(2)
	fmt.Println("-----------------------------------------------")
	problem3()
}

func problem1() {
	fmt.Println("The direction of the arrow specifies the channel type. For example,")
	fmt.Println("func ping(c chan <- string) // only allows SENDING to channel c")
	fmt.Println("func ping(c <- chan string) .// only allows RECEIVING from channel c")
}

func Sleep(duration uint64) {
	// https://golang.org/doc/effective_go.html
	/* From documentation:
	After waits for the duration to elapse and then sends the current time on the returned channel.
	It is equivalent to NewTimer(d).C.
	The underlying Timer is not recovered by the garbage collector until the timer fires.
	If efficiency is a concern, use NewTimer instead and call Timer.Stop if the timer is no longer needed. */
	<-time.After(time.Second * time.Duration(duration))
}

func problem3() {
	fmt.Println("A buffered channel is an asynchronous channel, whereas regular channels are synchronous.")
	fmt.Println(`To make one with a specific capacity, just do:
		 c := make(chan int, 20) // channel with capacity of 20`)
}
