package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for index := range 10 {
		wg.Add(1)
		go func() {
			fmt.Println(index)
			wg.Done()
		}()
	}
	fmt.Println("Done creating goroutines")
	wg.Wait()
}
