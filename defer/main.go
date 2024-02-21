package main

import "fmt"

func main() {
	fmt.Println("Main code")

	defer fmt.Println("Defer first code")

	fmt.Println("Main code again")

	defer fmt.Println("Clean up code")
}
