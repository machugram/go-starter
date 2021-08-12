package main

import (
	"fmt"
)


type rect struct {
	width ,length int
}

func (r rect) area() int{
	return r.length * r.width
}
func (r rect) perimeter() int{
	return 2 * (r.length + r.width)
}
func main(){
rectOne := rect{4,6}
fmt.Println("Area: ",rectOne.area())
fmt.Println("Perimeter: ",rectOne.perimeter())

}