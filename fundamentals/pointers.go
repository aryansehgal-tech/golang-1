package main

import (
	"fmt"
)

func main() {
	i, j:= 42, 2701

	p := &i	// point to i
	k := &j
	fmt.Println(*p, *k)
	*p = 21	// setting the value of i through pointer
	fmt.Println(i)

	*k = *k / 37	// divide j throught the pointer
	fmt.Println(j)
}
