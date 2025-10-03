package main

import "fmt"

func main() {
	fmt.Println("This package has no main, only tests.")
	fmt.Println("To run the tests use:")
	fmt.Println("go test -bench=. -benchmem")
}
