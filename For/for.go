package main 
import (
	"fmt"
)

func main(){
	i:=1
	for  i <3{
		fmt.Println(i)
		i++
	}


	for j:= 2; j<=9 ; j++{
		fmt.Println(j)
	}


	for w:=0;w<8;w++{
		fmt.Println("Break Loop")
		break
	}

	for e:=0; e<=9;e++{
		if e%2 == 0{
			continue
		}
		fmt.Println(e)
	}
}