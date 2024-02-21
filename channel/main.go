package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	sum := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			received := <-ch
			fmt.Printf("Received from the channel: %d\n", received)
			sum += received
			fmt.Printf("Total sum is now %d\n", sum)
			if sum == 45 {
				break
			}
		}
	}()

	for index := range 10 {
		wg.Add(1)
		go func() {
			ch <- index
			fmt.Printf("Sent %d to channel\n", index)
			defer wg.Done()
		}()
	}

	wg.Wait()
	fmt.Printf("Sum is: %d", sum)
}
