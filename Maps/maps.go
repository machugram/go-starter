package main 
import(
	"fmt"
)
func main(){
	smap := make ( map [string] int)
	smap["key0"] = 23
	smap["key1"] = 25

	fmt.Println("Map: ", smap)

	valKey0 := smap["key0"]
	fmt.Println(valKey0)

	fmt.Println("len: ",len(smap))

	delete(smap, "key0")
	fmt.Println("Deleted key0. Map : ", smap)

	smap2 := map [string] int{"one": 1, "two":2}
	fmt.Println(smap2)

}