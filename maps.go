package main
import "fmt"
func main(){
	x := make(map[string] int)
	x["a"]= 6
	x["e"]= 5
	x["i"]= 4
	x["o"]= 3
	x["u"]= 2
	if value,ok:=x["a"]; ok{
		fmt.Println(value,ok)
	}

	//fmt.Println(value)
	//fmt.Println(ok)
	fmt.Println(x)
	fmt.Println(len(x))
	v := x
	fmt.Println(v)
}