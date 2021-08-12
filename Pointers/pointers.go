package main

import (
	"fmt"
)

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
func main() {
	i := 1
	fmt.Println("Initial Value: ", i)
	zeroval(i)
	fmt.Println("zeroval: ", i)
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)
	fmt.Println("Pointer: ", &i)
}
