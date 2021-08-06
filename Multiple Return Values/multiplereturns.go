package main 
import(
	"fmt"
)
func main(){
fmt.Println(divide(11,5))
}

func divide(a,b int) (int, int){
	quotient := a/b
	remainder := a%b
	return  quotient,remainder
}