package main

import (
	"fmt"
)

var total int

func addAll(nums ...int) int {
	for _, num := range nums {
		total = +num
		fmt.Println(total)
	}
	fmt.Println(total)

	return total
}
func main() {
	fmt.Println(addAll(3, 4, 5, 9, 3, 4, 2))
}
