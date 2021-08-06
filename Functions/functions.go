package main

import "fmt"

func main(){
sumFunction := plus(2,3)
fmt.Println(sumFunction)

minusFunction := minus(6,3,2)
fmt.Println(minusFunction)
}


func plus(a int , b int ) int{
	return a + b 
}

func minus(a  , b  , c int ) int {
	return a - b - c
}