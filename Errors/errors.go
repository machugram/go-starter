package main
import(
	"errors"
	"fmt"
)

func f1(arg int) (int, error){
	if arg == 400{
		return -1,errors.New("HTTP status code 400 is an error")
	}else{
		return arg,nil
	}
}

func main(){
	fmt.Println(f1(400))
}