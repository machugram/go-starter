// Slice is a powerful data type in go because it allows you to have an array-like structure with other data types.
// Diff b/n Slice and Array
//Slice does not have a fixed size, Array has.

package main

import "fmt"

func main(){
	s:= make([]string, 4)
	s[0] = "Hello"
	s[1] = "Rexford"


	arr :=  [5] int{}
	fmt.Println(arr)

	newSlice := []int{}
	fmt.Println("len newSlice", len(newSlice))
	fmt.Println("Empty Slice newSlice: ", newSlice)

	// newSlice[5] = 70 
	// Why does this give : "panic: runtime error: index out of range [5] with length 0" error when run?

	fmt.Println("One Sliced newSlice: ", newSlice)

	fmt.Println("Before adding World to s: ", s)

	s = append(s,"World")
	fmt.Println("Appended World to s: ", s)

	var sliceCopy = make([]string,5)
	copy(sliceCopy, s)
	fmt.Println("Printing sliceCopy: ", sliceCopy)
}