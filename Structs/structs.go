package main

import (
	"fmt"
)

type person struct {
	name string 
	age int
}


func newPerson(name string) *person{
	p := person{name: name}
	p.age = 24
	return &p
}
func main(){
fmt.Println(person{name: "Bob", age: 44})
fmt.Println(person{"Ann",45})
fmt.Println(&person{name: "Jon"})
fmt.Println(newPerson("Tony"))

s := person{"Kai",56}
fmt.Println(s.name)

sp := &s
fmt.Println(sp.age)

sp.age = 54
fmt.Println(sp.age)
}
