package main 

import(
	"fmt"
)

func main(){
	if 7%2 == 0{
		fmt.Println("Seven is even")
	}else {
		fmt.Println("Seven is odd")
	}


	if num:= 9; num <0 {
		fmt.Println("Number is negative") 
	}	else if num <10 {  //If the else if goes to the next line , there is an error. GoLang special.
		fmt.Println("Number is 1 digit")
	}	else {
		fmt.Println("XXX")
	}
}