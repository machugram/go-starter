package main

import (
	"fmt"
	"time"
)

func main(){
	i := 2
	fmt.Print("Write ",i , " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Another number")
	}


	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
			fmt.Println("Its the weekend")
	default:
			fmt.Println("Its the weekday")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type) {
		case bool:
			fmt.Printf("I am a Bool %T\n" ,t)
		case int:
			fmt.Println("I am an Int")
		default:
            fmt.Printf("Don't know type %T\n", t) // I think the %T helps insert the type
		}
	}

	whatAmI("hey")
	whatAmI(false)
}