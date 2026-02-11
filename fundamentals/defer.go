package main

import (
	"fmt"
)

func stacking_defers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}	

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello")
	stacking_defers()
}
