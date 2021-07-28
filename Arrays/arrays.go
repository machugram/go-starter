package main 
import (
	"fmt"
	
)


func main(){
	var a  [5] int
	fmt.Println(a)
	fmt.Println("Setting a[4]: ", a[4])

	a[4] = 100
	fmt.Println("Getting a ", a)
}