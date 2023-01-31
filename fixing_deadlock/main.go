package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 500; i++ {
			ch <- i
		}
		// closing the channel, so that channel no longer waits for the next value to be written.
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}
