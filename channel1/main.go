package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	sum := 0

	for index := range 10 {
		sumValue := index
		go func() {
			ch <- sumValue
			fmt.Printf("Sent %d to channel\n", sumValue)
		}()
	}

	for msg := range ch {
		fmt.Printf("Received from the channel: %d\n", msg)
		sum += msg
		fmt.Printf("Total sum is now %d\n", sum)

		if sum == 45 {
			close(ch)
		}
	}
	fmt.Printf("Sum is: %d", sum)
}
