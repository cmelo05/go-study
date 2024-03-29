package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	sum := 0

	go func() {
		for index := range 10 {
			ch <- index
			fmt.Printf("Sent %d to channel\n", index)
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Printf("Received from the channel: %d\n", msg)
		sum += msg
		fmt.Printf("Total sum is now %d\n", sum)
	}
	fmt.Printf("Sum is: %d", sum)
}
