package main

import "fmt"
func main(){
	nums := []int{23,45,89,90,78,98}
	sum := 0

	for _, num := range nums{
		sum = sum + num
		fmt.Println("sum: ",sum)
	}

	newmap := map [string] int{"0":0,"1":1}
	for  k := range newmap {
        fmt.Println("key:", k)
    }

	var word = "go"

	for charIndex,unicode := range word{
		// Could reference unicode characters here: https://unicode-table.com/en/#006F
		fmt.Println("char:", map [int] rune {charIndex:unicode})
	}
}